package main

import "fmt"

func main() {
	fmt.Println("Hola Luisa")

	// Actualiza la variable dbDSN (data source name de una base de datos) cumpliendo las condiciones:
	// TODO: 1. Elimina todos los espacios en el valor del string dsn.
	// TODO: 2. Sustituye PASSWORD, HOST y PORT por valores inventados pero válidos.
	// TODO: 3. Muestra el resultado de dbDSN

	dbDSN := "imbee:PASSWORD@tcp(HOST:PORT)/db1?charset=utf8mb4&parseTime=true   "
	// TODO 1, 2 y 3.
	dbDSNUpdated(dbDSN)

}

// TODO: 1. Elimina todos los espacios en el valor del string dsn.

func dbDSNRemoveSpaces(dbDSN string) string {
	var dbDSNNoSpaces string
	for _, char := range dbDSN {
		if char != ' ' {
			dbDSNNoSpaces += string(char)
		}
	}
	return dbDSNNoSpaces
}

// TODO: 2. Sustituye PASSWORD, HOST y PORT por valores inventados pero válidos.
// TODO: 3. Muestra el resultado de dbDSN

func dbDSNUpdated(dbDSN string) {
	dbDSNNoSpaces := dbDSNRemoveSpaces(dbDSN) // Elimina los espacios aquí
	dbDSNUpdated := ""                        // Variable para almacenar el resultado

	/*Verifica que haya al menos n caracteres restantes en la cadena dbDSNNoSpaces a partir del índice i.
	Esto asegura que la operación de slicing dbDSNNoSpaces[i:i+n] no cause un error de índice fuera de rango.
	*/
	for i := 0; i < len(dbDSNNoSpaces); i++ { // Itera sobre cada caracter de dbDSNNoSpaces
		if i+8 < len(dbDSNNoSpaces) && dbDSNNoSpaces[i:i+8] == "PASSWORD" { // Corrige el tamaño de "PASSWORD"
			dbDSNUpdated += "pass123"
			i += 7 // Salta 7 posiciones porque en la siguiente iteración se incrementará en 1.
		} else if i+4 < len(dbDSNNoSpaces) && dbDSNNoSpaces[i:i+4] == "HOST" {
			dbDSNUpdated += "localhost"
			i += 3 // Salta el tamaño de "HOST"
		} else if i+4 < len(dbDSNNoSpaces) && dbDSNNoSpaces[i:i+4] == "PORT" {
			dbDSNUpdated += "3306"
			i += 3 // Salta el tamaño de "PORT"
		} else {
			dbDSNUpdated += string(dbDSNNoSpaces[i]) // Agrega el caracter de dbDSNNoSpaces
		}
	}
	fmt.Println("DB DSN actualizado:", dbDSNUpdated)
}
