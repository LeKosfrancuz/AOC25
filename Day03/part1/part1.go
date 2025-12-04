package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func Process(file_name string) {

	fmt.Println("	 ---- Part 1 ----	 ")
	fmt.Printf("Reading file `%v`\n", file_name)

	input, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	defer input.Close()


	input_scanner := bufio.NewScanner(input)

	sum_max_battery_combos := 0

	for input_scanner.Scan() {
		line := strings.Trim(input_scanner.Text(), "\n")

		combo := find_max_battery_combo_in_line(line)

		// Sum all the max battery combinations
		sum_max_battery_combos += combo
	}

	if err := input_scanner.Err(); err != nil {
		log.Fatal("Could not read input file: ", err)
	}

	fmt.Printf("The part 1 password is: %v\n", sum_max_battery_combos)

}


type Id_range struct {
	lower_id int
	upper_id int
}

func find_max_battery(line string, start, end int) (max_battery, index_of_max int, _ error) {

	max_battery = 0
	index_of_max = -1

	for i := start; i < len(line) && i < end; i++ {
		battery, err := strconv.Atoi(string(line[i]))
		if err != nil {
			return 0, 0, fmt.Errorf("Line contains non-number characters!")
		}

		if battery > max_battery {
			max_battery = battery
			index_of_max = i
		}
	}

	return max_battery, index_of_max, nil
}

func find_max_battery_combo_in_line(line string) int {

	higher, index_of_higher, err := find_max_battery(line, 0, len(line) - 1)
	if err != nil {
		log.Fatal("Could not get a battery number: ", err)
	}

	lower, _, err := find_max_battery(line, index_of_higher + 1, len(line))

	if DEBUG {
		fmt.Printf("Battery pack: %v\n", line)
		fmt.Printf("Max battery: %v\n", higher * 10 + lower)
	}


	return higher * 10 + lower
}
