package base64

import (
	"math"
)

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "+", "/", "="}

// return a uint8 slice with all the bits from
// the string
func getBitsFromString(chars string) []uint8 {
	var bits []uint8
	for _, c := range chars {
		bits = append(bits, getBitsSlice(uint8(c), 8)...)
	}

	return bits
}

// returns the (size) bits of a char(int)
func getBitsSlice(char uint8, size int) []uint8 {
	var bits []uint8
	for char > 0 {
		rest := char % 2
		bits = append(bits, rest)
		res := char / 2
		char = res
	}
	bits = fillingBitsSlice(size, bits)

	var retBits []uint8
	for i := len(bits) - 1; i >= 0; i-- {
		retBits = append(retBits, bits[i])
	}

	return retBits
}

func fillingBitsSlice(n int, slice []uint8) []uint8 {
	for len(slice) < n {
		slice = append(slice, 0)
	}

	return slice
}

// return the correct char that must be
// used to encode using the alphabet
func getFromTheAlphabet(idx int) string {
	return alphabet[idx]
}

// returns the decimal rperesentation of the 6
// bits received that will be used as index in the alphabet
func getIntFromBits(bits []uint8) int {
	start := 0
	end := float64(len(bits) - 1)
	var res float64
	for start < len(bits) {
		res += float64(bits[start]) * math.Pow(2, end)
		end--
		start++
	}

	return int(res)
}

func encodeBits(bits []uint8) string {
	rounds := len(bits) / 6 // how many times will be able to get 6 bits
	start := 0
	end := start + 6

	var encodeString string
	for rounds > 0 {
		sliceBits := bits[start:end]
		idx := getIntFromBits(sliceBits)
		encodeString += getFromTheAlphabet(idx)
		rounds--
		start = end
		end = start + 6
	}
	rest := len(bits) % 6 // shows the last block bits, if the last block does not have 6 bits.
	if rest > 0 {
		sliceBits := bits[start:]
		paddings := 0
		// adding padding
		for len(sliceBits) < 6 {
			sliceBits = append(sliceBits, 0)
			paddings++
		}
		idx := getIntFromBits(sliceBits)
		encodeString += getFromTheAlphabet(idx)
		for paddings > 0 {
			encodeString += getFromTheAlphabet(64)
			paddings -= 2
		}
	}

	return encodeString
}

// return the index of the str in the base64
// alphabet
func getIdxFromAlphabet(str string) int {
	for idx, v := range alphabet {
		if v == str {
			return idx
		}
	}

	return -1
}

// removes the = which represents padding in
// base64 encoded strings
func removePadding(str string) string {
	var s string
	for _, c := range str {
		if string(c) != "=" {
			s += string(c)
		}
	}

	return s
}

// receiveis a string and return the string
// encoded in base64
func EncodeString(str string) string {
	bits := getBitsFromString(str)
	encodedString := encodeBits(bits)
	return encodedString
}

// reads the bits received in blocks of 8
// decoding the int resulted from the bits
// using the ascii
func decodeBits(bits []uint8) string {
	start := 0
	end := start + 8
	rounds := len(bits) / 8
	var str string
	for rounds > 0 {
		sliceBits := bits[start:end]
		b := getIntFromBits(sliceBits)
		str += string(byte(b))
		rounds--
		start = end
		end = start + 8
	}

	return str
}

// receives a base64 encoded string and return the
// ascii representation of it
func DecodeString(encodedString string) string {
	encodedString = removePadding(encodedString)
	var bits []uint8
	for _, c := range encodedString {
		idx := uint8(getIdxFromAlphabet(string(c)))
		bits = append(bits, getBitsSlice(idx, 6)...)
	}
	decodedString := decodeBits(bits)
	return decodedString
}
