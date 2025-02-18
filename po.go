package translator

import (
	"errors"
	"github.com/gouef/finder"
	"github.com/gouef/utils"
	"github.com/leonelquinteros/gotext"
	"sync"
)

type PO struct {
	mu         sync.RWMutex
	locales    map[string]*gotext.Locale
	activeLang string
}

// NewPO create new Translator implementation for PO
func NewPO(dir string) (*PO, error) {
	pt := &PO{locales: make(map[string]*gotext.Locale)}

	locales := finder.FindDirectories("[a-z][a-z]", "[a-z][a-z]_[A-Z][A-Z]").In(dir).Get()

	for path, d := range locales {
		code := d.Name
		locale := gotext.NewLocale(dir, code)

		entries := finder.FindFiles("*.po", "*.mo").In(path).Get()

		for _, entry := range entries {
			domain := entry.Name
			locale.AddDomain(domain)
		}

		pt.locales[code] = locale
	}

	if len(pt.locales) == 0 {
		return nil, errors.New("no valid PO translations found")
	}

	return pt, nil
}

// SetLanguage set active language
func (p *PO) SetLanguage(lang string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, exists := p.locales[lang]; !exists {
		return errors.New("translator not found for language: " + lang)
	}
	p.activeLang = lang
	return nil
}

// GetLanguage return active language
func (p *PO) GetLanguage() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.activeLang
}

// Translate translating by active language
func (p *PO) Translate(key string, args ...interface{}) (error, string) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	keys := utils.Explode(".", key)
	locale := p.locales[p.activeLang]

	if len(keys) == 2 {
		locale.AddDomain(keys[0])
		return nil, locale.GetD(keys[0], keys[1], args...)
	}

	return nil, locale.Get(key, args...)
}
