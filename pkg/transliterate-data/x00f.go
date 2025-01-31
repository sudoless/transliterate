package transliterate_data

var x00f = []string{
	"AUM",     // 0x00
	"",        // 0x01
	"",        // 0x02
	"",        // 0x03
	"",        // 0x04
	"",        // 0x05
	"",        // 0x06
	"",        // 0x07
	" // ",    // 0x08
	" * ",     // 0x09
	"",        // 0x0a
	"-",       // 0x0b
	" / ",     // 0x0c
	" / ",     // 0x0d
	" // ",    // 0x0e
	" -/ ",    // 0x0f
	" +/ ",    // 0x10
	" X/ ",    // 0x11
	" /XX/ ",  // 0x12
	" /X/ ",   // 0x13
	", ",      // 0x14
	"",        // 0x15
	"",        // 0x16
	"",        // 0x17
	"",        // 0x18
	"",        // 0x19
	"",        // 0x1a
	"",        // 0x1b
	"",        // 0x1c
	"",        // 0x1d
	"",        // 0x1e
	"",        // 0x1f
	"0",       // 0x20
	"1",       // 0x21
	"2",       // 0x22
	"3",       // 0x23
	"4",       // 0x24
	"5",       // 0x25
	"6",       // 0x26
	"7",       // 0x27
	"8",       // 0x28
	"9",       // 0x29
	".5",      // 0x2a
	"1.5",     // 0x2b
	"2.5",     // 0x2c
	"3.5",     // 0x2d
	"4.5",     // 0x2e
	"5.5",     // 0x2f
	"6.5",     // 0x30
	"7.5",     // 0x31
	"8.5",     // 0x32
	"-.5",     // 0x33
	"+",       // 0x34
	"*",       // 0x35
	"^",       // 0x36
	"_",       // 0x37
	"",        // 0x38
	"~",       // 0x39
	"[?]",     // 0x3a
	"]",       // 0x3b
	"[[",      // 0x3c
	"]]",      // 0x3d
	"",        // 0x3e
	"",        // 0x3f
	"k",       // 0x40
	"kh",      // 0x41
	"g",       // 0x42
	"gh",      // 0x43
	"ng",      // 0x44
	"c",       // 0x45
	"ch",      // 0x46
	"j",       // 0x47
	"[?]",     // 0x48
	"ny",      // 0x49
	"tt",      // 0x4a
	"tth",     // 0x4b
	"dd",      // 0x4c
	"ddh",     // 0x4d
	"nn",      // 0x4e
	"t",       // 0x4f
	"th",      // 0x50
	"d",       // 0x51
	"dh",      // 0x52
	"n",       // 0x53
	"p",       // 0x54
	"ph",      // 0x55
	"b",       // 0x56
	"bh",      // 0x57
	"m",       // 0x58
	"ts",      // 0x59
	"tsh",     // 0x5a
	"dz",      // 0x5b
	"dzh",     // 0x5c
	"w",       // 0x5d
	"zh",      // 0x5e
	"z",       // 0x5f
	"'",       // 0x60
	"y",       // 0x61
	"r",       // 0x62
	"l",       // 0x63
	"sh",      // 0x64
	"ssh",     // 0x65
	"s",       // 0x66
	"h",       // 0x67
	"a",       // 0x68
	"kss",     // 0x69
	"r",       // 0x6a
	"[?]",     // 0x6b
	"[?]",     // 0x6c
	"[?]",     // 0x6d
	"[?]",     // 0x6e
	"[?]",     // 0x6f
	"[?]",     // 0x70
	"aa",      // 0x71
	"i",       // 0x72
	"ii",      // 0x73
	"u",       // 0x74
	"uu",      // 0x75
	"R",       // 0x76
	"RR",      // 0x77
	"L",       // 0x78
	"LL",      // 0x79
	"e",       // 0x7a
	"ee",      // 0x7b
	"o",       // 0x7c
	"oo",      // 0x7d
	"M",       // 0x7e
	"H",       // 0x7f
	"i",       // 0x80
	"ii",      // 0x81
	"",        // 0x82
	"",        // 0x83
	"",        // 0x84
	"",        // 0x85
	"",        // 0x86
	"",        // 0x87
	"",        // 0x88
	"",        // 0x89
	"",        // 0x8a
	"",        // 0x8b
	"[?]",     // 0x8c
	"[?]",     // 0x8d
	"[?]",     // 0x8e
	"[?]",     // 0x8f
	"k",       // 0x90
	"kh",      // 0x91
	"g",       // 0x92
	"gh",      // 0x93
	"ng",      // 0x94
	"c",       // 0x95
	"ch",      // 0x96
	"j",       // 0x97
	"[?]",     // 0x98
	"ny",      // 0x99
	"tt",      // 0x9a
	"tth",     // 0x9b
	"dd",      // 0x9c
	"ddh",     // 0x9d
	"nn",      // 0x9e
	"t",       // 0x9f
	"th",      // 0xa0
	"d",       // 0xa1
	"dh",      // 0xa2
	"n",       // 0xa3
	"p",       // 0xa4
	"ph",      // 0xa5
	"b",       // 0xa6
	"bh",      // 0xa7
	"m",       // 0xa8
	"ts",      // 0xa9
	"tsh",     // 0xaa
	"dz",      // 0xab
	"dzh",     // 0xac
	"w",       // 0xad
	"zh",      // 0xae
	"z",       // 0xaf
	"'",       // 0xb0
	"y",       // 0xb1
	"r",       // 0xb2
	"l",       // 0xb3
	"sh",      // 0xb4
	"ss",      // 0xb5
	"s",       // 0xb6
	"h",       // 0xb7
	"a",       // 0xb8
	"kss",     // 0xb9
	"w",       // 0xba
	"y",       // 0xbb
	"r",       // 0xbc
	"[?]",     // 0xbd
	"X",       // 0xbe
	" :X: ",   // 0xbf
	" /O/ ",   // 0xc0
	" /o/ ",   // 0xc1
	" \\o\\ ", // 0xc2
	" (O) ",   // 0xc3
	"",        // 0xc4
	"",        // 0xc5
	"",        // 0xc6
	"",        // 0xc7
	"",        // 0xc8
	"",        // 0xc9
	"",        // 0xca
	"",        // 0xcb
	"",        // 0xcc
	"[?]",     // 0xcd
	"[?]",     // 0xce
	"",        // 0xcf
	"[?]",     // 0xd0
	"[?]",     // 0xd1
	"[?]",     // 0xd2
	"[?]",     // 0xd3
	"[?]",     // 0xd4
	"[?]",     // 0xd5
	"[?]",     // 0xd6
	"[?]",     // 0xd7
	"[?]",     // 0xd8
	"[?]",     // 0xd9
	"[?]",     // 0xda
	"[?]",     // 0xdb
	"[?]",     // 0xdc
	"[?]",     // 0xdd
	"[?]",     // 0xde
	"[?]",     // 0xdf
	"[?]",     // 0xe0
	"[?]",     // 0xe1
	"[?]",     // 0xe2
	"[?]",     // 0xe3
	"[?]",     // 0xe4
	"[?]",     // 0xe5
	"[?]",     // 0xe6
	"[?]",     // 0xe7
	"[?]",     // 0xe8
	"[?]",     // 0xe9
	"[?]",     // 0xea
	"[?]",     // 0xeb
	"[?]",     // 0xec
	"[?]",     // 0xed
	"[?]",     // 0xee
	"[?]",     // 0xef
	"[?]",     // 0xf0
	"[?]",     // 0xf1
	"[?]",     // 0xf2
	"[?]",     // 0xf3
	"[?]",     // 0xf4
	"[?]",     // 0xf5
	"[?]",     // 0xf6
	"[?]",     // 0xf7
	"[?]",     // 0xf8
	"[?]",     // 0xf9
	"[?]",     // 0xfa
	"[?]",     // 0xfb
	"[?]",     // 0xfc
	"[?]",     // 0xfd
	"[?]",     // 0xfe
}
