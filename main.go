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

	//Method two
	m2, err := os.Create("m2.txt")
	if err != nil {
		fmt.Println("Cannot create file ", err)
		return
	}
	defer m2.Close()

	n, err := m2.WriteString(string(data))
	fmt.Printf("wrote %d bytes\n", n)

}
