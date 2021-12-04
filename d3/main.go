package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)
	slice := strings.Split(content, "\n")
	n := len(slice)
	m := len(slice[0])
	c := make([]int, m)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		bytes := []rune(slice[i])
		for j := range matrix[0] {
			matrix[i][j], _ = strconv.Atoi(string(bytes[j]))

			if matrix[i][j] == 1 {
				c[j] += 1
			} else {
				c[j] -= 1
			}
		}
	}
	var b1, b2 string
	for _, elem := range c {
		if elem > 0 {
			b1 += "1"
			b2 += "0"
		} else {
			b1 += "0"
			b2 += "1"
		}
	}
	x, _ := strconv.ParseInt(b1, 2, 64)
	y, _ := strconv.ParseInt(b2, 2, 64)
	fmt.Println(x * y)
}
