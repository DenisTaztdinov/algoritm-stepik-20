package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var a = []rune(text)
	text = strings.Trim(text, "\n")
	text = strings.Trim(text, "\r")
	if unicode.IsUpper(a[0]) == true && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
