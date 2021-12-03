package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
