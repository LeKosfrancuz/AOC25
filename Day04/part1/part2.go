package part1

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const DEBUG = false

func Process(file_name string) {

	fmt.Println("	 ---- Part 1 ----	 ")
	fmt.Printf("Reading file `%v`\n", file_name)

	input, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	roll_map := strings.Split(strings.Trim(string(input), "\n"), "\n")


	count_accessible := 0

	for row := range roll_map {
		for col := range roll_map[row] {
			current_tile := roll_map[row][col]

			if current_tile == TILE_EMPTY { continue }

			if is_roll_accessible(roll_map, row, col) {
				count_accessible++
			}
		}


	}

	fmt.Printf("The part 1 password is: %v\n", count_accessible)
}


const (
	TILE_EMPTY = '.'
	TILE_ROLL  = '@'
)

func count_adjecent_roles_of_paper(roll_map []string, row, col int) (count_adjecent int) {

	count_adjecent = 0

	last_row := len(roll_map) - 1
	last_col := len(roll_map[0]) - 1

	// Check top
	if row != 0 {

		// Check left
		if col != 0 {
			if roll_map[row - 1][col - 1] == TILE_ROLL { count_adjecent++ }
		}

		// Check middle
		if roll_map[row - 1][col] == TILE_ROLL { count_adjecent++ }

		// Check right
		if col != last_col {
			if roll_map[row - 1][col + 1] == TILE_ROLL { count_adjecent++ }
		}
	}

	// Check bottom
	if row != last_row {

		// Check left
		if col != 0 {
			if roll_map[row + 1][col - 1] == TILE_ROLL { count_adjecent++ }
		}

		// Check middle
		if roll_map[row + 1][col] == TILE_ROLL { count_adjecent++ }

		// Check right
		if col != last_col {
			if roll_map[row + 1][col + 1] == TILE_ROLL { count_adjecent++ }
		}
	}

	//Check middle

	// Check left
	if col != 0 {
		if roll_map[row][col - 1] == TILE_ROLL { count_adjecent++ }
	}

	// Check right
	if col != last_col {
		if roll_map[row][col + 1] == TILE_ROLL { count_adjecent++ }
	}


	return count_adjecent
}

func is_roll_accessible(roll_map []string, row, col int) bool {

	count_adjecent := count_adjecent_roles_of_paper(roll_map, row, col)

	if count_adjecent < 4 {
		return true
	}


	// if DEBUG {
	// 	fmt.Printf("Battery pack: %v\n", line)
	// 	fmt.Printf("Max battery: %v\n", higher * 10 + lower)
	// }


	return false
}
