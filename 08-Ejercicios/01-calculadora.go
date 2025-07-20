// Ejercicio 1: Calculadora Básica
// Objetivo: Crear una calculadora que tome dos números y una operación desde la consola
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== CALCULADORA BÁSICA EN GO ===")

	// Verificar que se pasen los argumentos correctos
	if len(os.Args) != 4 {
		fmt.Println("Uso: go run calculadora.go <numero1> <operacion> <numero2>")
		fmt.Println("Operaciones disponibles: +, -, *, /")
		fmt.Println("Ejemplo: go run calculadora.go 10 + 5")
		return
	}

	// Obtener argumentos de la línea de comandos
	num1Str := os.Args[1]
	operacion := os.Args[2]
	num2Str := os.Args[3]

	// Convertir strings a números
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil {
		fmt.Printf("Error: '%s' no es un número válido\n", num1Str)
		return
	}

	if err2 != nil {
		fmt.Printf("Error: '%s' no es un número válido\n", num2Str)
		return
	}

	// Realizar la operación
	var resultado float64
	var error bool

	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: División por cero no permitida")
			error = true
		} else {
			resultado = num1 / num2
		}
	default:
		fmt.Printf("Error: Operación '%s' no válida. Usa +, -, *, /\n", operacion)
		error = true
	}

	// Mostrar resultado
	if !error {
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operacion, num2, resultado)
	}
}
