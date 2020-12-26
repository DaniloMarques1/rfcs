package base64

import "testing"

func Test_getBitsFromString(t *testing.T) {
	s := "Ola"
	// Ola in binary
	expected := []uint8{0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1}
	bits := getBitsFromString(s)
	if len(bits) != len(expected) {
		t.Fatal("Invalid bits returned")
	}
	for i, _ := range bits {
		if bits[i] != expected[i] {
			t.Fatal("Invalid bits returned")
		}
	}
	// Danilo marques in binary
	expected = []uint8{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1}
	s = "Danilo marques"
	bits = getBitsFromString(s)
	if len(bits) != len(expected) {
		t.Fatal("Invalid bits returned. Lengths differ")
	}
	for i, _ := range bits {
		if bits[i] != expected[i] {
			t.Fatalf("Invalid bits returned. %v, %v", bits[i], expected[i])
		}
	}
}

func Test_GetIntFromBits(t *testing.T) {
	bits := []uint8{0, 1, 0, 0, 0, 1}
	idx1 := getIntFromBits(bits)
	bits = []uint8{0, 1, 0, 1, 1, 0}
	idx2 := getIntFromBits(bits)
	bits = []uint8{1, 0, 0, 1, 0, 0}
	idx3 := getIntFromBits(bits)
	bits = []uint8{1, 1, 1, 1, 1, 1, 1, 1}
	idx4 := getIntFromBits(bits)
	bits = []uint8{1, 0, 0, 0, 0, 0, 0, 0}
	idx5 := getIntFromBits(bits)
	if idx1 != 17 || idx2 != 22 || idx3 != 36 || idx4 != 255 || idx5 != 128 {
		t.Fatal("Wrong integer returned from the six bits group")
	}
}

func Test_GetFromTheAlphabet(t *testing.T) {
	ret1 := getFromTheAlphabet(0)  // A
	ret2 := getFromTheAlphabet(17) // R
	ret3 := getFromTheAlphabet(33) // h
	ret4 := getFromTheAlphabet(63) // /
	if ret1 != "A" || ret2 != "R" || ret3 != "h" || ret4 != "/" {
		t.Fatal("Returned the wrong encode")
	}
}

func Test_EncodeString(t *testing.T) {
	stringToBeEncoded := "Danilo Marques"
	expected := "RGFuaWxvIE1hcnF1ZXM="
	result := EncodeString(stringToBeEncoded)
	if expected != result {
		t.Fatal("Encode returned a wrong encoded string")
	}
	stringToBeEncoded = "Ticking away the moments that make up a dull day Fritter and waste the hours in an offhand way. Kicking around on a piece of ground in your home town Waiting for someone or something to show you the way."
	expected = "VGlja2luZyBhd2F5IHRoZSBtb21lbnRzIHRoYXQgbWFrZSB1cCBhIGR1bGwgZGF5IEZyaXR0ZXIgYW5kIHdhc3RlIHRoZSBob3VycyBpbiBhbiBvZmZoYW5kIHdheS4gS2lja2luZyBhcm91bmQgb24gYSBwaWVjZSBvZiBncm91bmQgaW4geW91ciBob21lIHRvd24gV2FpdGluZyBmb3Igc29tZW9uZSBvciBzb21ldGhpbmcgdG8gc2hvdyB5b3UgdGhlIHdheS4="
	result = EncodeString(stringToBeEncoded)
	if expected != result {
		t.Fatalf("Encode returned a wrong encoded string\n%v\n", result)
	}

	stringToBeEncoded = `Tired of lying in the sunshine staying home to watch the rain.
You are young and life is long and there is time to kill today.
And then one day you find ten years have got behind you.
No one told you when to run, you missed the starting gun.`
	expected = `VGlyZWQgb2YgbHlpbmcgaW4gdGhlIHN1bnNoaW5lIHN0YXlpbmcgaG9tZSB0byB3YXRjaCB0aGUgcmFpbi4KWW91IGFyZSB5b3VuZyBhbmQgbGlmZSBpcyBsb25nIGFuZCB0aGVyZSBpcyB0aW1lIHRvIGtpbGwgdG9kYXkuCkFuZCB0aGVuIG9uZSBkYXkgeW91IGZpbmQgdGVuIHllYXJzIGhhdmUgZ290IGJlaGluZCB5b3UuCk5vIG9uZSB0b2xkIHlvdSB3aGVuIHRvIHJ1biwgeW91IG1pc3NlZCB0aGUgc3RhcnRpbmcgZ3VuLg==`
	result = EncodeString(stringToBeEncoded)
	if expected != result {
		t.Fatalf("Encode returned a wrong encoded string\n%v\n", result)
	}
}

func Test_removePadding(t *testing.T) {
	str := "RWk="
	expected := "RWk"
	result := removePadding(str)
	if result != expected {
		t.Fatal("Padding not removed")
	}
}

func Test_getIdxFromAlphabet(t *testing.T) {
	c := "C"
	expected := 2
	idx := getIdxFromAlphabet(c)
	if idx != expected {
		t.Fatal("Returning wrong index")
	}
}

func Test_DecodeString(t *testing.T) {
	str := "this is my string that i will encode"
	encodedString := EncodeString(str)
	decodedString := DecodeString(encodedString)
	if decodedString != str {
		t.Fatal("Decoded string differs")
	}
	str = "In programming, Base64 is a group of binary-to-text encoding schemes that represent binary data (more specifically a sequence of 8-bit bytes) in an ASCII string format by translating it into a radix-64 representation."
	encodedString = EncodeString(str)
	decodedString = DecodeString(encodedString)
	if decodedString != str {
		t.Fatalf("Decoded string differs %v", decodedString)
	}
}
