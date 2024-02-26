package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	signs := readFile("morseCodes.txt")
	fmt.Println(translate("    ", signs))
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
