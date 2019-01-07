package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func one() {
	// inputString := "\x49\x27\x6d\x20\x6b\x69\x6c\x6c\x69\x6e\x67\x20\x79\x6f\x75\x72\x20\x62\x72\x61\x69\x6e\x20\x6c\x69\x6b\x65\x20\x61\x20\x70\x6f\x69\x73\x6f\x6e\x6f\x75\x73\x20\x6d\x75\x73\x68\x72\x6f\x6f\x6d"
	// expectedString := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// Check the examples section https://en.wikipedia.org/wiki/Base64
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedString := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	inputBytes, err := hex.DecodeString(hexString)
	check(err)

	// expectedBytes, err := base64.StdEncoding.DecodeString(expectedString)
	// if err != nil {
	// 	fmt.Println("Error decoding expected")
	// 	panic(err)
	// }

	fmt.Println("Input String: " + string(inputBytes))

	// encode the raw bytes to base64 string
	encoded := base64.StdEncoding.EncodeToString(inputBytes)

	fmt.Println("base64 encoded: " + encoded)
	if strings.EqualFold(encoded, expectedString) {
		fmt.Println("Success!")
	}

	output, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(output))
}

func two() {
	inputString := "1c0111001f010100061a024b53535009181c"
	xorString := "686974207468652062756c6c277320657965"
	expectedString := "746865206b696420646f6e277420706c6179"

	// hex decode strings
	inputBytes, err := hex.DecodeString(inputString)
	fmt.Println(inputBytes)
	check(err)

	xorBytes, err := hex.DecodeString(xorString)
	fmt.Println(xorBytes)
	check(err)

	if len(xorBytes) != len(inputBytes) {
		fmt.Println("Byte buffers different lengths")
		panic("ruh roh")
	}

	outputBytes := make([]byte, len(inputBytes))
	for i, iByte := range inputBytes {
		val := iByte ^ xorBytes[i]
		outputBytes[i] = val
	}

	fmt.Println(outputBytes)

	outputString := hex.EncodeToString(outputBytes)
	fmt.Println(outputString)

	if strings.EqualFold(expectedString, outputString) {
		fmt.Println("Success")
	}
}

// https://en.wikipedia.org/wiki/Letter_frequency
func score(s string) int {
	count := 0
	// freq := make(map[rune]int)
	// for _, c := range s {
	// 	if curr, ok := freq[c]; ok {
	// 		freq[c] = curr + 1
	// 	} else {
	// 		freq[c] = 1
	// 	}
	// }
	// for k, v := range freq {
	// 	fmt.Printf("%s:%d ", string(k), v)
	// }
	// fmt.Println("")

	// 1st idea: give it a point if it letter is in top half of alphabet
	// 		Works '$ b:58 s:18 -- Cooking MC's like a pound of bacon'
	// 		Can it be better?
	// 2: not 1 pt, but get extra points if more frequently used (1 -13)
	// 		Still worked
	// 3: added toLower to try and get more, but that royally messed things up
	// 4: try using actual percentages*100 rather than 1-13
	//		this at least got cOOKINGmcSLIKEAPOUNDOFBACON to be same score as answer
	// 5: try adding points for a space
	//  		back on tap woot woot
	// 6: this could be improved a lot I'm sure
	letters := map[string]int{
		" ": 13,
		"e": 13,
		"t": 9,
		"a": 8,
		"o": 8,
		"i": 7,
		"n": 7,
		"s": 6,
		"h": 6,
		"r": 6,
		"d": 4,
		"l": 4,
		"c": 3,
		"u": 3,
		"m": 2,
		"w": 2,
		"f": 2,
		"g": 2,
		"y": 2,
		"p": 2,
		"b": 1,
		"v": 1,
		"k": 1,
		"j": 0,
		"x": 0,
		"q": 0,
		"z": 0,
	}

	for _, c := range s {
		char := strings.ToLower(string(c))
		fmt.Print(char)
		if pts, ok := letters[char]; ok {
			count += pts
		}
	}
	fmt.Println("")
	return count
}

func three() {
	// a ^ b = encoded, a ^ b ^ b = a, find the key
	inputString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// does it mean a single character for each byte? or for the whole thing + 0s?
	// Probably for each byte, which makes the most sense

	// has to be ascii right? 0-255 to fit the byte or so I'd hope it's not unicode
	inputBytes, err := hex.DecodeString(inputString)
	fmt.Println(inputBytes)
	check(err)

	// after first run, looks like should be 0x58, need to come up with scoring method

	// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
	var xb byte
	type item struct {
		String string
		Cipher byte
		Score  int
	}
	outputBytes := make([]byte, len(inputBytes))
	var items []item
	for char := 0; char < 255; char++ {
		xb = byte(char)
		fmt.Printf("---> char is %x\n", xb)
		for i, b := range inputBytes {
			outputBytes[i] = b ^ xb
		}
		str := string(outputBytes)
		items = append(items, item{str, xb, score(str)})
		fmt.Println("-------------------------------")

	}

	// sort our scores map

	sort.Slice(items, func(i, j int) bool {
		return items[i].Score > items[j].Score
	})

	count := 0
	for _, i := range items {
		if count == 10 {
			break
		}
		fmt.Printf("$ b:%x s:%d -- %s\n", i.Cipher, i.Score, i.String)
		count++
	}
}

func main() {
	three()
}
