package main

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"math"
	"strconv"
)

var positionAdjustPatternTable = [][]int{
	{},
	{}, // version 1
	{6, 18},
	{6, 22},
	{6, 26},
	{6, 30}, // version 5
	{6, 34},
	{6, 22, 38},
	{6, 24, 42},
	{6, 26, 46},
	{6, 28, 50}, // version 10
	{6, 30, 54},
	{6, 32, 58},
	{6, 34, 62},
	{6, 26, 46, 66},
	{6, 26, 48, 70},
	{6, 26, 50, 74},
	{6, 30, 54, 78},
	{6, 30, 56, 82},
	{6, 30, 58, 86},
	{6, 34, 62, 90}, // version 20
	{6, 28, 50, 72, 94},
	{6, 26, 50, 74, 98},
	{6, 30, 54, 78, 102},
	{6, 28, 54, 80, 106},
	{6, 32, 58, 84, 110},
	{6, 30, 58, 86, 114},
	{6, 34, 62, 90, 118},
	{6, 26, 50, 74, 98, 122},
	{6, 30, 54, 78, 102, 126},
	{6, 26, 52, 78, 104, 130}, // version 30
	{6, 30, 56, 82, 108, 134},
	{6, 34, 60, 86, 112, 138},
	{6, 30, 58, 86, 114, 142},
	{6, 34, 62, 90, 118, 146},
	{6, 30, 54, 78, 102, 126, 150},
	{6, 24, 50, 76, 102, 128, 154},
	{6, 28, 54, 80, 106, 132, 158},
	{6, 32, 58, 84, 110, 136, 162},
	{6, 26, 54, 82, 110, 138, 166},
	{6, 30, 58, 86, 114, 142, 170}} // version 40

type ECLevel int

const (
	ECLevelL ECLevel = 1 << iota
	ECLevelM ECLevel = 1 << iota
	ECLevelQ ECLevel = 1 << iota
	ECLevelH ECLevel = 1 << iota
)

var typeInformationTable = map[ECLevel][]int{
	ECLevelL: {0x77c4, 0x72f3, 0x7daa, 0x789d, 0x662f, 0x6318, 0x6c41, 0x6976},
	ECLevelM: {0x5412, 0x5125, 0x5e7c, 0x5b4b, 0x45f9, 0x40ce, 0x4f97, 0x4aa0},
	ECLevelQ: {0x355f, 0x3068, 0x3f31, 0x3a06, 0x24b4, 0x2183, 0x2eda, 0x2bed},
	ECLevelH: {0x1689, 0x13be, 0x1ce7, 0x19d0, 0x0762, 0x0255, 0x0d0c, 0x083b},
}

var versionInformationTable = []int{
	0x07c94, 0x085bc, 0x09a99, 0x0a4d3, 0x0bbf6, 0x0c762, 0x0d847, 0x0e60d,
	0x0f928, 0x10b78, 0x1145d, 0x12a17, 0x13532, 0x149a6, 0x15683, 0x168c9,
	0x177ec, 0x18ec4, 0x191e1, 0x1Afab, 0x1b08e, 0x1cc1a, 0x1d33f, 0x1ef75,
	0x1f250, 0x209d5, 0x216f0, 0x228ba, 0x2379f, 0x24b0b, 0x2542e, 0x26a64,
	0x27541, 0x28c69,
}

