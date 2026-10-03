package langlib

// languageISO2 packs the interned ISO-639-1 code for each catalog index, two bytes each.
// Index 0 is UndefinedISO2.
const languageISO2 = UndefinedISO2 + "abaaafaksqamaranhyasavaeayazbmbaeubebnbhbibsbrbgmycachcenyzhcvkwcocrhrcsdadvnleneoeteefofjfifrffglkadeelgnguhthahehzhihohuiaidiegaigikioisitiujajvklknkrkskkkmkirwkykvkgkokukjlalblglilnloltlulvgvmkmgmsmlmtmimrmhmnnanvnbndnengnnnoiinrocojcuomorospapifaplpsptqurmrnrorusascsdsesmsgsrgdsnsiskslsostessuswsssvtatetgthtibotktltntotrtstttwtyugukuruzvevivowacywofyxhyiyoza"

// codeIndex maps a lowercase ISO-639-1 pair to its catalog index.
// Unlisted pairs are 0 (undefined).
var codeIndex = [26][26]uint8{
	0: {
		0: 2, 'b' - 'a': 1, 'e' - 'a': 12, 'f' - 'a': 3, 'k' - 'a': 4, 'm' - 'a': 6, 'n' - 'a': 8, 'r' - 'a': 7, 's' - 'a': 10, 'v' - 'a': 11, 'y' - 'a': 13, 'z' - 'a': 14,
	},
	'b' - 'a': {
		0: 16, 'e' - 'a': 18, 'g' - 'a': 24, 'h' - 'a': 20, 'i' - 'a': 21, 'm' - 'a': 15, 'n' - 'a': 19, 'o' - 'a': 158, 'r' - 'a': 23, 's' - 'a': 22,
	},
	'c' - 'a': {
		0: 26, 'e' - 'a': 28, 'h' - 'a': 27, 'o' - 'a': 33, 'r' - 'a': 34, 's' - 'a': 36, 'u' - 'a': 119, 'v' - 'a': 31, 'y' - 'a': 176,
	},
	'd' - 'a': {
		0: 37, 'e' - 'a': 51, 'v' - 'a': 38,
	},
	'e' - 'a': {
		'e' - 'a': 43, 'l' - 'a': 52, 'n' - 'a': 40, 'o' - 'a': 41, 's' - 'a': 148, 't' - 'a': 42, 'u' - 'a': 17,
	},
	'f' - 'a': {
		0: 125, 'f' - 'a': 48, 'i' - 'a': 46, 'j' - 'a': 45, 'o' - 'a': 44, 'r' - 'a': 47, 'y' - 'a': 178,
	},
	'g' - 'a': {
		0: 65, 'd' - 'a': 141, 'l' - 'a': 49, 'n' - 'a': 53, 'u' - 'a': 54, 'v' - 'a': 97,
	},
	'h' - 'a': {
		0: 56, 'e' - 'a': 57, 'i' - 'a': 59, 'o' - 'a': 60, 'r' - 'a': 35, 't' - 'a': 55, 'u' - 'a': 61, 'y' - 'a': 9, 'z' - 'a': 58,
	},
	'i' - 'a': {
		0: 62, 'd' - 'a': 63, 'e' - 'a': 64, 'g' - 'a': 66, 'i' - 'a': 115, 'k' - 'a': 67, 'o' - 'a': 68, 's' - 'a': 69, 't' - 'a': 70, 'u' - 'a': 71,
	},
	'j' - 'a': {
		0: 72, 'v' - 'a': 73,
	},
	'k' - 'a': {
		0: 50, 'g' - 'a': 84, 'i' - 'a': 80, 'j' - 'a': 87, 'k' - 'a': 78, 'l' - 'a': 74, 'm' - 'a': 79, 'n' - 'a': 75, 'o' - 'a': 85, 'r' - 'a': 76, 's' - 'a': 77, 'u' - 'a': 86, 'v' - 'a': 83, 'w' - 'a': 32, 'y' - 'a': 82,
	},
	'l' - 'a': {
		0: 88, 'b' - 'a': 89, 'g' - 'a': 90, 'i' - 'a': 91, 'n' - 'a': 92, 'o' - 'a': 93, 't' - 'a': 94, 'u' - 'a': 95, 'v' - 'a': 96,
	},
	'm' - 'a': {
		'g' - 'a': 99, 'h' - 'a': 105, 'i' - 'a': 103, 'k' - 'a': 98, 'l' - 'a': 101, 'n' - 'a': 106, 'r' - 'a': 104, 's' - 'a': 100, 't' - 'a': 102, 'y' - 'a': 25,
	},
	'n' - 'a': {
		0: 107, 'b' - 'a': 109, 'd' - 'a': 110, 'e' - 'a': 111, 'g' - 'a': 112, 'l' - 'a': 39, 'n' - 'a': 113, 'o' - 'a': 114, 'r' - 'a': 116, 'v' - 'a': 108, 'y' - 'a': 29,
	},
	'o' - 'a': {
		'c' - 'a': 117, 'j' - 'a': 118, 'm' - 'a': 120, 'r' - 'a': 121, 's' - 'a': 122,
	},
	'p' - 'a': {
		0: 123, 'i' - 'a': 124, 'l' - 'a': 126, 's' - 'a': 127, 't' - 'a': 128,
	},
	'q' - 'a': {
		'u' - 'a': 129,
	},
	'r' - 'a': {
		'm' - 'a': 130, 'n' - 'a': 131, 'o' - 'a': 132, 'u' - 'a': 133, 'w' - 'a': 81,
	},
	's' - 'a': {
		0: 134, 'c' - 'a': 135, 'd' - 'a': 136, 'e' - 'a': 137, 'g' - 'a': 139, 'i' - 'a': 143, 'k' - 'a': 144, 'l' - 'a': 145, 'm' - 'a': 138, 'n' - 'a': 142, 'o' - 'a': 146, 'q' - 'a': 5, 'r' - 'a': 140, 's' - 'a': 151, 't' - 'a': 147, 'u' - 'a': 149, 'v' - 'a': 152, 'w' - 'a': 150,
	},
	't' - 'a': {
		0: 153, 'e' - 'a': 154, 'g' - 'a': 155, 'h' - 'a': 156, 'i' - 'a': 157, 'k' - 'a': 159, 'l' - 'a': 160, 'n' - 'a': 161, 'o' - 'a': 162, 'r' - 'a': 163, 's' - 'a': 164, 't' - 'a': 165, 'w' - 'a': 166, 'y' - 'a': 167,
	},
	'u' - 'a': {
		'g' - 'a': 168, 'k' - 'a': 169, 'r' - 'a': 170, 'z' - 'a': 171,
	},
	'v' - 'a': {
		'e' - 'a': 172, 'i' - 'a': 173, 'o' - 'a': 174,
	},
	'w' - 'a': {
		0: 175, 'o' - 'a': 177,
	},
	'x' - 'a': {
		'h' - 'a': 179,
	},
	'y' - 'a': {
		'i' - 'a': 180, 'o' - 'a': 181,
	},
	'z' - 'a': {
		0: 182, 'h' - 'a': 30,
	},
}

func asciiLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func lookupID(b0, b1 byte) ID {
	b0, b1 = asciiLower(b0), asciiLower(b1)
	if b0 >= 'a' && b0 <= 'z' && b1 >= 'a' && b1 <= 'z' {
		return ID(codeIndex[b0-'a'][b1-'a'])
	}
	return UndefinedID
}

func iso2(id ID) string {
	if uint(id) >= uint(len(languageISO2)/2) {
		return UndefinedISO2
	}
	i := uint(id) * 2
	return languageISO2[i : i+2]
}

func lookup(b0, b1 byte) *Language {
	return &languages[lookupID(b0, b1)]
}

// GetLanguageByCodeString resolves a BCP-47 / ISO-639-1 string (first two letters).
// Unknown and short values return Languages[0] (never nil).
func GetLanguageByCodeString(code string) *Language {
	if len(code) < 2 {
		return &languages[0]
	}
	return lookup(code[0], code[1])
}
