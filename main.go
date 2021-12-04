package main

import (
	"advent-calendar-2021/challenges"
	"fmt"
	"log"
)

func main() {
	// d1c1, err := challenges.Day1Challenge1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 1 Challenge 1: %d\n", d1c1)

	// d1c2, err := challenges.Day1Challenge2()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 1 Challenge 2: %d\n", d1c2)

	// d2c1, err := challenges.Day2Challenge1()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Day 2 Challenge 1: %d\n", d2c1)

	d2c2, err := challenges.Day2Challenge2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2 Challenge 2: %d\n", d2c2)
}
