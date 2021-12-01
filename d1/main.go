package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var prev, result int = -1, 0

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())

		if prev != -1 {
			if number > prev {
				result = result + 1
			}
		}
		prev = number
	}
	fmt.Println(result)
}

func p2() {

}

func main() {
	p1()
	p2()
}
