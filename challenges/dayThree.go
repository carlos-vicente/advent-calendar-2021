package challenges

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3Challenge1() (int64, error) {
	file, err := os.Open("./inputs/day_3.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	var columnsWithOne []int
	
	for scanner.Scan() {
		line := scanner.Text()

		for pos, char := range line {
			if pos == len(columnsWithOne) {
				columnsWithOne = append(columnsWithOne, 0)
				fmt.Println("adding column")
			}

			c := string(char)
			
			if c == "1"{
				columnsWithOne[pos]++
			}
		}

		lines++
	}

	fmt.Println(columnsWithOne)
	var gama string
	var epsilon string
	for _, counter := range columnsWithOne{
		if counter > (lines/2){
			fmt.Println("most comum is 1")
			gama += "1"
			epsilon += "0"
		}else{
			fmt.Println("most comum is 0")
			gama += "0"
			epsilon += "1"
		}
	}

	fmt.Printf("number of lines: %d\n", lines)
	fmt.Println(gama)
	fmt.Println(epsilon)

	gamaInt, err := strconv.ParseInt(gama, 2, 64)
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gamaInt)
	fmt.Println(epsilonInt)

	return gamaInt*epsilonInt, nil
}