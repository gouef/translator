package translator

import (
	"errors"
	"fmt"
	"github.com/gouef/finder"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync"
)

type Yaml struct {
	mu         sync.RWMutex
	locales    map[string]map[string]interface{}
	activeLang string
}

// NewYaml create new Translator implementation for Yaml
func NewYaml(dir string) (*Yaml, error) {
	yt := &Yaml{
		locales: make(map[string]map[string]interface{}),
	}
	locales := finder.FindDirectories("[a-z][a-z]", "[a-z][a-z]_[A-Z][A-Z]").In(dir).Get()

	for path, d := range locales {
		code := d.Name
		yt.locales[code] = make(map[string]interface{})

		entries := finder.FindFiles("*.yml", "*.yaml").In(path).Get()

		for _, entry := range entries {
			data, err := os.ReadFile(entry.Path)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error reading file: %v", err))
			}

			var result map[string]interface{}
			if err := yaml.Unmarshal([]byte(data), &result); err != nil {
				return nil, errors.New(fmt.Sprintf("no valid yaml translations \"%s\"", entry.Path))
			}
			yt.locales[code][entry.Name] = result
		}
	}

	if len(yt.locales) == 0 {
		return nil, errors.New("no valid yaml translations found")
	}

	return yt, nil
}

// SetLanguage set active language
func (p *Yaml) SetLanguage(lang string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, exists := p.locales[lang]; !exists {
		return errors.New("translator not found for language: " + lang)
	}
	p.activeLang = lang
	return nil
}

// GetLanguage return active language
func (p *Yaml) GetLanguage() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.activeLang
}

// Translate translating by active language
func (p *Yaml) Translate(key string, args ...interface{}) (error, string) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	value := p.getKey(key, p.locales[p.activeLang])
	if value == nil {
		return errors.New(fmt.Sprintf("missing translation for \"%s\"", key)), key
	}

	if str, ok := value.(string); ok {
		return nil, fmt.Sprintf(str, args...)
	}

	return nil, fmt.Sprintf("%v", value)
}

func (p *Yaml) getKey(key string, data map[string]interface{}) interface{} {
	keys := strings.Split(key, ".")
	return p.getKeyRecursive(keys, data)
}

func (p *Yaml) getKeyRecursive(keys []string, data map[string]interface{}) interface{} {
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	value, exists := data[key]
	if !exists {
		return nil
	}

	if len(keys) == 1 {
		return value
	}

	if nested, ok := value.(map[string]interface{}); ok {
		return p.getKeyRecursive(keys[1:], nested)
	}

	return nil
}