var errorCorrectionTable = []map[ECLevel][]int{
	// ec code word per block
	// block 1 count
	// block 1 data code words
	// block 2 count
	// block 2 data code words
	{},
	{ECLevelL: {7, 1, 19, 0, 0}, ECLevelM: {10, 1, 16, 0, 0}, ECLevelQ: {13, 1, 13, 0, 0}, ECLevelH: {17, 1, 9, 0, 0}}, // version 1
	{ECLevelL: {10, 1, 34, 0, 0}, ECLevelM: {16, 1, 28, 0, 0}, ECLevelQ: {22, 1, 22, 0, 0}, ECLevelH: {28, 1, 16, 0, 0}},
	{ECLevelL: {15, 1, 55, 0, 0}, ECLevelM: {26, 1, 44, 0, 0}, ECLevelQ: {18, 2, 17, 0, 0}, ECLevelH: {22, 2, 13, 0, 0}},
	{ECLevelL: {20, 1, 80, 0, 0}, ECLevelM: {18, 2, 32, 0, 0}, ECLevelQ: {26, 2, 24, 0, 0}, ECLevelH: {16, 4, 9, 0, 0}},
	{ECLevelL: {26, 1, 108, 0, 0}, ECLevelM: {24, 2, 43, 0, 0}, ECLevelQ: {18, 2, 15, 2, 16}, ECLevelH: {22, 2, 11, 2, 12}}, // version 5
	{ECLevelL: {18, 2, 68, 0, 0}, ECLevelM: {16, 4, 27, 0, 0}, ECLevelQ: {24, 4, 19, 0, 0}, ECLevelH: {28, 4, 15, 0, 0}},
	{ECLevelL: {20, 2, 78, 0, 0}, ECLevelM: {18, 4, 31, 0, 0}, ECLevelQ: {18, 2, 14, 4, 15}, ECLevelH: {26, 4, 13, 1, 14}},
	{ECLevelL: {24, 2, 97, 0, 0}, ECLevelM: {22, 2, 38, 2, 39}, ECLevelQ: {22, 4, 18, 2, 19}, ECLevelH: {26, 4, 14, 2, 15}},
	{ECLevelL: {30, 2, 116, 0, 0}, ECLevelM: {22, 3, 36, 2, 37}, ECLevelQ: {20, 4, 16, 4, 17}, ECLevelH: {24, 4, 12, 4, 13}},
	{ECLevelL: {18, 2, 68, 2, 69}, ECLevelM: {26, 4, 43, 1, 44}, ECLevelQ: {24, 6, 19, 2, 20}, ECLevelH: {28, 6, 15, 2, 16}}, // version 10
	{ECLevelL: {20, 4, 81, 0, 0}, ECLevelM: {31, 1, 50, 4, 51}, ECLevelQ: {28, 4, 22, 4, 23}, ECLevelH: {24, 3, 12, 8, 13}},
	{ECLevelL: {24, 2, 92, 2, 93}, ECLevelM: {22, 6, 36, 2, 37}, ECLevelQ: {26, 4, 20, 6, 21}, ECLevelH: {28, 7, 14, 4, 15}},
	{ECLevelL: {26, 4, 107, 0, 0}, ECLevelM: {22, 8, 37, 1, 38}, ECLevelQ: {24, 8, 20, 4, 21}, ECLevelH: {22, 12, 11, 4, 12}},
	{ECLevelL: {30, 3, 115, 1, 116}, ECLevelM: {24, 4, 40, 5, 41}, ECLevelQ: {20, 11, 16, 5, 17}, ECLevelH: {24, 11, 12, 5, 13}},
	{ECLevelL: {22, 5, 87, 1, 88}, ECLevelM: {24, 5, 41, 5, 42}, ECLevelQ: {30, 5, 24, 7, 25}, ECLevelH: {30, 5, 24, 7, 25}},
	{ECLevelL: {24, 5, 98, 1, 99}, ECLevelM: {28, 7, 45, 3, 46}, ECLevelQ: {24, 15, 19, 2, 20}, ECLevelH: {30, 3, 15, 13, 16}},
	{ECLevelL: {28, 1, 107, 5, 108}, ECLevelM: {28, 10, 46, 1, 47}, ECLevelQ: {28, 1, 22, 15, 23}, ECLevelH: {28, 2, 14, 17, 15}},
	{ECLevelL: {30, 5, 120, 1, 121}, ECLevelM: {26, 9, 43, 4, 44}, ECLevelQ: {28, 17, 22, 1, 23}, ECLevelH: {28, 2, 14, 19, 15}},
	{ECLevelL: {28, 3, 113, 4, 114}, ECLevelM: {26, 3, 44, 11, 45}, ECLevelQ: {26, 17, 21, 4, 22}, ECLevelH: {26, 9, 13, 16, 14}},
	{ECLevelL: {28, 3, 107, 6, 108}, ECLevelM: {26, 3, 41, 13, 42}, ECLevelQ: {30, 15, 24, 5, 25}, ECLevelH: {28, 15, 15, 10, 16}}, // version 20
	{ECLevelL: {28, 4, 116, 4, 117}, ECLevelM: {26, 17, 42, 0, 0}, ECLevelQ: {28, 17, 22, 6, 23}, ECLevelH: {30, 19, 16, 6, 17}},
	{ECLevelL: {28, 2, 111, 7, 112}, ECLevelM: {28, 17, 46, 0, 0}, ECLevelQ: {30, 7, 24, 16, 25}, ECLevelH: {24, 34, 13, 0, 0}},
	{ECLevelL: {30, 4, 121, 5, 122}, ECLevelM: {28, 4, 47, 14, 48}, ECLevelQ: {30, 11, 24, 14, 25}, ECLevelH: {30, 16, 15, 14, 16}},
	{ECLevelL: {30, 6, 117, 4, 118}, ECLevelM: {28, 6, 45, 14, 46}, ECLevelQ: {30, 11, 24, 16, 25}, ECLevelH: {30, 30, 16, 2, 17}},
	{ECLevelL: {26, 8, 106, 4, 107}, ECLevelM: {28, 8, 47, 13, 48}, ECLevelQ: {230, 7, 24, 22, 25}, ECLevelH: {30, 22, 15, 13, 16}},
	{ECLevelL: {28, 10, 114, 2, 115}, ECLevelM: {28, 19, 46, 4, 47}, ECLevelQ: {28, 28, 22, 6, 23}, ECLevelH: {30, 33, 16, 4, 17}},
	{ECLevelL: {30, 8, 122, 4, 123}, ECLevelM: {28, 22, 45, 3, 46}, ECLevelQ: {30, 8, 23, 26, 24}, ECLevelH: {30, 12, 15, 28, 16}},
	{ECLevelL: {30, 3, 117, 10, 118}, ECLevelM: {28, 3, 45, 23, 46}, ECLevelQ: {30, 4, 24, 31, 25}, ECLevelH: {30, 11, 15, 31, 16}},
	{ECLevelL: {30, 7, 116, 7, 117}, ECLevelM: {28, 21, 45, 7, 46}, ECLevelQ: {30, 1, 23, 37, 24}, ECLevelH: {30, 19, 15, 26, 16}},
	{ECLevelL: {30, 5, 115, 10, 116}, ECLevelM: {28, 19, 47, 10, 48}, ECLevelQ: {30, 15, 24, 25, 25}, ECLevelH: {30, 23, 15, 25, 16}}, // version 30
	{ECLevelL: {30, 13, 115, 3, 116}, ECLevelM: {28, 2, 46, 29, 47}, ECLevelQ: {30, 42, 24, 1, 25}, ECLevelH: {30, 23, 15, 28, 16}},
	{ECLevelL: {30, 17, 115, 0, 0}, ECLevelM: {28, 10, 46, 23, 47}, ECLevelQ: {30, 10, 24, 35, 25}, ECLevelH: {30, 19, 15, 35, 16}},
	{ECLevelL: {30, 17, 115, 1, 116}, ECLevelM: {28, 14, 46, 21, 47}, ECLevelQ: {30, 29, 24, 19, 25}, ECLevelH: {30, 11, 15, 46, 16}},
	{ECLevelL: {30, 13, 115, 6, 116}, ECLevelM: {28, 14, 46, 13, 47}, ECLevelQ: {30, 44, 24, 7, 25}, ECLevelH: {30, 59, 16, 1, 17}},
	{ECLevelL: {30, 12, 121, 7, 122}, ECLevelM: {28, 12, 47, 26, 48}, ECLevelQ: {30, 39, 24, 14, 25}, ECLevelH: {30, 22, 15, 41, 16}},
	{ECLevelL: {30, 6, 121, 14, 122}, ECLevelM: {28, 6, 47, 34, 48}, ECLevelQ: {30, 46, 24, 10, 25}, ECLevelH: {30, 2, 15, 64, 16}},
	{ECLevelL: {30, 17, 122, 4, 123}, ECLevelM: {28, 29, 46, 14, 47}, ECLevelQ: {30, 49, 24, 10, 25}, ECLevelH: {30, 24, 15, 46, 16}},
	{ECLevelL: {30, 4, 122, 18, 123}, ECLevelM: {28, 13, 46, 32, 47}, ECLevelQ: {30, 48, 24, 14, 25}, ECLevelH: {30, 42, 15, 32, 16}},
	{ECLevelL: {30, 20, 117, 4, 118}, ECLevelM: {28, 40, 47, 7, 48}, ECLevelQ: {30, 43, 24, 22, 25}, ECLevelH: {30, 10, 15, 67, 16}},
	{ECLevelL: {30, 19, 118, 6, 119}, ECLevelM: {28, 18, 47, 31, 48}, ECLevelQ: {30, 34, 24, 34, 25}, ECLevelH: {30, 20, 15, 61, 16}}, // version 40
}

