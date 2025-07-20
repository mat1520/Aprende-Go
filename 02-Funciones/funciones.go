// Ejemplo: Funciones en Go - Básico a Avanzado
package main

import (
	"fmt"
	"strings"
)

// 1. Función simple sin parámetros ni return
func saludar() {
	fmt.Println("¡Hola desde Go!")
}

// 2. Función con parámetros
func saludarPersona(nombre string) {
	fmt.Printf("¡Hola %s!\n", nombre)
}

// 3. Función con return
func sumar(a, b int) int {
	return a + b
}

// 4. Función con múltiples returns
func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

// 5. Función con returns nombrados
func operacionesBasicas(a, b int) (suma, resta, multiplicacion int) {
	suma = a + b
	resta = a - b
	multiplicacion = a * b
	return // return automático de las variables nombradas
}

// 6. Función variádica (acepta múltiples argumentos)
func calcularPromedio(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	suma := 0.0
	for _, numero := range numeros {
		suma += numero
	}
	return suma / float64(len(numeros))
}

// 7. Función que recibe otra función como parámetro
func aplicarOperacion(a, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

// 8. Función anónima (closure)
func obtenerContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// 9. Función recursiva
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 10. Función con defer (se ejecuta al final)
func ejemploDefer() {
	fmt.Println("Inicio de la función")
	defer fmt.Println("Esto se ejecuta al final (defer 1)")
	defer fmt.Println("Esto se ejecuta penúltimo (defer 2)")
	fmt.Println("Mitad de la función")
	fmt.Println("Final de la función")
}

func main() {
	fmt.Println("=== FUNCIONES EN GO ===")

	// 1. Función simple
	fmt.Println("\n1. Función simple:")
	saludar()

	// 2. Función con parámetros
	fmt.Println("\n2. Función con parámetros:")
	saludarPersona("Carlos")

	// 3. Función con return
	fmt.Println("\n3. Función con return:")
	resultado := sumar(15, 25)
	fmt.Printf("15 + 25 = %d\n", resultado)

	// 4. Múltiples returns
	fmt.Println("\n4. Múltiples returns:")
	cociente, err := dividir(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", cociente)
	}

	// 5. Returns nombrados
	fmt.Println("\n5. Returns nombrados:")
	s, r, m := operacionesBasicas(8, 3)
	fmt.Printf("8 + 3 = %d, 8 - 3 = %d, 8 * 3 = %d\n", s, r, m)

	// 6. Función variádica
	fmt.Println("\n6. Función variádica:")
	promedio := calcularPromedio(8.5, 9.0, 7.5, 6.0, 8.0)
	fmt.Printf("Promedio de notas: %.2f\n", promedio)

	// 7. Función como parámetro
	fmt.Println("\n7. Función como parámetro:")
	multiplicar := func(x, y int) int {
		return x * y
	}
	resultado = aplicarOperacion(6, 4, multiplicar)
	fmt.Printf("6 * 4 = %d\n", resultado)

	// 8. Closure
	fmt.Println("\n8. Closure (función anónima):")
	contador := obtenerContador()
	fmt.Printf("Contador: %d\n", contador())
	fmt.Printf("Contador: %d\n", contador())
	fmt.Printf("Contador: %d\n", contador())

	// 9. Recursión
	fmt.Println("\n9. Recursión:")
	fmt.Printf("Factorial de 5: %d\n", factorial(5))

	// 10. Defer
	fmt.Println("\n10. Defer:")
	ejemploDefer()

	// BONUS: Función con strings
	fmt.Println("\n=== BONUS: Manipulación de strings ===")
	procesarTexto("Hola Mundo con Go")
}

func procesarTexto(texto string) {
	fmt.Printf("Texto original: %s\n", texto)
	fmt.Printf("Mayúsculas: %s\n", strings.ToUpper(texto))
	fmt.Printf("Minúsculas: %s\n", strings.ToLower(texto))
	fmt.Printf("Longitud: %d caracteres\n", len(texto))
	fmt.Printf("Contiene 'Go': %t\n", strings.Contains(texto, "Go"))
}
