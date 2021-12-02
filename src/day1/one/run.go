package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("../depths.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counter int
	var prev int

	for scanner.Scan() {
		var input, _ = strconv.Atoi(scanner.Text())

		if prev > 0 && input > prev {
			counter++
		}

		prev = input
	}

	print(counter)
}