// Galois field antilog to log convertion table
var int2exp = []int{
	0xff, 0x00, 0x01, 0x19, 0x02, 0x32, 0x1a, 0xc6, 0x03, 0xdf, 0x33, 0xee, 0x1b, 0x68, 0xc7, 0x4b,
	0x04, 0x64, 0xe0, 0x0e, 0x34, 0x8d, 0xef, 0x81, 0x1c, 0xc1, 0x69, 0xf8, 0xc8, 0x08, 0x4c, 0x71,
	0x05, 0x8a, 0x65, 0x2f, 0xe1, 0x24, 0x0f, 0x21, 0x35, 0x93, 0x8e, 0xda, 0xf0, 0x12, 0x82, 0x45,
	0x1d, 0xb5, 0xc2, 0x7d, 0x6a, 0x27, 0xf9, 0xb9, 0xc9, 0x9a, 0x09, 0x78, 0x4d, 0xe4, 0x72, 0xa6,
	0x06, 0xbf, 0x8b, 0x62, 0x66, 0xdd, 0x30, 0xfd, 0xe2, 0x98, 0x25, 0xb3, 0x10, 0x91, 0x22, 0x88,
	0x36, 0xd0, 0x94, 0xce, 0x8f, 0x96, 0xdb, 0xbd, 0xf1, 0xd2, 0x13, 0x5c, 0x83, 0x38, 0x46, 0x40,
	0x1e, 0x42, 0xb6, 0xa3, 0xc3, 0x48, 0x7e, 0x6e, 0x6b, 0x3a, 0x28, 0x54, 0xfa, 0x85, 0xba, 0x3d,
	0xca, 0x5e, 0x9b, 0x9f, 0x0a, 0x15, 0x79, 0x2b, 0x4e, 0xd4, 0xe5, 0xac, 0x73, 0xf3, 0xa7, 0x57,
	0x07, 0x70, 0xc0, 0xf7, 0x8c, 0x80, 0x63, 0x0d, 0x67, 0x4a, 0xde, 0xed, 0x31, 0xc5, 0xfe, 0x18,
	0xe3, 0xa5, 0x99, 0x77, 0x26, 0xb8, 0xb4, 0x7c, 0x11, 0x44, 0x92, 0xd9, 0x23, 0x20, 0x89, 0x2e,
	0x37, 0x3f, 0xd1, 0x5b, 0x95, 0xbc, 0xcf, 0xcd, 0x90, 0x87, 0x97, 0xb2, 0xdc, 0xfc, 0xbe, 0x61,
	0xf2, 0x56, 0xd3, 0xab, 0x14, 0x2a, 0x5d, 0x9e, 0x84, 0x3c, 0x39, 0x53, 0x47, 0x6d, 0x41, 0xa2,
	0x1f, 0x2d, 0x43, 0xd8, 0xb7, 0x7b, 0xa4, 0x76, 0xc4, 0x17, 0x49, 0xec, 0x7f, 0x0c, 0x6f, 0xf6,
	0x6c, 0xa1, 0x3b, 0x52, 0x29, 0x9d, 0x55, 0xaa, 0xfb, 0x60, 0x86, 0xb1, 0xbb, 0xcc, 0x3e, 0x5a,
	0xcb, 0x59, 0x5f, 0xb0, 0x9c, 0xa9, 0xa0, 0x51, 0x0b, 0xf5, 0x16, 0xeb, 0x7a, 0x75, 0x2c, 0xd7,
	0x4f, 0xae, 0xd5, 0xe9, 0xe6, 0xe7, 0xad, 0xe8, 0x74, 0xd6, 0xf4, 0xea, 0xa8, 0x50, 0x58, 0xaf,
}

