package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Process(file_name string) {

    fmt.Println("    ---- Part 1 ----    ")
    fmt.Printf("Reading file `%v`\n", file_name)

    input, err := os.Open(file_name)
    if err != nil {
        log.Fatal("Could not open input file: ", err)
    }

    defer input.Close()


    input_scanner := bufio.NewScanner(input)

    count_of_0_stops := 0
    dial_state := 50

    for input_scanner.Scan() {
        line := input_scanner.Text()

        do_dial_action(&dial_state, line)
        if dial_state == 0 {
            count_of_0_stops++
        }
    }

    if err := input_scanner.Err(); err != nil {
        log.Fatal("Could not read input file: ", err)
    }

    fmt.Printf("The part 1 password is: %v\n", count_of_0_stops)

}


const (
    I_Sub byte = 'L'
    I_Add byte = 'R'
)

const COUNT_STATES int = 100

func do_dial_action(dial_state *int, line string) {
    move_ammount, err := strconv.Atoi(line[1:])
    if err != nil {
        log.Fatal("In line: `", line, "` error at converting to int: ", err)
    }

    switch line[0] {
    case I_Add:
        *dial_state = (*dial_state + move_ammount)%COUNT_STATES
    case I_Sub:
        *dial_state = (*dial_state + COUNT_STATES - move_ammount)%COUNT_STATES
    }
}
