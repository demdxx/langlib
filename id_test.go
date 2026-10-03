package langlib

import (
	"encoding/json"
	"errors"
	"strconv"
	"testing"
)

func Test_IDCatalog(t *testing.T) {
	undef := UndefinedID.Language()
	if !UndefinedID.IsUndefined() || undef.Name != "Undefined" || undef.ISO2() != UndefinedISO2 {
		t.Fatalf("undefined: %+v %q %q", undef, undef.Name, undef.ISO2())
	}
	if undef.Code != UndefinedCode || undef.NativeName != "" {
		t.Fatalf("undefined code=%q native=%q", undef.Code, undef.NativeName)
	}

	for i, lg := range Languages {
		id := ID(i)
		got := id.Language()
		if got != &languages[i] {
			t.Fatalf("ID(%d).Language is not catalog slot", i)
		}
		if id.Code() != lg.Code || id.ISO2() != lg.ISO2() {
			t.Fatalf("ID(%d) code %q ISO2 %q", i, id.Code(), id.ISO2())
		}
		if got.Code != lg.Code || got.Name != lg.Name || got.NativeName != lg.NativeName || got.ISO2() != lg.ISO2() {
			t.Fatalf("ID(%d) data mismatch: %+v", i, lg)
		}
		if got.Code.Int() != lg.Code.Int() {
			t.Fatalf("ID(%d) numeric mismatch", i)
		}
		if id.IsUndefined() != (i == 0) {
			t.Fatalf("ID(%d) IsUndefined=%v", i, id.IsUndefined())
		}
		if i == 0 {
			continue
		}
		if IDFromString(lg.ISO2()) != id {
			t.Fatalf("IDFromString(%q) != %d", lg.ISO2(), id)
		}
		if lg.Code.ID() != id {
			t.Fatalf("Code.ID round-trip %d", id)
		}
	}
}

func Test_IDFromString(t *testing.T) {
	en := IDFromString("en")
	if en.Language().Name != "English" {
		t.Fatalf("en: %s", en.Language().Name)
	}
	tests := []struct {
		in   string
		want ID
	}{
		{"EN", en},
		{"en-US", en},
		{"en_GB", en},
		{"**", UndefinedID},
		{"xx", UndefinedID},
		{"", UndefinedID},
		{"X", UndefinedID},
	}
	for _, test := range tests {
		if got := IDFromString(test.in); got != test.want {
			t.Errorf("IDFromString(%q)=%d, want %d", test.in, got, test.want)
		}
	}
	if !ID(len(languages)).IsUndefined() || ID(len(languages)).Language() != &languages[0] {
		t.Fatal("out-of-range ID should be undefined")
	}
}

func Test_IDJSON(t *testing.T) {
	en := IDFromString("en")
	data, err := json.Marshal(en)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `"`+en.Language().ISO2()+`"` {
		t.Fatalf("marshal = %s", data)
	}

	tests := []struct {
		source string
		want   ID
	}{
		{strconv.FormatUint(uint64(en), 10), en},
		{`"en"`, en},
		{`"EN"`, en},
		{`"en-US"`, en},
		{`"**"`, UndefinedID},
		{`""`, UndefinedID},
		{"null", UndefinedID},
		{"0", UndefinedID},
	}
	for _, test := range tests {
		var id ID
		if err := json.Unmarshal([]byte(test.source), &id); err != nil {
			t.Errorf("unmarshal %s: %v", test.source, err)
			continue
		}
		if id != test.want {
			t.Errorf("unmarshal %s = %d, want %d", test.source, id, test.want)
		}
	}

	var id ID
	if err := json.Unmarshal([]byte(`"12x"`), &id); !errors.Is(err, ErrIDInvalidValue) {
		t.Fatalf("unmarshal garbage err = %v", err)
	}
	if err := json.Unmarshal([]byte("true"), &id); err == nil {
		t.Fatal("unmarshal bool should fail")
	}
}

func Test_IDScan(t *testing.T) {
	var id ID
	if err := id.Scan(int64(40)); err != nil {
		t.Fatal(err)
	}
	if id.Language().Name != "English" {
		t.Fatalf("scan int64: %s", id.Language().Name)
	}
	if err := id.Scan("DE"); err != nil {
		t.Fatal(err)
	}
	if id != IDFromString("de") {
		t.Fatalf("scan string: %d", id)
	}
	if err := id.Scan([]byte("en-US")); err != nil {
		t.Fatal(err)
	}
	if id != IDFromString("en") {
		t.Fatalf("scan bytes: %d", id)
	}
	if err := id.Scan(nil); err != nil || id != UndefinedID {
		t.Fatalf("scan nil: %d %v", id, err)
	}
	if err := id.Scan(int64(-1)); !errors.Is(err, ErrIDInvalidValue) {
		t.Fatalf("scan negative err = %v", err)
	}
	if err := id.Scan("12x"); !errors.Is(err, ErrIDInvalidValue) {
		t.Fatalf("scan garbage err = %v", err)
	}
	if err := id.Scan(1.5); !errors.Is(err, ErrIDInvalidScanType) {
		t.Fatalf("scan float err = %v", err)
	}

	v, err := ID(40).Value()
	if err != nil {
		t.Fatal(err)
	}
	if v != int64(40) {
		t.Fatalf("Value = %v", v)
	}
}

func Test_IDNoAlloc(t *testing.T) {
	id := IDFromString("en")
	if n := testing.AllocsPerRun(1000, func() { _ = id.Language() }); n != 0 {
		t.Fatalf("ID.Language allocs = %v", n)
	}
	if n := testing.AllocsPerRun(1000, func() { _ = id.ISO2() }); n != 0 {
		t.Fatalf("ID.ISO2 allocs = %v", n)
	}
	if n := testing.AllocsPerRun(1000, func() { _ = id.Code() }); n != 0 {
		t.Fatalf("ID.Code allocs = %v", n)
	}
	if n := testing.AllocsPerRun(1000, func() { _ = IDFromString("en-US") }); n != 0 {
		t.Fatalf("IDFromString allocs = %v", n)
	}
}
