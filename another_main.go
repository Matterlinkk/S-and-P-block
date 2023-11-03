package main

import (
	"fmt"
	"strings"
)

var pBox = [8]int{2, 4, 1, 6, 3, 0, 5, 7}

func pboxencrypt(input string) string {

	encryptValue := make([]string, len(input))

	for i := range input {
		encryptValue[i] = string(input[pBox[i]])
	}

	return strings.Join(encryptValue, "")
}

func pboxdecrypt(input string) string {

	encryptValue := make([]string, len(input))

	for i := range input {
		encryptValue[pBox[i]] = string(input[i])
	}

	return strings.Join(encryptValue, "")
}

func main() {
	msg := "qwertyui"

	fmt.Println(pboxencrypt(msg))

}
