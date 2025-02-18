package translator

import (
	"sync"
)

var (
	mu               sync.RWMutex
	translators      []Translator
	activeLang       string
	activeTranslator Translator
)

// Translator interface for translators
type Translator interface {
	Translate(key string, args ...interface{}) (error, string)
	SetLanguage(lang string) error
	GetLanguage() string
}

// Init initializing map of translators
func Init() {
	mu.Lock()
	defer mu.Unlock()
	translators = []Translator{}
}

// Register add Translator
func Register(translator Translator) {
	mu.Lock()
	defer mu.Unlock()
	translators = append(translators, translator)
	activeTranslator = translator
}

// SetLanguage set active language
func SetLanguage(lang string) error {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range translators {
		err := t.SetLanguage(lang)
		if err != nil {
			return err
		}
	}

	activeLang = lang
	return nil
}

// GetLanguage return active language
func GetLanguage() string {
	mu.RLock()
	defer mu.RUnlock()
	return activeLang
}

// Translate translating by active language
func Translate(key string, args ...interface{}) string {
	mu.RLock()
	defer mu.RUnlock()

	for _, t := range translators {
		err, tr := t.Translate(key, args...)
		if err != nil {
			continue
		}

		return tr
	}

	return key
}
