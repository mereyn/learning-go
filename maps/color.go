package main

import "fmt"

type colors map[string]string

func (c colors) printMap() {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
