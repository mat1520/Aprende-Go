// Ejemplo práctico de funciones variádicas
package main

import "fmt"

// Función variádica para calcular promedio
func calcularPromedio(numeros ...float64) float64 {
	fmt.Printf("📊 Recibí %d números: %v\n", len(numeros), numeros)

	if len(numeros) == 0 {
		fmt.Println("⚠️  No hay números para calcular")
		return 0
	}

	suma := 0.0
	for _, numero := range numeros {
		suma += numero
	}

	promedio := suma / float64(len(numeros))
	fmt.Printf("➕ Suma: %.2f | ➗ Promedio: %.2f\n", suma, promedio)
	return promedio
}

// Función variádica para concatenar strings
func concatenar(separador string, palabras ...string) string {
	fmt.Printf("📝 Concatenando %d palabras con '%s'\n", len(palabras), separador)

	resultado := ""
	for i, palabra := range palabras {
		if i > 0 {
			resultado += separador
		}
		resultado += palabra
	}
	return resultado
}

// Función variádica para encontrar el máximo
func encontrarMaximo(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}

	maximo := numeros[0]
	for _, numero := range numeros {
		if numero > maximo {
			maximo = numero
		}
	}
	return maximo
}

func main() {
	fmt.Println("=== FUNCIONES VARIÁDICAS EXPLICADAS ===")

	// 1. PROMEDIO CON DIFERENTES CANTIDADES
	fmt.Println("\n1. 📊 Calculando promedios:")

	calcularPromedio(8.5)                  // 1 número
	calcularPromedio(8.5, 9.0, 7.5)        // 3 números
	calcularPromedio(10, 8, 9, 7, 6, 8, 9) // 7 números
	calcularPromedio()                     // 0 números

	// 2. CONCATENACIÓN DE STRINGS
	fmt.Println("\n2. 📝 Concatenando palabras:")

	resultado1 := concatenar(" ", "Hola", "mundo", "desde", "Go")
	fmt.Printf("Resultado: '%s'\n", resultado1)

	resultado2 := concatenar("-", "HTML", "CSS", "JavaScript", "Go")
	fmt.Printf("Resultado: '%s'\n", resultado2)

	resultado3 := concatenar(", ", "manzana", "banana", "naranja")
	fmt.Printf("Resultado: '%s'\n", resultado3)

	// 3. ENCONTRAR MÁXIMO
	fmt.Println("\n3. 🔢 Encontrando máximos:")

	max1 := encontrarMaximo(3, 7, 2, 9, 1)
	fmt.Printf("Máximo de [3,7,2,9,1]: %d\n", max1)

	max2 := encontrarMaximo(100)
	fmt.Printf("Máximo de [100]: %d\n", max2)

	max3 := encontrarMaximo(50, 25, 75, 12, 88, 33)
	fmt.Printf("Máximo de [50,25,75,12,88,33]: %d\n", max3)

	// 4. EJEMPLO CON SLICE EXISTENTE
	fmt.Println("\n4. 📋 Usando slice existente:")

	notas := []float64{8.5, 9.2, 7.8, 9.0, 8.7}
	fmt.Println("Notas del estudiante:", notas)

	// Para pasar un slice, usar ... al final
	promedio := calcularPromedio(notas...)
	fmt.Printf("El promedio final es: %.2f\n", promedio)

	// 5. VENTAJAS DE FUNCIONES VARIÁDICAS
	fmt.Println("\n5. ✅ Ventajas:")
	fmt.Println("   • Flexibilidad: acepta cualquier cantidad de argumentos")
	fmt.Println("   • Simplicidad: una sola función para diferentes casos")
	fmt.Println("   • Legibilidad: código más limpio y expresivo")

	fmt.Println("\n6. 📚 Ejemplos de Go que las usan:")
	fmt.Println("   • fmt.Printf(formato, args...)")
	fmt.Println("   • fmt.Println(args...)")
	fmt.Println("   • append(slice, elementos...)")
}
