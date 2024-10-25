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

// askData solicita al usuario la contraseña, host y puerto de la base de datos.
func askData() (string, string, string) {
	var password, host, port string

	fmt.Println("Enter password")
	fmt.Scan(&password)

	fmt.Println("Enter host")
	fmt.Scan(&host)

	fmt.Println("Enter port")
	fmt.Scan(&port)

	return password, host, port
}

// updateDSN actualiza el DSN de la base de datos con los valores de password, host y port.
func updateDSN(dbDSN string) {
	if dbDSN == "" {
		log.Fatalf("Error updating DSN: %v", errors.New("empty DSN"))
	}

	password, host, port := askData()

	//Elimina todos los espacios en el valor del string dsn.
	dbDSNNoSpaces := strings.ReplaceAll(dbDSN, " ", "")

	//Sustituye PASSWORD, HOST y PORT por valores inventados pero válidos.
	dbDSNUpdated := strings.ReplaceAll(dbDSNNoSpaces, "PASSWORD", password)
	dbDSNUpdated = strings.ReplaceAll(dbDSNUpdated, "HOST", host)
	dbDSNUpdated = strings.ReplaceAll(dbDSNUpdated, "PORT", port)

	//Muestra el resultado de dbDSN
	fmt.Println("Updated DB DSN:", dbDSNUpdated)
}
