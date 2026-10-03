package langlib

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

// UndefinedID is the catalog ID of an unknown language (Languages[0]).
const UndefinedID ID = 0

var (
	ErrIDInvalidScanType = errors.New("[langlib.id] invalid scan type, supports only integer, bytes and string")
	ErrIDInvalidValue    = errors.New("[langlib.id] invalid value")
)

// ID is a catalog language identifier. Index in [Languages] equals the ID.
// Zero and out-of-range values resolve to Languages[0].
type ID uint8

// IDFromString resolves a BCP-47 / ISO-639-1 string (first two letters).
// Unknown and short values return [UndefinedID].
func IDFromString(s string) ID {
	if len(s) < 2 {
		return UndefinedID
	}
	return lookupID(s[0], s[1])
}

// ISO2 returns the interned two-letter catalog code, or [UndefinedISO2].
func (id ID) ISO2() string {
	return iso2(id)
}

// Code returns the catalog ISO code, or [UndefinedCode].
func (id ID) Code() Code {
	s := id.ISO2()
	return Code{s[0], s[1]}
}

// Language returns the catalog entry. Unknown IDs resolve to Languages[0] (never nil).
func (id ID) Language() *Language {
	if uint(id) >= uint(len(languages)) {
		return &languages[0]
	}
	return &languages[id]
}

// IsUndefined reports whether id maps to Languages[0].
func (id ID) IsUndefined() bool {
	return id == UndefinedID || uint(id) >= uint(len(languages))
}

// Value implementation of sql driver.Valuer interface. The value is int64.
func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}

// Scan implementation of sql.Scanner interface.
// Integers are catalog IDs. Strings and byte slices are either a decimal ID
// or a language code ([IDFromString]).
func (id *ID) Scan(data any) error {
	switch v := data.(type) {
	case nil:
		*id = UndefinedID
		return nil
	case int64:
		return id.setInt64(v)
	case int32:
		return id.setInt64(int64(v))
	case int16:
		return id.setInt64(int64(v))
	case int8:
		return id.setInt64(int64(v))
	case int:
		return id.setInt64(int64(v))
	case uint64:
		return id.setUint64(v)
	case uint32:
		return id.setUint64(uint64(v))
	case uint16:
		return id.setUint64(uint64(v))
	case uint8:
		return id.setUint64(uint64(v))
	case uint:
		return id.setUint64(uint64(v))
	case []byte:
		return id.scanBytes(v)
	case string:
		return id.scanText(v)
	default:
		return ErrIDInvalidScanType
	}
}

func (id *ID) setInt64(v int64) error {
	if v < 0 {
		return ErrIDInvalidValue
	}
	return id.setUint64(uint64(v))
}

func (id *ID) setUint64(v uint64) error {
	if v > math.MaxUint8 {
		return ErrIDInvalidValue
	}
	*id = ID(v)
	return nil
}

func (id *ID) scanText(s string) error {
	if s == "" || s == UndefinedISO2 {
		*id = UndefinedID
		return nil
	}
	if len(s) >= 2 && isASCIILetter(s[0]) && isASCIILetter(s[1]) {
		*id = lookupID(s[0], s[1])
		return nil
	}
	if n, err := strconv.ParseUint(s, 10, 64); err == nil {
		return id.setUint64(n)
	}
	return ErrIDInvalidValue
}

func (id *ID) scanBytes(b []byte) error {
	if len(b) == 0 || (len(b) == 2 && b[0] == '*' && b[1] == '*') {
		*id = UndefinedID
		return nil
	}
	if len(b) >= 2 && isASCIILetter(b[0]) && isASCIILetter(b[1]) {
		*id = lookupID(b[0], b[1])
		return nil
	}
	if n, err := strconv.ParseUint(string(b), 10, 64); err == nil {
		return id.setUint64(n)
	}
	return ErrIDInvalidValue
}

func isASCIILetter(b byte) bool {
	b = asciiLower(b)
	return b >= 'a' && b <= 'z'
}

// MarshalJSON encodes the ISO-639-1 code. Unknown IDs are [UndefinedISO2].
func (id ID) MarshalJSON() ([]byte, error) {
	s := id.ISO2()
	return []byte{'"', s[0], s[1], '"'}, nil
}

// UnmarshalJSON accepts a JSON number or a language-code string.
// Null, "", and "**" become [UndefinedID].
func (id *ID) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && data[0] == 'n' && data[1] == 'u' && data[2] == 'l' && data[3] == 'l' {
		*id = UndefinedID
		return nil
	}
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		return id.scanBytes(data[1 : len(data)-1])
	}
	var n uint64
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	return id.setUint64(n)
}

var (
	_ driver.Valuer    = UndefinedID
	_ sql.Scanner      = (*ID)(nil)
	_ json.Marshaler   = UndefinedID
	_ json.Unmarshaler = (*ID)(nil)
)
