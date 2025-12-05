set -e

go build main.go

./main part2 input.in

rm main
