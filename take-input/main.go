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
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(line)
	}

	// data := make([]byte, 100)

	// _, err = reader.Read(data)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(data))

}