// Galois field log to antilog convertion table
var exp2int = []int{
	0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1d, 0x3a, 0x74, 0xe8, 0xcd, 0x87, 0x13, 0x26,
	0x4c, 0x98, 0x2d, 0x5a, 0xb4, 0x75, 0xea, 0xc9, 0x8f, 0x03, 0x06, 0x0c, 0x18, 0x30, 0x60, 0xc0,
	0x9d, 0x27, 0x4e, 0x9c, 0x25, 0x4a, 0x94, 0x35, 0x6a, 0xd4, 0xb5, 0x77, 0xee, 0xc1, 0x9f, 0x23,
	0x46, 0x8c, 0x05, 0x0a, 0x14, 0x28, 0x50, 0xa0, 0x5d, 0xba, 0x69, 0xd2, 0xb9, 0x6f, 0xde, 0xa1,
	0x5f, 0xbe, 0x61, 0xc2, 0x99, 0x2f, 0x5e, 0xbc, 0x65, 0xca, 0x89, 0x0f, 0x1e, 0x3c, 0x78, 0xf0,
	0xfd, 0xe7, 0xd3, 0xbb, 0x6b, 0xd6, 0xb1, 0x7f, 0xfe, 0xe1, 0xdf, 0xa3, 0x5b, 0xb6, 0x71, 0xe2,
	0xd9, 0xaf, 0x43, 0x86, 0x11, 0x22, 0x44, 0x88, 0x0d, 0x1a, 0x34, 0x68, 0xd0, 0xbd, 0x67, 0xce,
	0x81, 0x1f, 0x3e, 0x7c, 0xf8, 0xed, 0xc7, 0x93, 0x3b, 0x76, 0xec, 0xc5, 0x97, 0x33, 0x66, 0xcc,
	0x85, 0x17, 0x2e, 0x5c, 0xb8, 0x6d, 0xda, 0xa9, 0x4f, 0x9e, 0x21, 0x42, 0x84, 0x15, 0x2a, 0x54,
	0xa8, 0x4d, 0x9a, 0x29, 0x52, 0xa4, 0x55, 0xaa, 0x49, 0x92, 0x39, 0x72, 0xe4, 0xd5, 0xb7, 0x73,
	0xe6, 0xd1, 0xbf, 0x63, 0xc6, 0x91, 0x3f, 0x7e, 0xfc, 0xe5, 0xd7, 0xb3, 0x7b, 0xf6, 0xf1, 0xff,
	0xe3, 0xdb, 0xab, 0x4b, 0x96, 0x31, 0x62, 0xc4, 0x95, 0x37, 0x6e, 0xdc, 0xa5, 0x57, 0xae, 0x41,
	0x82, 0x19, 0x32, 0x64, 0xc8, 0x8d, 0x07, 0x0e, 0x1c, 0x38, 0x70, 0xe0, 0xdd, 0xa7, 0x53, 0xa6,
	0x51, 0xa2, 0x59, 0xb2, 0x79, 0xf2, 0xf9, 0xef, 0xc3, 0x9b, 0x2b, 0x56, 0xac, 0x45, 0x8a, 0x09,
	0x12, 0x24, 0x48, 0x90, 0x3d, 0x7a, 0xf4, 0xf5, 0xf7, 0xf3, 0xfb, 0xeb, 0xcb, 0x8b, 0x0b, 0x16,
	0x2c, 0x58, 0xb0, 0x7d, 0xfa, 0xe9, 0xcf, 0x83, 0x1b, 0x36, 0x6c, 0xd8, 0xad, 0x47, 0x8e, 0x00,
}

var alphanumTable = map[byte]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
	'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
	'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, ' ': 36, '$': 37, '%': 38, '*': 39,
	'+': 40, '-': 41, '.': 42, '/': 43, ':': 44,
}

// Errors introduced by this package
var (
	errInvalidVersion        = errors.New("goqr: invalid qrcode version (1-40)")
	errInvalidLevel          = errors.New("goqr: invalid qrcode error correctionlevel (ECLevelL,M,Q,H)")
	errInvalidModuleSize     = errors.New("goqr: invalid qrcode module size (>=1)")
	errInvalidQuietZoneWidth = errors.New("goqr: invalid qrcode quiet zone width (>=0)")
	errInvalidDataSize       = errors.New("goqr: invalid data size")
)

type Qrcode struct {
	Version        int     // qrcode version 1-40 (0:auto)
	Level          ECLevel // error correction Level (0:auto)
	mode           string  //
	module         [][]int // 2-D matrix [row][col] ( +-1 : data, +-2 : pattern )
	data           string  // data to encode
	ModuleSize     int     // module size (default 1)
	QuietZoneWidth int     // quiet zone width
	encodedData    []byte  // encoded data
	penalty        []int   // calculated mask penalty (0-7)
}

func Encode(data string, version int, level ECLevel) (image.Image, error) {

	qr := new(Qrcode)
	qr.data = data
	qr.Version = version
	qr.Level = level
	qr.ModuleSize = 1
	qr.QuietZoneWidth = 4

	return qr.Encode()
}

// Module count on a side
func (qr *Qrcode) len() int { return qr.Version*4 + (7+1)*2 + 1 }

