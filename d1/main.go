package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getsum(queue []int) int {
	sum := 0
	for _, v := range queue {
		sum += v
	}
	return sum
}

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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var queue []int
	firstTripletComplete := false
	sum := 0
	count := 0

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())

		if firstTripletComplete {
			queue = queue[1:]
			queue = append(queue, number)
			newsum := getsum(queue)

			if newsum > sum {
				count += 1
			}
			sum = newsum
		} else {
			queue = append(queue, number)

			if len(queue) == 3 {
				firstTripletComplete = true
				sum = getsum(queue)
			}
		}
	}
	fmt.Println(count)
}

func main() {
	p1()
	p2()
}
