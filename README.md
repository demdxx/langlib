# langlib

[![Test](https://github.com/demdxx/langlib/actions/workflows/test.yaml/badge.svg)](https://github.com/demdxx/langlib/actions/workflows/test.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/demdxx/langlib.svg)](https://pkg.go.dev/github.com/demdxx/langlib)
[![Go Report Card](https://goreportcard.com/badge/github.com/demdxx/langlib)](https://goreportcard.com/report/github.com/demdxx/langlib)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

Catalog of ISO-639-1 languages and two value types for storing one: `Code` (two letters) and `ID` (catalog index).

Lookup uses the first two letters, so BCP-47 tags work (`en-US`, `zh-Hans`). Case does not matter. An unknown or short value resolves to `Languages[0]`: `ID` 0, ISO code `**`, name `Undefined`. The catalog index is the language ID. Resolving a code does not allocate.

```
go get github.com/demdxx/langlib
```

## Lookup

```go
lang := langlib.GetLanguageByCodeString("en-US")
fmt.Println(lang.ISO2(), lang.Name, lang.ID) // en English 40

code := langlib.CodeFromString("DE")
fmt.Println(code.ISO2(), code.ID(), code.Language().Name) // de 51 German

id := langlib.IDFromString("fr")
fmt.Println(id.Language().Name) // French

if langlib.IDFromString("xx").IsUndefined() {
    fmt.Println(langlib.UndefinedISO2) // **
}
```

`Languages` is the full catalog. `Languages[i].ID` equals `i`.

## JSON

`Code` and `ID` marshal as an ISO-639-1 string. Unknown values marshal as `"**"`.

```go
type prefs struct {
    Lang langlib.ID   `json:"lang"`
    Code langlib.Code `json:"code"`
}

var p prefs
_ = json.Unmarshal([]byte(`{"lang":"en-US","code":"DE"}`), &p)
data, _ := json.Marshal(p)
fmt.Println(string(data)) // {"lang":"en","code":"de"}
```

`ID` also accepts a JSON number, the catalog index.

## Lists of codes

`LangCodes2IDs` keeps input order. `LangCodes2SortedIDs` sorts the IDs. The element type is any integer.

```go
ids := langlib.LangCodes2IDs[uint32]([]string{"en", "xx", "DE"})
// 40, 0, 51

sorted := langlib.LangCodes2SortedIDs[uint32]([]string{"en", "xx", "DE"})
// 0, 40, 51
```

## SQL

`Code` is stored as the two-letter string (`en`). `ID` is stored as `int64` and can be scanned from an integer or from a language code.
