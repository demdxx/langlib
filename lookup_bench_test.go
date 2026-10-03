package langlib

import "testing"

var (
	benchLanguage *Language
	benchCode     Code
	benchString   string
	benchID       ID
	benchBytes    []byte
	benchIDs      []uint64
	benchBool     bool
)

func BenchmarkGetLanguageByCodeString(b *testing.B) {
	b.ReportAllocs()
	var lang *Language
	for i := 0; i < b.N; i++ {
		lang = GetLanguageByCodeString("en")
	}
	benchLanguage = lang
}

func BenchmarkCodeLanguage(b *testing.B) {
	code := Code{'e', 'n'}
	b.ReportAllocs()
	var lang *Language
	for i := 0; i < b.N; i++ {
		lang = code.Language()
	}
	benchLanguage = lang
}

func BenchmarkCodeISO2(b *testing.B) {
	code := Code{'e', 'n'}
	b.ReportAllocs()
	var s string
	for i := 0; i < b.N; i++ {
		s = code.ISO2()
	}
	benchString = s
}

func BenchmarkCodeFromStringBCP47(b *testing.B) {
	b.ReportAllocs()
	var c Code
	for i := 0; i < b.N; i++ {
		c = CodeFromString("en-US")
	}
	benchCode = c
}

func BenchmarkIDFromString(b *testing.B) {
	b.ReportAllocs()
	var id ID
	for i := 0; i < b.N; i++ {
		id = IDFromString("en")
	}
	benchID = id
}

func BenchmarkIDLanguage(b *testing.B) {
	id := ID(40)
	b.ReportAllocs()
	var lang *Language
	for i := 0; i < b.N; i++ {
		lang = id.Language()
	}
	benchLanguage = lang
}

func BenchmarkCodeID(b *testing.B) {
	code := Code{'e', 'n'}
	b.ReportAllocs()
	var id ID
	for i := 0; i < b.N; i++ {
		id = code.ID()
	}
	benchID = id
}

func BenchmarkIDIsUndefined(b *testing.B) {
	id := ID(40)
	b.ReportAllocs()
	var undef bool
	for i := 0; i < b.N; i++ {
		undef = id.IsUndefined()
	}
	benchBool = undef
}

func BenchmarkLanguageISO2(b *testing.B) {
	lang := ID(40).Language()
	b.ReportAllocs()
	var s string
	for i := 0; i < b.N; i++ {
		s = lang.ISO2()
	}
	benchString = s
}

func BenchmarkGetLanguageByCodeStringMiss(b *testing.B) {
	b.ReportAllocs()
	var lang *Language
	for i := 0; i < b.N; i++ {
		lang = GetLanguageByCodeString("xx")
	}
	benchLanguage = lang
}

func BenchmarkIDMarshalJSON(b *testing.B) {
	id := ID(40)
	b.ReportAllocs()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf, _ = id.MarshalJSON()
	}
	benchBytes = buf
}

func BenchmarkIDUnmarshalJSON(b *testing.B) {
	data := []byte(`"en-US"`)
	b.ReportAllocs()
	var id ID
	for i := 0; i < b.N; i++ {
		_ = id.UnmarshalJSON(data)
	}
	benchID = id
}

func BenchmarkCodeMarshalJSON(b *testing.B) {
	code := Code{'E', 'N'}
	b.ReportAllocs()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf, _ = code.MarshalJSON()
	}
	benchBytes = buf
}

func BenchmarkCodeUnmarshalJSON(b *testing.B) {
	data := []byte(`"EN"`)
	b.ReportAllocs()
	var code Code
	for i := 0; i < b.N; i++ {
		_ = code.UnmarshalJSON(data)
	}
	benchCode = code
}

func BenchmarkLangCodes2IDs(b *testing.B) {
	codes := []string{"en", "de", "fr", "xx", "zh-Hans"}
	b.ReportAllocs()
	var ids []uint64
	for i := 0; i < b.N; i++ {
		ids = LangCodes2IDs[uint64](codes)
	}
	benchIDs = ids
}

func BenchmarkLangCodes2SortedIDs(b *testing.B) {
	codes := []string{"en", "de", "fr", "xx", "zh-Hans"}
	b.ReportAllocs()
	var ids []uint64
	for i := 0; i < b.N; i++ {
		ids = LangCodes2SortedIDs[uint64](codes)
	}
	benchIDs = ids
}
