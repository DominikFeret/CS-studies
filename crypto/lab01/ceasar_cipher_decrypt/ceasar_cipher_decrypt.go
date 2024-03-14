package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dict map[rune]int
var dictRev map[int]rune

const dictLen = 52

func main() {
	initDict()
	str, k := getInput()

	enc := ceasarCipher(str, k, dictLen)

	fmt.Printf("%s", enc)
}

func ceasarCipher(str []rune, k int, z int) string {
	enc := ""
	for _, c := range str {
		if !isLiteral(c) {
			enc += string(c)
			continue
		}

		encVal := mod(dict[c]-k, z)
		enc += string(dictRev[encVal])
	}

	return enc
}

func isLiteral(c rune) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
}

func mod(num int, z int) int {
	num = num % z
	if num < 0 {
		num += z
	}
	return num
}

func initDict() {
	dict = make(map[rune]int)
	dictRev = make(map[int]rune)

	offset := 97
	for i := 0; i < 26; i++ {
		dict[rune(offset+i)] = i
		dictRev[i] = rune(offset + i)
	}
	offset = 65
	for i := 0; i < 26; i++ {
		dict[rune(offset+i)] = i + 26
		dictRev[i+26] = rune(offset + i)
	}
}

func getInput() ([]rune, int) {
	input := readInput()

	var s []rune
	for i, c := range input {
		if i == len(input)-1 {
			break
		}
		s = append(s, []rune(c)...)
		if i != len(input)-2 {
			s = append(s, '\n')
		}
	}
	k, err := strconv.Atoi(input[len(input)-1])
	if err != nil {
		panic(err)
	}

	return s, k
}

func readInput() []string {
	r := bufio.NewScanner(os.Stdin)
	s := []string{}

	for r.Scan() {
		s = append(s, r.Text())
	}

	return s
}
