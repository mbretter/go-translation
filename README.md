[![](https://github.com/mbretter/go-translation/actions/workflows/test.yml/badge.svg)](https://github.com/mbretter/go-translation/actions/workflows/test.yml)
[![](https://goreportcard.com/badge/mbretter/go-translation)](https://goreportcard.com/report/mbretter/go-translation "Go Report Card")
[![codecov](https://codecov.io/gh/mbretter/go-translation/graph/badge.svg?token=YMBMKY7W9X)](https://codecov.io/gh/mbretter/go-translation)
[![GoDoc](https://godoc.org/github.com/mbretter/go-translation?status.svg)](https://pkg.go.dev/github.com/mbretter/go-translation)

Simple translation system based on JSON files.

It also provides a helper function for parsing the AcceptLanguage header.

## JSON File format

The JSON file which contains the translations has at the top level object the language code as property, this is 
the only prerequisite. You can put as many levels as you want below the language code.

```json
{
    "de":{
        "user":{
            "username":"Benutzername",
            "password":"Passwort"
        },
        "validation":{
            "required":"%s ist erforderlich!"
        }
    },
    "en":{
        "user":{
            "username":"Username",
            "password":"Password"
        },
        "validation":{
            "required":"%s is mandatory!"
        }
    }
}
```