func (qr *Qrcode) Encode() (image.Image, error) {

	// check version
	if qr.Version < 0 || 40 < qr.Version {
		return nil, errInvalidVersion
	}

	// check level
	if qr.Level != 0 && len(errorCorrectionTable[1][qr.Level]) == 0 {
		return nil, errInvalidLevel
	}

	// check module size
	if qr.ModuleSize < 1 {
		return nil, errInvalidModuleSize
	}

	// check quiet zone width
	if qr.QuietZoneWidth < 0 {
		return nil, errInvalidQuietZoneWidth
	}

	// set encoding mode (kanji mode not supported)
	qr.selectMode()

	// set version & level
	qr.selectVersionLevel()
	if qr.Version == 0 || qr.Level == 0 {
		return nil, errInvalidDataSize
	}

	// initialize qrcode
	for i := 0; i < qr.len(); i++ {
		s := make([]int, qr.len())
		qr.module = append(qr.module, s)
	}

	// place pattern to qr.module
	qr.placePatterns()

	// encode data -> qr.encodedData
	qr.encodeData()

	// place encoded data to qr.module
	qr.mapData()

	// mask
	qr.maskData()

	// [][]int to [][]bool
	module := make([][]bool, 0)
	for _, row := range qr.module {
		r := make([]bool, qr.len())
		for c, cell := range row {
			if cell > 0 {
				r[c] = true
			} else {
				r[c] = false
			}
		}
		module = append(module, r)
	}

	// quiet zone
	for i := 0; i < qr.QuietZoneWidth; i++ {
		module = append(module, make([]bool, qr.len()))
		module = append([][]bool{make([]bool, qr.len())}, module...)
	}
	for i, row := range module {
		row = append(make([]bool, qr.QuietZoneWidth), row...)
		row = append(row, make([]bool, qr.QuietZoneWidth)...)
		module[i] = row
	}

	// module size
	if qr.ModuleSize > 1 {
		l := len(module) * qr.ModuleSize
		m := make([][]bool, l)

		for i := 0; i < l; i++ {
			m[i] = make([]bool, l)
		}

		for r, row := range module {
			for c, cell := range row {
				if !cell {
					continue
				}
				for x := 0; x < qr.ModuleSize; x++ {
					for y := 0; y < qr.ModuleSize; y++ {
						m[r*qr.ModuleSize+x][c*qr.ModuleSize+y] = true
					}
				}
			}
		}
		module = m
	}

	rgba := image.NewRGBA(image.Rect(0, 0, len(module), len(module)))

	for r, row := range module {
		for c, cell := range row {
			if cell {
				rgba.SetRGBA(c, r, color.RGBA{0, 0, 0, 255})
			} else {
				rgba.SetRGBA(c, r, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	return rgba, nil
}

// qrcode version information table
func (qr *Qrcode) table() []int         { return errorCorrectionTable[qr.Version][qr.Level] }
func (qr *Qrcode) ecCodeWords() int     { return qr.table()[0] }
func (qr *Qrcode) dataCodeWords() []int { return []int{qr.table()[2], qr.table()[4]} }
func (qr *Qrcode) blkCount() []int      { return []int{qr.table()[1], qr.table()[3]} }

// total code words
func (qr *Qrcode) totalCodeWords() int {
	table := qr.table()
	return (table[0]+table[2])*table[1] + (table[0]+table[4])*table[3]
}

// total data code words
func (qr *Qrcode) totalDataCodeWords() int {
	table := qr.table()
	return table[1]*table[2] + table[3]*table[4]
}

// total data code bits
func (qr *Qrcode) totalDataCodeBits() int {
	return qr.totalDataCodeWords() * 8
}

func (qr *Qrcode) selectMode() {
	qr.mode = "numeric"
	for _, c := range qr.data {
		if c >= '0' && c <= '9' {
			continue
		}
		if alphanumTable[byte(c)] != 0 {
			qr.mode = "alphanum"
			continue
		}
		qr.mode = "8bitbyte"
		return
	}
}

func (qr *Qrcode) selectVersionLevel() {
	firstVer, lastVer := 1, 40
	errorCorrectionLevel := []ECLevel{ECLevelH, ECLevelQ, ECLevelM, ECLevelL}

	if qr.Version > 0 {
		firstVer, lastVer = qr.Version, qr.Version
	}

	if qr.Level > 0 {
		errorCorrectionLevel = []ECLevel{qr.Level}
	}

	for i := firstVer; i <= lastVer; i++ {
		for _, j := range errorCorrectionLevel {
			if maxDataSize(i, j, qr.mode) >= len(qr.data) {
				qr.Version, qr.Level = i, j
				return
			}
		}
	}
	qr.Version, qr.Level = 0, 0
}

func (qr *Qrcode) setTypeBits(mask int) {
	qr.module = setTypeBits(qr.module, qr.Level, mask)
}

func (qr *Qrcode) placePatterns() {

	isBlank := func(r, c int) bool { return qr.module[r][c] == 0 }

	// finder pattern
	for _, p := range [3][2]int{{0, 0}, {qr.len() - 7, 0}, {0, qr.len() - 7}} {
		for r := -1; r <= 7; r++ {
			if p[0]+r < 0 || qr.len() <= p[0]+r {
				continue
			}
			for c := -1; c <= 7; c++ {
				if p[1]+c < 0 || qr.len() <= p[1]+c {
					continue
				}
				if (0 <= r && r < 7 && (c == 0 || c == 6)) ||
					(0 <= c && c < 7 && (r == 0 || r == 6)) ||
					(2 <= r && r < 5 && 2 <= c && c < 5) {
					qr.module[p[0]+r][p[1]+c] = 2
				} else {
					qr.module[p[0]+r][p[1]+c] = -2
				}
			}
		}
	}

	// alignment pattern
	pat := positionAdjustPatternTable[qr.Version]

	for i := 0; i < len(pat); i++ {
		for j := 0; j < len(pat); j++ {
			row, col := pat[i], pat[j]
			if !isBlank(row, col) {
				continue
			}

			for r := -2; r <= 2; r++ {
				for c := -2; c <= 2; c++ {
					if r == -2 || r == 2 || c == -2 || c == 2 || (r == 0 && c == 0) {
						qr.module[row+r][col+c] = 2
					} else {
						qr.module[row+r][col+c] = -2
					}
				}
			}
		}
	}

	// timing pattern
	for i := 0; i < qr.len(); i++ {
		if !isBlank(i, 6) {
			continue
		}
		if i%2 == 0 {
			qr.module[i][6] = -2
		} else {
			qr.module[i][6] = 2
		}
	}
	for i := 0; i < qr.len(); i++ {
		if !isBlank(6, i) {
			continue
		}
		if i%2 == 0 {
			qr.module[6][i] = -2
		} else {
			qr.module[6][i] = 2
		}
	}

	// version information
	if qr.Version >= 7 {
		versionBits := versionInformationTable[qr.Version-7]

		for j := 0; j < 6; j++ {
			for k := qr.len() - 11; k < qr.len()-8; k++ {
				if (versionBits & 1) > 0 {
					qr.module[k][j] = 2
					qr.module[j][k] = 2
				} else {
					qr.module[k][j] = -2
					qr.module[j][k] = -2
				}
				versionBits = versionBits >> 1
			}
		}
	}

	// single bit
	qr.module[qr.len()-8][8] = 2

	// dummy format infomation (mask v0)
	qr.setTypeBits(0)

}

func (qr *Qrcode) mapData() {

	// byte slice to bit slice
	bitbuf := new(bitBuffer)
	for _, b := range qr.encodedData {
		bitbuf.appendByte(b)
	}

	// reminder bits
	switch qr.Version {
	case 2, 3, 4, 5, 6:
		bitbuf.append(0, 7)
	case 14, 15, 16, 17, 18, 19, 20, 28, 29, 30, 31, 32, 33, 34:
		bitbuf.append(0, 3)
	case 21, 22, 23, 24, 25, 26, 27:
		bitbuf.append(0, 4)
	}

	r := qr.len() - 1 // row
	c := qr.len() - 1 // column
	v := 1            // direction 1:up, -1:down
	h := 1            // 1:right, -1:left
	i := 0            // index

	for {
		if qr.module[r][c] == 0 {
			if bitbuf.get(i) == true {
				qr.module[r][c] = 1
			} else {
				qr.module[r][c] = -1
			}
			i++
		}
		if i >= bitbuf.len() {
			break
		}
		if c == 6 { // skip
			c--
			h = 1
		} else if h == 1 { // right
			if c != 0 {
				c--
				h *= -1
			}
		} else { // left
			if (v > 0 && r == 0) || (v < 0 && r == qr.len()-1) {
				v *= -1
				c--
				h *= -1
			} else {
				c++
				h *= -1
				r -= v
			}
		}
	}
}

func (qr *Qrcode) maskData() {

	qr.penalty = make([]int, 8)
	c := make(chan int, 8) // channel

	for i := 0; i < 8; i++ {
		// Penalty
		go qr.calcPenalty(i, c)
	}

	for k := 0; k < 8; k++ {
		<-c
	}

	mask := 0
	min := qr.penalty[0]

	for i := 1; i < 8; i++ {
		if qr.penalty[i] < min {
			min = qr.penalty[i]
			mask = i
		}
	}

	qr.setTypeBits(mask)
	qr.module = maskData(mask, qr.module)
}

func (qr *Qrcode) errorCorrectionCode(data []byte, block, blockIndex int, c chan int) {

	eccw := qr.ecCodeWords()                             // ec code word per block
	totalBlkCount := qr.blkCount()[0] + qr.blkCount()[1] // total blocks
	dcw := qr.totalDataCodeWords()                       // total data code words

	// message polynominal
	msg := polynominal{}
	for i := 0; i < len(data); i++ {
		x := len(data) + eccw - 1 - i
		msg[x] = int2exp[data[i]]
	}

	// generator polynominal
	gen := polynominal{0: 0}
	for i := 0; i < eccw; i++ {
		p := polynominal{1: 0, 0: i}
		gen = gen.multiply(p)
	}

	// error collection code words
	for i := len(data) - 1; i >= 0; i-- {
		p := polynominal{i: msg[i+eccw]}
		g := gen.multiply(p)
		msg = msg.xor(g)
		delete(msg, i+eccw)
	}

	ec := make([]byte, eccw)
	for i := 0; i < eccw; i++ {
		ec[i] = byte(exp2int[msg[eccw-1-i]])
	}

	for i, v := range data {
		j := blockIndex + i*totalBlkCount
		if j >= dcw {
			j -= qr.blkCount()[0]
		}
		qr.encodedData[j] = v
	}

	for i, v := range ec {
		qr.encodedData[dcw+blockIndex+i*totalBlkCount] = v
	}

	c <- 1
}

func (qr *Qrcode) encodeData() {

	// total data code words
	dcw := qr.totalDataCodeWords()

	// total code words
	cw := qr.totalCodeWords()

	// total blocks
	totalBlkCount := qr.blkCount()[0] + qr.blkCount()[1]

	// encode data
	var bytebuf *bytes.Buffer
	switch qr.mode {
	case "numeric":
		bytebuf = bytes.NewBuffer(qr.encodeNumeric())
	case "alphanum":
		bytebuf = bytes.NewBuffer(qr.encodeAlphanum())
	default: // default
		bytebuf = bytes.NewBuffer(qr.encodeByte())
	}

	// pad codeword
	for bytebuf.Len() < dcw {
		bytebuf.WriteByte(0xEC)
		if bytebuf.Len() < dcw {
			bytebuf.WriteByte(0x11)
		}
	}

	// 1. split encoded data into blocks
	// 2. generate error correction code for each block
	// 3. reallocate data
	qr.encodedData = make([]byte, cw)

	// goroutine
	c := make(chan int, totalBlkCount)

	for i, k := 0, 0; i < 2; i++ { // rs block 1, 2
		for j := 0; j < qr.blkCount()[i]; j, k = j+1, k+1 {
			go qr.errorCorrectionCode(bytebuf.Next(qr.dataCodeWords()[i]), i, k, c)
		}
	}

	// wait
	for k := 0; k < totalBlkCount; k++ {
		<-c
	}
}

// encode data ( numeric mode )
func (qr *Qrcode) encodeNumeric() []byte {

	// working bit array
	bitbuf := new(bitBuffer)

	// mode indicator
	bitbuf.append(1, 4)

	// character count indicator
	switch {
	case qr.Version <= 9:
		bitbuf.append(len(qr.data), 10)
	case qr.Version <= 26:
		bitbuf.append(len(qr.data), 12)
	default:
		bitbuf.append(len(qr.data), 14)
	}

	// string to bitstream
	bytebuf := bytes.NewBufferString(qr.data)
	for b := bytebuf.Next(3); len(b) > 0; b = bytebuf.Next(3) {
		switch len(b) {
		case 2:
			n, _ := strconv.Atoi(string(b))
			bitbuf.append(n, 7)
		case 1:
			n, _ := strconv.Atoi(string(b))
			bitbuf.append(n, 4)
		default: // 3
			n, _ := strconv.Atoi(string(b))
			bitbuf.append(n, 10)
		}
	}

	// terminator
	for i := 0; i < 4 && bitbuf.len() < qr.totalDataCodeBits(); i++ {
		bitbuf.append(0, 1)
	}

	return bitbuf.bytes()

}

// encode data ( alphanumeric mode )
func (qr *Qrcode) encodeAlphanum() []byte {

	// working bit array
	bitbuf := new(bitBuffer)

	// mode indicator
	bitbuf.append(1<<1, 4)

	// character count indicator
	switch {
	case qr.Version <= 9:
		bitbuf.append(len(qr.data), 9)
	case qr.Version <= 26:
		bitbuf.append(len(qr.data), 11)
	default:
		bitbuf.append(len(qr.data), 13)
	}

	// string to bitstream
	bytebuf := bytes.NewBufferString(qr.data)
	for b := bytebuf.Next(2); len(b) > 0; b = bytebuf.Next(2) {
		switch len(b) {
		case 1:
			bitbuf.append(alphanumTable[b[0]], 6)
		default: // 2
			bitbuf.append(alphanumTable[b[0]]*45+alphanumTable[b[1]], 11)
		}
	}

	// terminator
	for i := 0; i < 4 && bitbuf.len() < qr.totalDataCodeBits(); i++ {
		bitbuf.append(0, 1)
	}

	return bitbuf.bytes()
}

// encode data ( 8bitbyte mode )
func (qr *Qrcode) encodeByte() []byte {

	// working bit array
	bitbuf := new(bitBuffer)

	// mode indicator
	bitbuf.append(1<<2, 4)

	// character count indicator
	switch {
	case qr.Version <= 9:
		bitbuf.append(len(qr.data), 8)
	case qr.Version <= 26:
		bitbuf.append(len(qr.data), 16)
	default:
		bitbuf.append(len(qr.data), 16)
	}

	// string to bitstream
	bytebuf := bytes.NewBufferString(qr.data)
	for c, e := bytebuf.ReadByte(); e == nil; c, e = bytebuf.ReadByte() {
		bitbuf.appendByte(c)
	}

	// terminator
	bitbuf.append(0, 4)

	return bitbuf.bytes()
}

func (qr *Qrcode) calcPenalty(mask int, c chan int) {

	penalty := 0

	module := maskData(mask, setTypeBits(qr.module, qr.Level, mask))

	isColored := func(r, c int) bool { return module[r][c] > 0 }

	// Rule #1: five or more same colored pixels
	for r := 0; r < qr.len(); r++ { // horizontal scan
		for i := 0; i < qr.len()-1; i++ {
			for j := 1; ; j++ {
				if !(j < qr.len()-i &&
					isColored(r, i) == isColored(r, i+j)) {
					i += j - 1
					break
				}
				if j == 4 {
					penalty += 3
				}
				if j > 4 {
					penalty++
				}
			}
		}
	}

	for c := 0; c < qr.len(); c++ { // vertical scan
		for i := 0; i < qr.len()-1; i++ {
			for j := 1; ; j++ {
				if !(j < qr.len()-i &&
					isColored(i, c) == isColored(i+j, c)) {
					i += j - 1
					break
				}
				if j == 4 {
					penalty += 3
				}
				if j > 4 {
					penalty++
				}
			}
		}
	}

	// Rule #2: 2x2 block of same-colored pixels
	for r := 0; r < qr.len()-1; r++ {
		for c := 0; c < qr.len()-1; c++ {
			if isColored(r, c) == isColored(r+1, c) &&
				isColored(r, c) == isColored(r, c+1) &&
				isColored(r, c) == isColored(r+1, c+1) {
				penalty += 3
			}
		}
	}

	// Rule #3: o-x-o-o-o-x-o that have four light pixels on either or both sides
	for r := 0; r < qr.len(); r++ {
		for c := 0; c < qr.len()-6; c++ {
			if isColored(r, c) && !isColored(r, c+1) && isColored(r, c+2) &&
				isColored(r, c+3) && isColored(r, c+4) && !isColored(r, c+5) &&
				isColored(r, c+6) {
				if c < qr.len()-10 &&
					!isColored(r, c+7) && !isColored(r, c+8) &&
					!isColored(r, c+9) && !isColored(r, c+10) {
					penalty += 40
				} else if c >= 4 &&
					!isColored(r, c-1) && !isColored(r, c-2) &&
					!isColored(r, c-3) && !isColored(r, c-4) {
					penalty += 40
				}
				c += 4
			}
		}
	}

	for c := 0; c < qr.len(); c++ {
		for r := 0; r < qr.len()-6; r++ {
			if isColored(r, c) && !isColored(r+1, c) && isColored(r+2, c) &&
				isColored(r+3, c) && isColored(r+4, c) && !isColored(r+5, c) &&
				isColored(r+6, c) {
				if r < qr.len()-10 &&
					!isColored(r+7, c) && !isColored(r+8, c) &&
					!isColored(r+9, c) && !isColored(r+10, c) {
					penalty += 40
				} else if r >= 4 &&
					!isColored(r-1, c) && !isColored(r-2, c) &&
					!isColored(r-3, c) && !isColored(r-4, c) {
					penalty += 40
				}
				r += 4
			}
		}
	}

	// Rule #4: Color Ratio
	total := 0.0
	black := 0.0

	for r, row := range module {
		for c, _ := range row {
			if isColored(r, c) {
				black++
			}
			total++
		}
	}
	penalty += int(float64(int(math.Abs(black/total*100-50))) / 5 * 10)

	qr.penalty[mask] = penalty

	c <- 1
}

func maskData(mask int, data [][]int) [][]int {

	// mask 0-7
	maskFuncs := map[int]func(int, int) bool{
		0: func(row, col int) bool { return (row+col)%2 == 0 },
		1: func(row, col int) bool { return row%2 == 0 },
		2: func(row, col int) bool { return col%3 == 0 },
		3: func(row, col int) bool { return (row+col)%3 == 0 },
		4: func(row, col int) bool { return (row/2+col/3)%2 == 0 },
		5: func(row, col int) bool { return row*col%2+row*col%3 == 0 },
		6: func(row, col int) bool { return (row*col%2+row*col%3)%2 == 0 },
		7: func(row, col int) bool { return (row*col%3+(row+col)%2)%2 == 0 },
	}

	// copy
	module := make([][]int, 0)
	for _, row := range data {
		c := make([]int, len(row))
		copy(c, row)
		module = append(module, c)
	}

	for r, row := range module {
		for c, _ := range row {
			if !(module[r][c] == 1 || module[r][c] == -1) {
				continue
			}
			if maskFuncs[mask](r, c) {
				module[r][c] *= -1
			}
		}
	}
	return module
}

func setTypeBits(data [][]int, level ECLevel, mask int) [][]int {

	module := make([][]int, 0)
	for _, row := range data {
		c := make([]int, len(row))
		copy(c, row)
		module = append(module, c)
	}

	bits := new(bitBuffer)
	bits.append(typeInformationTable[level][mask], 15)

	len := len(module)

	for i, b := range *bits {
		bit := 2
		if b {
			bit = 2
		} else {
			bit = -2
		}

		switch {
		case i < 6:
			module[8][i] = bit
			module[len-1-i][8] = bit
		case i == 6:
			module[8][i+1] = bit
			module[len-1-i][8] = bit
		case i == 7:
			module[8][len-15+i] = bit
			module[8][8] = bit
		case i == 8:
			module[8][len-15+i] = bit
			module[7][8] = bit
		case i == 9:
			module[8][len-15+i] = bit
			module[5][8] = bit
		case i > 9:
			module[8][len-15+i] = bit
			module[(7*2)-i][8] = bit
		}
	}
	return module
}

func maxDataSize(version int, level ECLevel, mode string) (max int) {

	table := errorCorrectionTable[version][level]
	dcw := table[1]*table[2] + table[3]*table[4]

	switch mode {
	case "numeric":
		bitsCount := dcw*8 - 4 // mode bits(4 bits)

		switch { // data length bits
		case version <= 9:
			bitsCount -= 10
		case version <= 26:
			bitsCount -= 12
		default:
			bitsCount -= 14
		}

		max = bitsCount / 10 * 3 // 10bits for 3 chars
		if bitsCount%10 >= 7 {   // 7bits for 2 chars
			max += 2
		} else if bitsCount%10 >= 4 { // 4bits for 1 chars
			max++
		}
	case "alphanum":
		bitsCount := dcw*8 - 4 // mode bits(4 bits)

		switch { // data length bits
		case version <= 9:
			bitsCount -= 9
		case version <= 26:
			bitsCount -= 11
		default:
			bitsCount -= 13
		}

		max = bitsCount / 11 * 2 // 11bits for 2 characters
		if bitsCount%11 >= 6 {
			max++
		}
	case "8bitbyte":
		max = dcw - 1 // mode(4bit) + terminal(4bit) = 1byte

		switch { // data length bits
		case version <= 9:
			max -= 1
		default:
			max -= 2
		}
	}
	return
}

// bitBuffer
type bitBuffer []bool

func (p *bitBuffer) len() int {
	return len(*p)
}

func (p *bitBuffer) append(c, l int) {
	for i := 0; i < l; i++ {
		if c&(1<<uint(l-i-1)) > 0 {
			*p = append(*p, true)
		} else {
			*p = append(*p, false)
		}
	}
}

func (p *bitBuffer) get(i int) bool {
	return (*p)[i]
}

func (p *bitBuffer) appendByte(b byte) {
	p.append(int(b), 8)
}

func (p *bitBuffer) bytes() []byte {

	// copy
	s := make(bitBuffer, p.len())
	copy(s, *p)

	for s.len()%8 != 0 {
		s.append(0, 1)
	}

	l := len(s) / 8
	bytebuf := bytes.NewBuffer([]byte{})

	for i := 0; i < l; i++ {
		bits := s[(i * 8):(i*8 + 8)]
		c := byte(0)
		for j := 0; j < 8; j++ {
			if bits[j] {
				c += 1 << uint(8-j-1)
			}
		}
		bytebuf.WriteByte(c)
	}
	return bytebuf.Bytes()
}

// polynominal ( Galois field 256 )
// used to calculate error correction code
// (a^x)*(x^25) + (a^y)*(x^24) + (a^z)*(x^23) + .... => {25:x,24:y,23:z,.....}
type polynominal map[int]int

func (p0 polynominal) multiply(p1 polynominal) polynominal {
	poly := polynominal{}
	for x1, a1 := range p1 {
		for x0, a0 := range p0 {
			x := x1 + x0
			for x >= 255 {
				x -= 255
				x = (x >> 8) + (x & 255)
			}
			a := a1 + a0
			for a >= 255 {
				a -= 255
				a = (a >> 8) + (a & 255)
			}
			if poly[x] == 0 {
				poly[x] = exp2int[a]
			} else {
				anti := poly[x] ^ exp2int[a]
				poly[x] = anti
			}
		}
	}
	for x, a := range poly {
		poly[x] = int2exp[a]
	}
	return poly
}

func (p0 polynominal) xor(p1 polynominal) polynominal {
	poly := polynominal{}
	for x0, a0 := range p0 {
		poly[x0] = exp2int[a0]
	}
	for x1, a1 := range p1 {
		if poly[x1] == 0 {
			poly[x1] = exp2int[a1]
		} else {
			poly[x1] = poly[x1] ^ exp2int[a1]
		}
	}
	for x, a := range poly {
		poly[x] = int2exp[a]
	}
	return poly
}
