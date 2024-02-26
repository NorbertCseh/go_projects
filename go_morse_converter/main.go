package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	signs := readFile("morseCodes.txt")
	text := getSentence()
	fmt.Println(translate(text, signs))
}

func getSentence() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please add a sentence to translate")
		text, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		return text
	}
}

func readFile(fileName string) map[string]string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")
	codes := make(map[string]string)

	for _, v := range lines {
		if v != "" {
			temp := strings.Split(v, " ")
			codes[temp[0]] = temp[1]
		}
	}
	return codes
}

func translate(sentece string, signs map[string]string) string {
	chars := strings.Split(sentece, "")
	translated_stuff := ""

	for _, v := range chars {
		translated_stuff = translated_stuff + signs[strings.ToUpper(v)]
	}

	return translated_stuff
}
