// Ejemplo 1: Variables y Tipos Básicos en Go
package main

import "fmt"

func main() {
	fmt.Println("=== VARIABLES Y TIPOS BÁSICOS ===")

	// 1. Declaración con var (forma explícita)
	var nombre string = "Ana García"
	var edad int = 25
	var altura float64 = 1.65
	var esEstudiante bool = true

	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Edad: %d años\n", edad)
	fmt.Printf("Altura: %.2f metros\n", altura)
	fmt.Printf("¿Es estudiante?: %t\n", esEstudiante)

	// 2. Declaración corta (inferencia de tipo)
	ciudad := "Madrid"
	temperatura := 22.5
	llueve := false

	fmt.Printf("\nCiudad: %s\n", ciudad)
	fmt.Printf("Temperatura: %.1f°C\n", temperatura)
	fmt.Printf("¿Llueve?: %t\n", llueve)

	// 3. Múltiples variables
	var x, y, z int = 1, 2, 3
	a, b, c := "Go", "es", "genial"

	fmt.Printf("\nNúmeros: %d, %d, %d\n", x, y, z)
	fmt.Printf("Palabras: %s %s %s\n", a, b, c)

	// 4. Variables con valor cero (zero values)
	var numeroEntero int      // 0
	var textoVacio string     // ""
	var booleanoFalso bool    // false
	var numeroDecimal float64 // 0.0

	fmt.Printf("\nValores por defecto:\n")
	fmt.Printf("int: %d\n", numeroEntero)
	fmt.Printf("string: '%s'\n", textoVacio)
	fmt.Printf("bool: %t\n", booleanoFalso)
	fmt.Printf("float64: %f\n", numeroDecimal)

	// 5. Constantes
	const PI = 3.14159
	const EMPRESA = "MiCompany"

	fmt.Printf("\nConstantes:\n")
	fmt.Printf("PI: %f\n", PI)
	fmt.Printf("Empresa: %s\n", EMPRESA)
}
