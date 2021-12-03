package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func p2(file *os.File) {
	scanner := bufio.NewScanner(file)
	aim, hor, depth := 0, 0, 0

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		command := tokens[0]
		steps, _ := strconv.Atoi(tokens[1])

		if command == "down" {
			aim += steps
		} else if command == "up" {
			aim -= steps
		} else {
			hor += steps
			depth += aim * steps
		}
	}

	result := hor * depth
	fmt.Printf("result = %v\n", result)
}

func p1(file *os.File) {
	scanner := bufio.NewScanner(file)
	m := make(map[string]int)

	for scanner.Scan() {
		input := scanner.Text()
		tokens := strings.Split(input, " ")
		command := tokens[0]
		steps, _ := strconv.Atoi(tokens[1])

		temp, ok := m[command]
		if ok {
			m[command] = temp + steps
		} else {
			m[command] = steps
		}
	}

	result := (m["down"] - m["up"]) * m["forward"]
	fmt.Printf("result = %v\n", result)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1(file)
	file.Seek(0, io.SeekStart)
	p2(file)
}
