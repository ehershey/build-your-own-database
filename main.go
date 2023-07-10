package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calling SaveData1")
	err := SaveData1("/tmp/somedata", []byte("There are these two young fish swimming along and they happen to meet an older fish swimming the other way"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("Calling SaveData2")
	err = SaveData2("/tmp/somedata", []byte("There are these two young fish swimming along and they happen to meet an older fish swimming the other way"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("Calling SaveData3")
	err = SaveData3("/tmp/somedata", []byte("There are these two young fish swimming along and they happen to meet an older fish swimming the other way"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("Data saved in /tmp/somedata")

	fmt.Println("Creating Btree test struct")
	c := newC()
	fmt.Println(c)
	c.add("name", "Ernie")
	fmt.Println(c)
}
