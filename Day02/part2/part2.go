package part2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const DEBUG = false


func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r != ' ' {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func Process(file_name string) {

	fmt.Println("	 ---- Part 2 ----	 ")
	fmt.Printf("Reading file `%v`\n", file_name)

	input, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	defer input.Close()


	input_scanner := bufio.NewScanner(input)
	input_scanner.Split(scanCommas)

	invalid_id_sum := 0

	for input_scanner.Scan() {
		// Parse and create the id_range struct
		var id_range Id_range
		{
			id_range_text := strings.Trim(input_scanner.Text(), "\n")
			numbers := strings.Split(id_range_text, "-")
			lower, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatal("Could not convert a range text (", id_range_text, ") to a number")
			}

			upper, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatal("Could not convert a range text (", id_range_text, ") to a number")
			}

			id_range = Id_range{
				lower_id: lower,
				upper_id: upper,
			}
		}

		ids := find_invalid_ids_in_range(id_range)

		// Sum all the invalid ids
		for i := range ids {
			invalid_id_sum += ids[i]
		}
		if DEBUG {
			fmt.Println("----------------------------------------------")
		}
	}

	if err := input_scanner.Err(); err != nil {
		log.Fatal("Could not read input file: ", err)
	}

	fmt.Printf("The part 2 password is: %v\n", invalid_id_sum)

}


type Id_range struct {
	lower_id int
	upper_id int
}

func reverseASCII(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

func is_id_invalid(id int) bool {

	// Starts at number lenght
	repetition_len := int(math.Floor(math.Log10(float64(id)))) + 1
	str_id := make([]string, 0)
	for i := 0; i < repetition_len; i++ {
		str_id = append(str_id, strconv.Itoa(id % 10))
		id /= 10
	}
	str_id = reverseASCII(str_id)
	if DEBUG {
		fmt.Printf("ID: %v\n", str_id)
	}

	if repetition_len == 1 {
		return false
	}

	repetition_len /= 2

	// Check repetition for every repetition length
	for repetition_len > 0 {
		pattern := str_id[0:repetition_len]
		if DEBUG {
			fmt.Printf("Pattern: %v\n", pattern)
		}


		is_invalid := true
		for i := 0; i < len(str_id); i += repetition_len {
			if i + repetition_len - 1 >= len(str_id) {
				is_invalid = false
				break
			}
			for j := 0; j < repetition_len; j++ {
				if str_id[i + j] != pattern[j] {
					is_invalid = false
				}
			}
		}

		if is_invalid {
			if DEBUG {
				fmt.Printf("Pattern: REPEATING\n")
			}
			return true
		}

		repetition_len -= 1
	}

	return false
}

func find_invalid_ids_in_range(id_range Id_range) []int {

	invalid_ids := make([]int, 0)

	for id := id_range.lower_id; id <= id_range.upper_id; id++ {
		if is_id_invalid(id) {
			invalid_ids = append(invalid_ids, id)
		}
	}

	return invalid_ids
}
