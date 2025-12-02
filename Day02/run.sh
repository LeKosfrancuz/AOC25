set -e

go build main.go

./main both input.in

rm main
