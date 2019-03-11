package utils

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Header makes it look pretty
func Header(s string) {
	fmt.Println("========================")
	fmt.Println(s)
	fmt.Println("========================")

}

// Check checks if error needs panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func HexToBase64String(hexStr string) string {
	// decode the hex string to bytes
	bytes, err := hex.DecodeString(hexStr)
	Check(err)

	// encode raw bytes to base64
	return base64.StdEncoding.EncodeToString(bytes)
}

func CheckExpected(expected, output string) {
	fmt.Println("Expected: ", expected)
	fmt.Println("Output: ", output)
	if strings.EqualFold(expected, output) {
		fmt.Println("** Success! **")
	} else {
		fmt.Println("!!! You suck !!!")
	}
}

func EnglishScore(s string) int {
	// https://en.wikipedia.org/wiki/Letter_frequency
	count := 0
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
		if points, ok := letters[char]; ok {
			count += points
		}
	}
	return count
}

// GetCachedFile keeps a cached version of the file
func GetCachedFile(fileName, fileURL string) *os.File {
	var file *os.File

	_, err := os.Stat(fileName)
	// check local system first so we don't keep grabbing it
	if os.IsNotExist(err) {
		if fileURL == "" {
			panic(errors.New("not given a file URL to download a new copy"))
		}
		fmt.Println("Downloading from ", fileURL)
		resp, err := http.Get(fileURL)
		Check(err)

		defer resp.Body.Close()

		file, err = os.Create(fileName)
		Check(err)

		defer file.Close()

		io.Copy(file, resp.Body)

	} else {
		file, err = os.Open(fileName)
		Check(err)

		defer file.Close()
	}

	fmt.Println(file.Name())
	return file

}
