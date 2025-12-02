package main

import (
	"os"
        "AOC25_Day01/part1"
        "AOC25_Day01/part2"
)

func read_arg() string {
    arg := os.Args[0]
    os.Args = os.Args[1:]
    return arg
}

func main() {
    _ = read_arg()
    part_name := read_arg()
    file_name := read_arg()


    switch part_name {
    case "part1": part1.Process(file_name)
    case "part2": part2.Process(file_name)
    }

}
