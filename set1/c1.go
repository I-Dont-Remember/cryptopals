package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// func hexToBase64(hexData []byte) []byte {
// 	base64 := map[byte]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H",
// 		8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R",
// 		18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z", 26: "a", 27: "b",
// 		28: "c", 46: "u"}
// 	// handle if not multiple of 3
// 	subsets := [][]byte{}
// 	for {
// 		if len(hexData) < 3 {
// 			break
// 		}
// 		triplet := hexData[0:3]
// 		hexData = hexData[3:]
// 		subsets = append(subsets, triplet)
// 	}

// 	// have list of the triplets we need for each base64 chars.
// 	// convert 4 6-bit chunks from those 3 8 bit chunks
// 	for i := 0; i < len(subsets); i++ {
// 		bytes := subsets[i]
// 		first := bytes[0] >> 2 // shift right to ignore top two bits
// 		second := (bytes[0] & 3 << 4) + (bytes[1] >> 4)
// 		third := (bytes[1] & 0x0f << 2) + (bytes[2] >> 6)
// 		fourth := bytes[2] & 0x3f
// 		fmt.Printf("%s %s %s %s\n", first, second, third, fourth)
// 		quad := append([]byte(base64[first]), []byte(base64[second]), []byte(base64[third]), []byte(base64[fourth]))
// 		fmt.Println(quad)
// 		fmt.Printnln()
// 	}

// 	return nil
// }

func hexToB64(hex []byte) []byte {
	fmt.Println(b64.StdEncoding.EncodeToString(hex))
	return []byte(b64.URLEncoding.EncodeToString(hex))
}

// do some kind of bitmask like 11111100000... so we get just those first bits,
// then move it along the line of the string?

func main() {
	// inputString := "\x49\x27\x6d\x20\x6b\x69\x6c\x6c\x69\x6e\x67\x20\x79\x6f\x75\x72\x20\x62\x72\x61\x69\x6e\x20\x6c\x69\x6b\x65\x20\x61\x20\x70\x6f\x69\x73\x6f\x6e\x6f\x75\x73\x20\x6d\x75\x73\x68\x72\x6f\x6f\x6d"
	// expectedString := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// Check the examples section https://en.wikipedia.org/wiki/Base64
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedString := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	inputBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding input")
		panic(err)
	}

	// expectedBytes, err := b64.StdEncoding.DecodeString(expectedString)
	// if err != nil {
	// 	fmt.Println("Error decoding expected")
	// 	panic(err)
	// }

	fmt.Println("Input String: " + string(inputBytes))

	// encode the raw bytes to base64 string
	encoded := b64.StdEncoding.EncodeToString(inputBytes)

	fmt.Println("B64 encoded: " + encoded)
	if strings.EqualFold(encoded, expectedString) {
		fmt.Println("Success!")
	}

	output, _ := b64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(output))
}

// https://brianritchie.me/crytopals-solution-set-1-challenge-1
// https://gobyexample.com/base64-encoding
// https://xdmtk.wordpress.com/2018/03/26/cryptopals-walk-through-series-set-1-challenge-1/