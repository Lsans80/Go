package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("Hello Luisa")

	dbDSN := "imbee:PASSWORD@tcp(HOST:PORT)/db1?charset=utf8mb4&parseTime=true   "
	updateDSN(dbDSN)

}

func updateDSN(dbDSN string) {
	if dbDSN == "" {
		log.Fatalf("Error updating DSN: %v", errors.New("empty DSN"))
	}

	// TODO: 1. Elimina todos los espacios en el valor del string dsn.
	dbDSNNoSpaces := strings.ReplaceAll(dbDSN, " ", "")

	// TODO: 2. Sustituye PASSWORD, HOST y PORT por valores inventados pero válidos.
	dbDSNUpdated := strings.ReplaceAll(dbDSNNoSpaces, "PASSWORD", "pass123")
	dbDSNUpdated = strings.ReplaceAll(dbDSNUpdated, "HOST", "localhost")
	dbDSNUpdated = strings.ReplaceAll(dbDSNUpdated, "PORT", "3306")

	// TODO: 3. Muestra el resultado de dbDSN
	fmt.Println("Updated DB DSN:", dbDSNUpdated)
}
