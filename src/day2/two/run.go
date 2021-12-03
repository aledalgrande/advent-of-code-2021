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
	var aim int

	for scanner.Scan() {
		var command = scanner.Text()
		var parts = strings.Split(command, " ")
		var amount, _ = strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			depth = depth + aim*amount
			position = position + amount
		case "up":
			aim = aim - amount
		case "down":
			aim = aim + amount
		}

		fmt.Println(command, "depth: ", depth, "position: ", position, "aim: ", aim)
	}

	fmt.Println(depth * position)
}
