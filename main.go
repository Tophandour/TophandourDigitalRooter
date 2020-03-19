// TophandourDigitalRooter project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	shouldContinue := true
	for shouldContinue {
		fmt.Println("Please input some text")
		bytes, err := buffer.ReadBytes('\n')
		bytes = bytes[:len(bytes)-1]
		text := strings.Trim(string(bytes), "\n")
		if err != nil {
			fmt.Println(text)
		} else {
			lowercase := strings.ToLower(string(bytes))
			bytes = []byte(lowercase)
			fmt.Println("Digital root is: " + CalculateDigitalRoot(bytes))
		}
		fmt.Println("Continue? Input n to quit, anything else to continue")
		bytes, err = buffer.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		text = string(bytes)
		text = text[:len(text)-2]
		shouldContinue = text != "n"
	}
	os.Exit(0)
}

func CalculateDigitalRoot(text []byte) string {
	charSum := SumCharacters(text)
	digiRoot := (charSum-1)%9 + 1
	if digiRoot < 0 {
		digiRoot = 0
	}
	return strconv.Itoa(digiRoot)
}

func SumCharacters(text []byte) int {
	checkArray := " abcdefghijklmnopqrstuvwxyz"
	sum := 0
	for i := 0; i < len(text); i++ {
		charValue := strings.IndexByte(checkArray, text[i])
		if charValue != -1 {
			sum += charValue
		}
	}
	return sum
}
