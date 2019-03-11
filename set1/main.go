package set1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"

	"example.com/utils"
)

func Challenge1() {
	utils.Header("Convert hex to base64")
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedStr := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// decode the hex string to bytes
	bytes, err := hex.DecodeString(hexStr)
	utils.Check(err)

	// encode raw bytes to base64
	encoded := base64.StdEncoding.EncodeToString(bytes)

	utils.CheckExpected(expectedStr, encoded)
}

func Challenge2() {
	utils.Header("Fixed XOR")
	msg := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	// hex decode
	msgBytes, err := hex.DecodeString(msg)
	utils.Check(err)

	keyBytes, err := hex.DecodeString(key)
	utils.Check(err)

	if len(msgBytes) != len(keyBytes) {
		fmt.Println("Byte buffers different lengths")
		panic("ruh roh")
	}

	// xor against key, isn't this One Time Pad?
	outputBytes := make([]byte, len(msgBytes))
	for i, msgByte := range msgBytes {
		val := msgByte ^ keyBytes[i]
		outputBytes[i] = val
	}

	output := hex.EncodeToString(outputBytes)

	utils.CheckExpected(expected, output)
}

func Challenge3() {
	utils.Header("Single-byte XOR cipher")
	ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	cipherBytes, err := hex.DecodeString(ciphertext)
	utils.Check(err)

	// for each letter in the alphabet, have to decrypt byte
	// A - Z 65 - 90
	// a - z 97 - 122

	type item struct {
		score   int
		key     byte
		decoded []byte
	}

	xor := func(k byte) item {
		decoded := make([]byte, len(cipherBytes))
		for i, b := range cipherBytes {
			decoded[i] = b ^ k
		}

		score := utils.EnglishScore(string(decoded))
		return item{
			score:   score,
			key:     k,
			decoded: decoded,
		}
	}

	var items []item

	// add uppercase & lowercase to list
	for i := 0; i < 26; i++ {
		upper := byte(65 + i)
		lower := byte(97 + i)

		items = append(items, xor(upper))
		items = append(items, xor(lower))
	}

	// have scored list; sort and display
	sort.Slice(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	count := 0
	for _, i := range items {
		if count == 10 {
			break
		}
		fmt.Printf("$ %s:%x s:%d -- %s\n", string(i.key), i.key, i.score, i.decoded)
		count++
	}
}

func Challenge4() {
	utils.Header("Detect single-character XOR")
	// file is list of fixed-length strings of 60 chars
	// one of them was encrypted by a single char xor, gotta find it
	fileName := "set1Challenge4.txt"
	fileURL := "https://cryptopals.com/static/challenge-data/4.txt"

	file := utils.GetCachedFile(fileName, fileURL)
	defer file.Close()

	// this could be a struct again, but ehh
	bestScore := -1
	bestOutput := ""
	bestKey := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key, score, output := utils.GetBestScore(line)
		fmt.Printf("%d : %s\n", score, output)
		if score > bestScore {
			bestScore = score
			bestOutput = output
			bestKey = string(key)
		}
	}

	// 1st run: total garbage coming out
	// 		Possible help:
	//          https://stackoverflow.com/questions/52969757/stuck-on-cryptopals-challenge-4-in-go
	//			https://crypto.stackexchange.com/questions/30209/developing-algorithm-for-detecting-plain-text-via-frequency-analysis
	fmt.Println("**** Top Answer ****")
	fmt.Printf("%s:%d - %s\n", bestKey, bestScore, bestOutput)
}

func Challenge5() {
	fmt.Println("5")
}

func Challenge6() {
	fmt.Println("6")
}

func Challenge7() {
	fmt.Println("7")
}

func Challenge8() {
	fmt.Println("8")
}
