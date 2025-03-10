package main

import "fmt"

type Cat struct {
	Color string
	Age uint8
	Name string
}

func (cat *Cat) Rename(newName string) {
	cat.Name = newName
}

func (cat Cat) RenameV2(newName string) {
	cat.Name = newName
}

func main() {
	cat := Cat{"N1", 12, "Ben"}
	cat.Rename("ten")
	fmt.Println(cat.Name)
	cat.RenameV2("Nine")
	fmt.Println(cat.Name)
}