package main

import (
	"bufio"
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

	m2.WriteString(string(data))

	//Method three
	m3, err := os.Create("m3.tx")
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bufio.NewWriter(m3)    //opens the file for writing
	w.WriteString(string(data)) //writes the data
	err = w.Flush()
	if err != nil {
		return
	}

}
