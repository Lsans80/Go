package main

import "fmt"

func main() {
	values := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6}

	// TODO: 1. Calcula y muestra el resultado de la suma total de todos los elementos del slice

	totalSum := sumValues(values)
	fmt.Println("Suma total:", totalSum)

	// TODO: 2. Crea un nuevo slice que contenga solo los elementos impares

	oddValues := filterOddValues(values)
	fmt.Println("Elementos impares:", oddValues)

	// TODO: 3. Crea un nuevo slice con elementos únicos ordenados de mayor a menor

	sortedUniqueValues := uniqueSortedValues(values)
	fmt.Println("Elementos únicos ordenados de mayor a menor:", sortedUniqueValues)

}

// Sumar todos los elementos del slice

func sumValues(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

//Filtrar elementos impares
//range itera sobre los elementos de un slice devolviendo el índice y el valor de cada elemento en cada iteración.
//Como un for each en Java.

func filterOddValues(values []int) []int {
	var oddValues []int
	for _, v := range values {
		if v%2 != 0 {
			oddValues = append(oddValues, v)
		}
	}
	return oddValues
}

// Obtener elementos únicos

func uniqueValues(values []int) []int {

	// Un map es una colección de pares clave-valor, donde cada clave es única.
	// En este caso, usamos un map para almacenar los elementos únicos del slice. Ej. map[1:true, 2:true, 3:true]
	uniqueMap := make(map[int]bool)
	for _, v := range values {
		uniqueMap[v] = true
	}

	// Convertir las claves del map en un slice de elementos únicos y almacenarlos en un slice de enteros.
	var uniqueValues []int
	for k := range uniqueMap {
		uniqueValues = append(uniqueValues, k)
	}

	fmt.Println("Mapa de elementos únicos:", uniqueValues)

	return uniqueValues
}

// Ordenar elementos únicos de mayor a menor

func uniqueSortedValues(values []int) []int {
	uniqueValues := uniqueValues(values)

	for i := 0; i < len(uniqueValues)-1; i++ {
		for j := i + 1; j < len(uniqueValues); j++ {
			if uniqueValues[i] < uniqueValues[j] {
				//En Go no es necesaria una variable auxiliar para intercambiar valores. 
				//Se puede hacer directamente por que Go evalua primero el valor de la derecha.
				uniqueValues[i], uniqueValues[j] = uniqueValues[j], uniqueValues[i]
			}
		}
	}

	return uniqueValues
}
