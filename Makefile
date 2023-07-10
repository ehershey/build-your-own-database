main: *.go
	go build -o main main.go btree.go  tests.go  savedata.go

run: main
	./main
