package main

import (
	"fmt"
	"os"
)

func main() {

	data := []byte("Data to write \n")

	//Method one
	m1, err := os.Create("m1.txt")
	if err != nil {
		fmt.Println("Cannot create file ", err)
		return
	}
	defer m1.Close()

	fmt.Fprintf(m1, string(data))

}
