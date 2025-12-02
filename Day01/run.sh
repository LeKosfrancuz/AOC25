set -e

go build main.go

./main part1 input.in

echo

./main part2 input.in

rm main
