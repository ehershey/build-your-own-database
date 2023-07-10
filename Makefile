main: *.go
	go build -o main *.go

run: main
	./main
