package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)

	data := make([]byte, 100)

	_, err = reader.Read(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

}