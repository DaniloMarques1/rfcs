package base64

import (
	"math"
)

const alphabetBase = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const alphabetUrl  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_="

var alphabet string

// return a uint8 slice with all the bits from
// the string
func getBitsFromString(str string) []uint8 {
	var bits []uint8
	for _, c := range str {
		bits = append(bits, getBitsSlice(uint8(c), 8)...)
	}

	return bits
}

// returns the bits of the value passed
// with (size) bits
func getBitsSlice(value uint8, size int) []uint8 {
	var bits = make([]uint8, size)
	idx := size - 1
	for value > 0 {
		rest := value % 2
		bits[idx] = rest
		value /= 2
		idx--
	}

	bits = fillingBitsSlice(bits, size)

	return bits
}

// adds zeros to fill the required slice size
// and return the slice with newer values added
func fillingBitsSlice(slice []uint8, size int) []uint8 {
	for len(slice) < size {
		slice = append(slice, 0)
	}

	return slice
}

// return the correct char that must be
// used to encode using the alphabet
func getFromTheAlphabet(idx int) byte {
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

// takes the string bits and encode the six bits block
// using the base64 alphabet, returned a encodedString
func encodeBits(bits []uint8) string {
	rest := len(bits) % 6 // if rest > 0 then it needs to be add a padding
	paddingsAdded := 0
	for rest > 0 {
		bits = append(bits, 0)
		rest = len(bits) % 6
		paddingsAdded++ // if we added two 0s, we need one =, 4 zeros, ==
	}

	rounds := len(bits) / 6 // how many 6 bits blocks does bits has
	start := 0
	end := start + 6
	var encodedString string
	for rounds > 0 {
		sliceBits := bits[start:end]
		idxInAlphabet := getIntFromBits(sliceBits)
		encodedString += string(getFromTheAlphabet(idxInAlphabet))
		rounds--
		start = end
		end = start + 6
	}

	// adding the 64 char, aka padding
	for paddingsAdded > 0 {
		encodedString += string(getFromTheAlphabet(64))
		paddingsAdded -= 2
	}

	return encodedString
}

// return the index of the str in the base64
// alphabet
func getIdxFromAlphabet(char rune) int {
	for idx, v := range alphabet {
		if v == char {
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
func encodeString(str string) string {
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
func decodeString(encodedString string) string {
	encodedString = removePadding(encodedString)
	var bits []uint8
	for _, c := range encodedString {
		idx := uint8(getIdxFromAlphabet(c))
		bits = append(bits, getBitsSlice(idx, 6)...)
	}
	decodedString := decodeBits(bits)
	return decodedString
}

// encode a string using the base64 alphabet
func StdEncodeString(str string) string {
	alphabet = alphabetBase
	return encodeString(str)
}

// encode a string using the base64url slphabet
// as defined in the rfc 4648
func UrlEncodeString(str string) string {
	alphabet = alphabetUrl
	return encodeString(str)
}

// decoding a base64 encoded string
func StdDecodeString(str string) string {
	alphabet = alphabetBase
	return decodeString(str)
}

// decoding a url safe encoded string
func UrlDecodeString(str string) string {
	alphabet = alphabetUrl
	return decodeString(str)
}
