// Ejemplo 2: Control de Flujo - Condicionales y Bucles
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== CONTROL DE FLUJO ===")

	// 1. CONDICIONALES IF-ELSE
	fmt.Println("\n1. Condicionales:")

	edad := 20
	if edad >= 18 {
		fmt.Printf("Con %d años eres mayor de edad\n", edad)
	} else {
		fmt.Printf("Con %d años eres menor de edad\n", edad)
	}

	// If con inicialización
	if puntuacion := 85; puntuacion >= 90 {
		fmt.Println("Excelente!")
	} else if puntuacion >= 70 {
		fmt.Println("Muy bien!")
	} else if puntuacion >= 50 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Necesitas estudiar más")
	}

	// 2. SWITCH
	fmt.Println("\n2. Switch:")

	// Generar día aleatorio
	rand.Seed(time.Now().UnixNano())
	dia := rand.Intn(7) + 1

	switch dia {
	case 1:
		fmt.Println("Lunes - ¡Nueva semana!")
	case 2, 3, 4:
		fmt.Println("Martes a Jueves - Días productivos")
	case 5:
		fmt.Println("Viernes - ¡Casi fin de semana!")
	case 6, 7:
		fmt.Println("Fin de semana - ¡A descansar!")
	default:
		fmt.Println("Día inválido")
	}

	// Switch sin expresión (como if-else)
	nota := 8.5
	switch {
	case nota >= 9:
		fmt.Println("Sobresaliente")
	case nota >= 7:
		fmt.Println("Notable")
	case nota >= 5:
		fmt.Println("Aprobado")
	default:
		fmt.Println("Suspenso")
	}

	// 3. BUCLES FOR
	fmt.Println("\n3. Bucles:")

	// For tradicional
	fmt.Println("Contando del 1 al 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For como while
	fmt.Println("Cuenta regresiva:")
	contador := 5
	for contador > 0 {
		fmt.Printf("%d... ", contador)
		contador--
	}
	fmt.Println("¡Despegue!")

	// For infinito con break
	fmt.Println("\nBuscando número par:")
	for {
		numero := rand.Intn(10)
		fmt.Printf("Número: %d ", numero)
		if numero%2 == 0 {
			fmt.Printf("- ¡Par encontrado!\n")
			break
		}
		fmt.Println("- Impar, continuamos...")
	}

	// For con range (arrays)
	frutas := []string{"manzana", "banana", "naranja"}
	fmt.Println("\nFrutas disponibles:")
	for indice, fruta := range frutas {
		fmt.Printf("%d. %s\n", indice+1, fruta)
	}

	// 4. CONTINUE
	fmt.Println("\nNúmeros impares del 1 al 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Salta los números pares
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 5. testando if con inicialización
	fmt.Println("\n5. If con inicialización:")

	if j := 43; j > 0 {
		fmt.Printf("j es positivo: %d\n", j)
	} else {
		fmt.Println("j no es positivo")
	}
}
