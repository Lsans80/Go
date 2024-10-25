package main

import "fmt"

// Definir una estructura (clase en Java)
type Person struct {
    Name string
    Age  int
}

// Definir una interfaz
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

// Como implementar la interfaz Animal en la estructura Dog
func (d Dog) Speak() string {
    return "Woof!"
}

// Alias de tipo es un nombre alternativo para un tipo existente para hacerlo más legible y fácil de usar.
type MyInt int

// Tipo función es como una functional interface en Java, es una funcionalidad que se puede asignar a una variable.
type Adder func(int, int) int

func main() {
    // Usar la estructura
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p)

    // Usar la interfaz
    var a Animal = Dog{Name: "Buddy"}
    fmt.Println(a.Speak())
	fmt.Println("My dog", a, "says", a.Speak())

    // Usar el alias de tipo
    var x MyInt = 10
    fmt.Println(x)

    // Usar el tipo función, parecido a una lambda en Java.
    var add Adder = func(a, b int) int {
        return a + b
    }
    fmt.Println(add(2, 3))
}