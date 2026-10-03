package langlib

// Language is a catalog entry. Slice index equals ID.
type Language struct {
	ID         ID     `json:"ID"`
	Code       Code   `json:"code"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
}

// ISO2 returns the interned two-letter code (never allocates).
func (l *Language) ISO2() string {
	if l == nil {
		return UndefinedISO2
	}
	return iso2(l.ID)
}
