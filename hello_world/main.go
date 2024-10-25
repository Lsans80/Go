package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hola guapa, aprendamos Go juntas")

	dbDSN := "example"

	// Como un for each en Java, %d es para enteros y %c para caracteres
	for index, char := range dbDSN {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	fmt.Println("El día y la hora es:", time.Now())
}
