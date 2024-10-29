package translation

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var translationsTestBuf = []byte(`{
	"de":{
		"global":{
			"username":"Benutzername",
			"password":"Passwort"
		},
		"validation":{
			"required":"%s ist erforderlich!"
		},
		"channel": {
			"validation": {
				"name_required": "Der Channel-Name ist erforderlich!"
			}
		}
	},
	"en":{
		"global":{
			"username":"Username",
			"password":"Password"
		},
		"validation":{
			"required":"%s is mandatory!"
		}
	}
}`)

var translationsTestInvalidBuf = []byte(`{
	"de":
		"global":{
			"username":"Benutzername",
			"password":"Passwort"
		},
		"validation":{
			"required":"%s ist erforderlich!"
		}
	}
}`)

func TestTranslator_New(t *testing.T) {
	tr := New()

	assert.NotNil(t, tr)
	assert.Equal(t, DefaultLanguage, tr.language)
}

func TestTranslator_NewFromBuffer(t *testing.T) {
	tr, err := NewFromBuffer(&translationsTestBuf)

	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, "Benutzername", tr.T("global.username"))
	assert.Equal(t, "Username", tr.TL("en", "global.username"))
}

func TestTranslator_NewFromFile(t *testing.T) {
	fileReader = func(filename string) ([]byte, error) {
		return translationsTestBuf, nil
	}

	tr, err := NewFromFile("foo.json")

	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, "Benutzername", tr.T("global.username"))
	assert.Equal(t, "Username", tr.TL("en", "global.username"))
}

func TestTranslator_NewFromFileWithError(t *testing.T) {
	fileReader = func(filename string) ([]byte, error) {
		return nil, errors.New("could not read json file")
	}

	tr, err := NewFromFile("foo.json")

	assert.NotNil(t, err)
	assert.Nil(t, tr)
}

func TestTranslator_NewFromFileWithInvalidJson(t *testing.T) {
	fileReader = func(filename string) ([]byte, error) {
		return translationsTestInvalidBuf, nil
	}

	tr, err := NewFromFile("foo.json")

	assert.NotNil(t, err)
	assert.Nil(t, tr)
}

func TestTranslator_NewFromBufferInvalidJson(t *testing.T) {
	tr, err := NewFromBuffer(&translationsTestInvalidBuf)

	assert.NotNil(t, err)
	assert.Nil(t, tr)
}

func TestTranslator_NotFound(t *testing.T) {
	tr, err := NewFromBuffer(&translationsTestBuf)

	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, "global.foo", tr.T("global.foo"))
	assert.Equal(t, "global.foo", tr.TL("es", "global.foo"))
}

func TestTranslator_IncompletePath(t *testing.T) {
	tr, err := NewFromBuffer(&translationsTestBuf)

	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, "channel.validation", tr.T("channel.validation"))
	assert.Equal(t, "Der Channel-Name ist erforderlich!", tr.T("channel.validation.name_required"))
}

func TestTranslator_WithLanguage(t *testing.T) {
	tr, err := NewFromBuffer(&translationsTestBuf)

	assert.Nil(t, err)

	trEn := tr.WithLanguage("en")

	assert.NotSame(t, tr, trEn)
	assert.Equal(t, "Username", trEn.T("global.username"))
	assert.Equal(t, "Benutzername", tr.T("global.username"))
}
