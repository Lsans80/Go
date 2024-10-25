package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	values := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6}

	// Calcular y mostrar la suma total
	totalSum := sumValues(values)
	fmt.Println("Suma total:", totalSum)

	// Filtrar y mostrar los elementos impares
	oddValues := filterOddValues(values)
	fmt.Println("Elementos impares:", oddValues)

	// Obtener y mostrar elementos únicos ordenados de mayor a menor
	sortedUniqueValues := uniqueSortedValues(values)
	fmt.Println("Elementos únicos ordenados de mayor a menor:", sortedUniqueValues)

}

func sumValues(values []int) int {
	if values == nil {
		log.Fatal("Error: slice is empty")
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func filterOddValues(values []int) []int {
    if values == nil {
		log.Fatal("Error: slice is empty")
	}
	var oddValues []int
	for _, v := range values {
		if v%2 != 0 {
			oddValues = append(oddValues, v)
		}
	}
	return oddValues
}

func uniqueValues(values []int) []int {
	if values == nil {
		log.Fatal("Error: slice is empty")
	}
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

	return uniqueValues
}

func uniqueSortedValues(values []int) []int {
    if values == nil {
		log.Fatal("Error: slice is empty")
	}
    uniqueSortedValues := uniqueValues(values)
    sort.Sort(sort.Reverse(sort.IntSlice(uniqueSortedValues)))
	return uniqueSortedValues
}
