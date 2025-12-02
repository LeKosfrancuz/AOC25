package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Process(file_name string) {

    fmt.Println("    ---- Part 2 ----    ")
    fmt.Printf("Reading file `%v`\n", file_name)

    input, err := os.Open(file_name)
    if err != nil {
        log.Fatal("Could not open input file: ", err)
    }

    defer input.Close()


    input_scanner := bufio.NewScanner(input)

    count_of_0_passes := 0
    dial_state := 50

    for input_scanner.Scan() {
        line := input_scanner.Text()

        count_of_0_passes += do_dial_action(&dial_state, line)

        // fmt.Println(line)
        // fmt.Printf("state: %v, passes: %v\n", dial_state, count_of_0_passes)
        // fmt.Println("-------------------")
    }

    if err := input_scanner.Err(); err != nil {
        log.Fatal("Could not read input file: ", err)
    }

    fmt.Printf("The part 2 password is: %v\n", count_of_0_passes)

}


const (
    I_Sub byte = 'L'
    I_Add byte = 'R'
)

const COUNT_STATES int = 100

func do_dial_action(dial_state *int, line string) (num_of_passes_over_0 int) {
    move_ammount, err := strconv.Atoi(line[1:])
    if err != nil {
        log.Fatal("In line: `", line, "` error at converting to int: ", err)
    }

    prev_state := *dial_state
    passes := move_ammount/COUNT_STATES

    switch line[0] {
    case I_Add:
        *dial_state = (*dial_state + move_ammount)%COUNT_STATES
        if prev_state > *dial_state && prev_state != 0 {
            passes += 1
        }

        return passes
    case I_Sub:
        *dial_state = (*dial_state + (passes + 1)*COUNT_STATES - move_ammount)%COUNT_STATES
        if prev_state < *dial_state && prev_state != 0 {
            passes += 1
        }
    }

    if *dial_state == 0 && prev_state != 0 {
        passes += 1
    }

    return passes
}
