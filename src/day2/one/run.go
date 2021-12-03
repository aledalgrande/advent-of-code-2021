package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../commands.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var depth int
	var position int

	for scanner.Scan() {
		var command = scanner.Text()
		var parts = strings.Split(command, " ")
		var amount, _ = strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			position = position + amount
		case "up":
			depth = depth - amount
		case "down":
			depth = depth + amount
		}
	}

	fmt.Println(depth * position)
}
