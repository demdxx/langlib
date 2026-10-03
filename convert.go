package langlib

import "slices"

// NumberType is the integer type of the returned language IDs.
type NumberType interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// LangCodes2IDs converts language codes to language IDs, in input order.
// Unknown codes are 0. An empty input returns nil.
func LangCodes2IDs[T NumberType](codes []string) []T {
	if len(codes) == 0 {
		return nil
	}
	result := make([]T, len(codes))
	for i, lg := range codes {
		result[i] = T(IDFromString(lg))
	}
	return result
}

// LangCodes2SortedIDs converts language codes to language IDs, sorted ascending.
// Unknown codes are 0. An empty input returns nil.
func LangCodes2SortedIDs[T NumberType](codes []string) []T {
	res := LangCodes2IDs[T](codes)
	slices.Sort(res)
	return res
}
