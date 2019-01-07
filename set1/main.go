package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

func main() {
	two()
}
