set -e

go build main.go

./main both test.in

rm main
