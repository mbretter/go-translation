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

## Usage

Basically there are two translation functions, T and TL. T accepts the translation key only, while TL accepts the 
language code as first argument.

The translation key is a representation of the path to the final field name in dot format. If the key was not found 
the key itself will be returned. This has two advantages, first you will find easily non translated keys and second 
it helps a lot when testing your code.

The translations functions are supporting printf-like formating.

```go
translator, err := translation.NewFromFile("./assets/translations.json")
if err != nil {
    log.Error("unable to load translations: " + err.Error())
    return
}

translator = translator.WithLanguage("en")

fmt.Println(translator.T("user.username")) // prints Username
fmt.Println(translator.TL("de", "user.username")) // prints Benutzername
fmt.Println(translator.T("some.path.not.found")) // prints some.path.not.found
fmt.Println(tr.T("validation.required", tr.T("user.username")) // prints Username is mandatory!
```

## Accept-Language header parsing

When writing APIs it could be useful to change the translation language based on the Accept-Language header, which 
was sent by the client.

The ParseAcceptLanguage function parses the header by returning a slice of languages sorted by the quality value in 
descending order, i.e. languages with a higher quality level comes first.

```go
langs := ParseAcceptLanguage("de;q=1,de-AT;q=0.8,fr;q=0.2")
fmt.Println(langs) // outputs: [{de de  1} {de-AT de AT 0.8} {fr fr  0.2}]
```

The returned AcceptLanguage struct is defined like this:
```go
type AcceptLanguage struct {
    Lang    string // de-AT, de
    Base    string // de, de
    Region  string // AT, ""
    Quality float64 // defaults to 1
}
```

### http middleware

You could put the accept-language header parsing into a http middleware, here is a code snippet how this could be made.

In this case the first language with the highest quality is used.
The language as well as the translator is puted into the context and could later be used for accessing the translator.

```go
func HttpLanguageMiddleware(translator *translation.Translator) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        fn := func(w http.ResponseWriter, r *http.Request) {
    
            headerLine := r.Header.Get("Accept-Language")
            if headerLine == "" {
                headerLine = "en"
            }
    
            var lang translation.AcceptLanguage
            languages := translation.ParseAcceptLanguage(headerLine)
            if len(languages) > 0 {
                lang = languages[0]
            } else {
                lang = translation.AcceptLanguage{Lang: "en", Base: "en"}
            }
    
            ctx := context.WithValue(r.Context(), "lang", lang)
            ctx = context.WithValue(ctx, "translator", translator.WithLanguage(lang.Base))
    
            next.ServeHTTP(w, r.WithContext(ctx))
        }
    
        return http.HandlerFunc(fn)
    }
}

// somewhere in the request handler function
func (a *Api) MyHandler(w http.ResponseWriter, r *http.Request) {
    tr := r.Context().Value("translator").(*translation.Translator)
	
	// ...
	
	// tr.T("user.username")
}
```

