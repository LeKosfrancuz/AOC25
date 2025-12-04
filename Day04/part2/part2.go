package part2

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const DEBUG = false

func Process(file_name string) {

	fmt.Println("	 ---- Part 2 ----	 ")
	fmt.Printf("Reading file `%v`\n", file_name)

	input, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal("Could not open input file: ", err)
	}

	roll_map := strings.Split(strings.Trim(string(input), "\n"), "\n")


	count_removed := 0
	removed := 0
	first_loop := true

	for removed > 0 || first_loop {
		first_loop = false
		removed = remove_accessible(roll_map)

		if removed > 0 {
			count_removed += removed
		}
	}

	fmt.Printf("The part 2 password is: %v\n", count_removed)
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

func remove_accessible(roll_map []string) int {

	removed := 0
	roll_map_before_removing := make([]string, len(roll_map))
	{
		copied_items := copy(roll_map_before_removing, roll_map)
		if copied_items != len(roll_map) {
			log.Fatalf("Making a copy of the board FAILED!\nTried to copy %v elements but succeded in %v\n", len(roll_map), copied_items)
		}
	}

	for row := range roll_map_before_removing {
		for col := range roll_map_before_removing[row] {
			current_tile := roll_map_before_removing[row][col]

			if current_tile == TILE_EMPTY { continue }

			count_adjecent := count_adjecent_roles_of_paper(roll_map_before_removing, row, col)

			if count_adjecent < 4 {
				removed += 1

				// Ugly way to replace the col elementh with an empty tile
				roll_map[row] = roll_map[row][0:col] + string(TILE_EMPTY) + roll_map[row][col+1:]
			}
		}


	}

	return removed
}
