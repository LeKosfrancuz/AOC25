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
	input_scanner.Split(bufio.ScanLines)

	available_fresh_food := 0

	parsing := FRESH_IDS
	fresh_id_ranges := make([]Id_range, 0)

	for input_scanner.Scan() {
		line := strings.Trim(input_scanner.Text(), "\n")

		if line == "" {
			parsing = AVAILABLE_IDS
			continue
		}

		if parsing == FRESH_IDS {
			var id_range Id_range
			{
				id_range_text := line
				numbers := strings.Split(id_range_text, "-")
				lower, err := strconv.ParseInt(numbers[0], 10, 64)
				if err != nil {
					log.Fatal("Could not convert a range text (", id_range_text, ") to a number")
				}

				upper, err := strconv.ParseInt(numbers[1], 10, 64)
				if err != nil {
					log.Fatal("Could not convert a range text (", id_range_text, ") to a number")
				}

				id_range = Id_range{
					lower_id: lower,
					upper_id: upper,
				}
			}
			fresh_id_ranges = append(fresh_id_ranges, id_range)

			continue
		}

		if parsing == AVAILABLE_IDS {
			id, _ := strconv.ParseInt(line, 10, 64)
			spoiled := is_food_id_spoiled(id, fresh_id_ranges)

			if !spoiled {
				available_fresh_food++
			}

			continue
		}
	}

	if err := input_scanner.Err(); err != nil {
		log.Fatal("Could not read input file: ", err)
	}

	fmt.Printf("The part 1 password is: %v\n", available_fresh_food)

}

type File_part int
const (
	FRESH_IDS File_part = iota
	AVAILABLE_IDS
)


type Id_range struct {
	lower_id int64
	upper_id int64
}

func is_food_id_spoiled(food_id int64, fresh_food_ids []Id_range) bool {

	for i := range fresh_food_ids {
		if food_id <= fresh_food_ids[i].upper_id && food_id >= fresh_food_ids[i].lower_id {
			return false
		}
	}
	return true
}
