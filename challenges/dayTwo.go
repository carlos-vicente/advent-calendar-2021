package challenges

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day2Challenge1() (int, error) {
	file, err := os.Open("./inputs/day_2.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontal := 0
	depth := 0

	reg := regexp.MustCompile(`(?P<command>\w*)\s(?P<value>\d*)`)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		matches := reg.FindStringSubmatch(line)

		switch command := matches[1]; command {
		case "forward":
			value, _ := strconv.Atoi(matches[2])
			horizontal += value
			fmt.Println("f", horizontal)
		case "up":
			value, _ := strconv.Atoi(matches[2])
			depth -= value
			fmt.Println("u", depth)
		case "down":
			value, _ := strconv.Atoi(matches[2])
			depth += value
			fmt.Println("d", depth)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return horizontal * depth, nil
}

func Day2Challenge2() (int, error) {
	file, err := os.Open("./inputs/day_2.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontal := 0
	depth := 0
	aim := 0

	reg := regexp.MustCompile(`(?P<command>\w*)\s(?P<value>\d*)`)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		matches := reg.FindStringSubmatch(line)

		switch command := matches[1]; command {
		case "forward":
			value, _ := strconv.Atoi(matches[2])
			horizontal += value
			depth += aim * value
			fmt.Printf("f %d %d\n", horizontal, depth)
		case "up":
			value, _ := strconv.Atoi(matches[2])
			aim -= value
			fmt.Printf("u %d\n", aim)
		case "down":
			value, _ := strconv.Atoi(matches[2])
			aim += value
			fmt.Printf("d %d\n", aim)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return horizontal * depth, nil
}