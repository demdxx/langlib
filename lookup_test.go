package langlib

import "testing"

func Test_IDLanguage(t *testing.T) {
	if got := ID(40).Language(); got.Name != "English" {
		t.Fatalf("id 40: %+v", got)
	}
	if got := ID(0).Language(); !got.Code.IsUndefined() {
		t.Fatalf("id 0 should be undefined: %+v", got)
	}
	if got := ID(len(languages)).Language(); got != &languages[0] {
		t.Fatal("out-of-range id should be Languages[0]")
	}
}

func Test_CatalogIndexEqualsID(t *testing.T) {
	if Languages[0].Code != UndefinedCode {
		t.Fatalf("Languages[0] code = %q", Languages[0].Code)
	}
	for i, lg := range Languages {
		if lg.ID != ID(i) {
			t.Fatalf("index %d has ID %d", i, lg.ID)
		}
		if lg.ISO2() != string([]byte{lg.Code[0], lg.Code[1]}) {
			t.Fatalf("index %d ISO2 %q != code %q", i, lg.ISO2(), lg.Code)
		}
		if lg.ID.Language() != &languages[i] {
			t.Fatalf("ID(%d).Language is not catalog slot", lg.ID)
		}
		if i == 0 {
			continue
		}
		if lg.Code[0] < 'a' || lg.Code[0] > 'z' || lg.Code[1] < 'a' || lg.Code[1] > 'z' {
			t.Fatalf("code %q is not lowercase", lg.Code)
		}
		if lg.Code.Language() != &languages[i] {
			t.Fatalf("code %q does not round-trip", lg.Code)
		}
	}
}

func Test_LangCodes2IDs(t *testing.T) {
	if LangCodes2IDs[uint64](nil) != nil {
		t.Fatal("nil codes should return nil")
	}
	ids := LangCodes2IDs[uint64]([]string{"en", "xx", "DE"})
	if len(ids) != 3 {
		t.Fatalf("len=%d", len(ids))
	}
	if ids[0] != uint64(IDFromString("EN")) || ids[1] != 0 || ids[2] != uint64(IDFromString("DE")) {
		t.Fatalf("input order, got %v", ids)
	}

	sorted := LangCodes2SortedIDs[uint64]([]string{"en", "xx", "DE"})
	if len(sorted) != 3 || sorted[0] != 0 || sorted[1] != ids[0] || sorted[2] != ids[2] {
		t.Fatalf("sorted ids, got %v", sorted)
	}
}
