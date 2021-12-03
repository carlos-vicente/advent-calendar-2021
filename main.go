package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"regexp"
)

func main() {
	// d1c1, err := day1Challenge1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 1 Challenge 1: %d\n", d1c1)

	// d1c2, err := day1Challenge2()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 1 Challenge 2: %d\n", d1c2)

	// d2c1, err := day2Challenge1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 2 Challenge 1: %d\n", d2c1)

	d2c2, err := day2Challenge2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2 Challenge 2: %d\n", d2c2)
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func day1Challenge1() (int, error) {
	file, err := os.Open("./inputs/day_1.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastValue := -1
	counter := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		currentValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if lastValue != -1 {
			if currentValue > lastValue {
				counter++
			}
		}
		lastValue = currentValue

		fmt.Printf("value: %d | counter: %d\n", currentValue, counter)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func day1Challenge2() (int, error) {
	values, err := readLines("./inputs/day_1.txt")
	if err != nil {
		return 0, err
	}

	lastWindow := -1
	counter := 0

	for index := range values {
		if index >= 2 {
			window := values[index-2] + values[index-1] + values[index]

			if lastWindow != -1 {
				if window > lastWindow {
					counter++
				}
			}
			lastWindow = window
			fmt.Printf("value: %d | counter: %d\n", window, counter)
		}
	}

	return counter, nil
}


func day2Challenge1() (int, error) {
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


func day2Challenge2() (int, error) {
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