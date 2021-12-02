package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("../depths.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var measurements = make([]int, 0, 2000)

	for scanner.Scan() {
		var input, _ = strconv.Atoi(scanner.Text())
		measurements = append(measurements, input)
	}

	fmt.Println(measurements)

	var counter int
	var prevSum int

	for i := 0; i < len(measurements)-2; i++ {
		var sum = measurements[i] + measurements[i+1] + measurements[i+2]

		if prevSum > 0 && sum > prevSum {
			counter++
		}

		fmt.Println(measurements[i], " + ", measurements[i+1], " + ", measurements[i+2], " = ", sum, " > ", prevSum, " ", sum > prevSum)

		prevSum = sum
	}

	print(counter)
}
