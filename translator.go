// Package translation provides a translation service based on JSON files.
package translation

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"os"
)

const (
	DefaultLanguage = "de"
)

type Translator struct {
	container *gabs.Container
	language  string
}

// New creates a new instance of Translator with the default language.
func New() *Translator {
	t := Translator{
		language: DefaultLanguage,
	}
	return &t
}

// NewFromBuffer creates a new Translator instance from the provided JSON buffer.
func NewFromBuffer(buf *[]byte) (*Translator, error) {
	t := New()
	err := t.parse(buf)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// NewFromFile loads a translator from a JSON file specified by the filename.
func NewFromFile(filename string) (*Translator, error) {
	buf, err := load(filename)
	if err != nil {
		return nil, err
	}

	t, err := NewFromBuffer(buf)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// WithLanguage sets the language for the Translator instance and returns a new Translator instance with the specified language.
func (t *Translator) WithLanguage(lang string) *Translator {
	tc := *t
	tc.language = lang

	return &tc
}

// T translates the given key using the current language and optional arguments.
func (t *Translator) T(key string, args ...any) string {
	return t.TL(t.language, key, args...)
}

// TL translates the specified key using the provided language and optional arguments. Returns the key if translation fails.
func (t *Translator) TL(lang string, key string, args ...any) string {
	txt, ok := t.container.Path(lang + "." + key).Data().(string)
	if !ok {
		return key
	}

	return fmt.Sprintf(txt, args...)
}

var fileReader = os.ReadFile

func load(filename string) (*[]byte, error) {
	//content, err := os.ReadFile(filename)
	content, err := fileReader(filename)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (t *Translator) parse(data *[]byte) error {
	cont, err := gabs.ParseJSON(*data)
	if err != nil {
		return err
	}

	t.container = cont

	return nil
}
