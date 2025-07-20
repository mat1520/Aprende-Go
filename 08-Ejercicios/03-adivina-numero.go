// Ejercicio 3: Adivina el Número
// Objetivo: Juego donde el usuario debe adivinar un número aleatorio
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== JUEGO: ADIVINA EL NÚMERO ===")

	// Configurar el juego
	rand.Seed(time.Now().UnixNano())
	numeroSecreto := rand.Intn(100) + 1
	maxIntentos := 7
	intentos := 0

	fmt.Printf("¡Bienvenido al juego!\n")
	fmt.Printf("He pensado un número entre 1 y 100.\n")
	fmt.Printf("Tienes %d intentos para adivinarlo.\n\n", maxIntentos)

	// Bucle principal del juego
	for intentos < maxIntentos {
		intentos++

		fmt.Printf("Intento %d/%d - Ingresa tu número: ", intentos, maxIntentos)

		var guess int
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("⚠️  Por favor, ingresa un número válido.")
			intentos-- // No contar como intento válido
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("⚠️  El número debe estar entre 1 y 100.")
			intentos-- // No contar como intento válido
			continue
		}

		// Verificar el resultado
		if guess == numeroSecreto {
			fmt.Printf("🎉 ¡FELICIDADES! Has adivinado el número %d en %d intentos!\n", numeroSecreto, intentos)
			mostrarCalificacion(intentos, maxIntentos)
			return
		} else if guess < numeroSecreto {
			diferencia := numeroSecreto - guess
			if diferencia <= 5 {
				fmt.Println("🔥 ¡Muy cerca! El número es un poco más alto.")
			} else if diferencia <= 15 {
				fmt.Println("📈 El número es más alto.")
			} else {
				fmt.Println("⬆️  El número es mucho más alto.")
			}
		} else {
			diferencia := guess - numeroSecreto
			if diferencia <= 5 {
				fmt.Println("🔥 ¡Muy cerca! El número es un poco más bajo.")
			} else if diferencia <= 15 {
				fmt.Println("📉 El número es más bajo.")
			} else {
				fmt.Println("⬇️  El número es mucho más bajo.")
			}
		}

		// Mostrar intentos restantes
		intentosRestantes := maxIntentos - intentos
		if intentosRestantes > 0 {
			fmt.Printf("Te quedan %d intentos.\n\n", intentosRestantes)
		}
	}

	// El jugador se quedó sin intentos
	fmt.Printf("😞 ¡Se acabaron los intentos! El número era %d.\n", numeroSecreto)
	fmt.Println("¡Mejor suerte la próxima vez!")

	// Preguntar si quiere jugar de nuevo
	preguntarJugarDeNuevo()
}

func mostrarCalificacion(intentos, maxIntentos int) {
	fmt.Println("\n=== TU CALIFICACIÓN ===")

	switch {
	case intentos == 1:
		fmt.Println("🏆 ¡INCREÍBLE! ¡Adivinaste en el primer intento!")
		fmt.Println("⭐⭐⭐⭐⭐ (5/5 estrellas)")
	case intentos <= 3:
		fmt.Println("🥇 ¡EXCELENTE! Muy pocas personas lo logran tan rápido.")
		fmt.Println("⭐⭐⭐⭐⭐ (5/5 estrellas)")
	case intentos <= 5:
		fmt.Println("🥈 ¡MUY BIEN! Eres bueno en este juego.")
		fmt.Println("⭐⭐⭐⭐ (4/5 estrellas)")
	case intentos <= 6:
		fmt.Println("🥉 ¡BIEN! Lo lograste con tiempo de sobra.")
		fmt.Println("⭐⭐⭐ (3/5 estrellas)")
	default:
		fmt.Println("✅ ¡COMPLETADO! Lo importante es que no te rendiste.")
		fmt.Println("⭐⭐ (2/5 estrellas)")
	}
}

func preguntarJugarDeNuevo() {
	fmt.Print("\n¿Quieres jugar de nuevo? (s/n): ")
	var respuesta string
	fmt.Scanln(&respuesta)

	if respuesta == "s" || respuesta == "S" || respuesta == "si" || respuesta == "SI" {
		fmt.Println("\n" + strings.Repeat("=", 40))
		main() // Reiniciar el juego
	} else {
		fmt.Println("¡Gracias por jugar! 👋")
	}
}

// VERSIONES ADICIONALES DEL JUEGO:

// Versión con niveles de dificultad
func juegoConNiveles() {
	fmt.Println("Elige el nivel de dificultad:")
	fmt.Println("1. Fácil (1-50, 10 intentos)")
	fmt.Println("2. Medio (1-100, 7 intentos)")
	fmt.Println("3. Difícil (1-200, 5 intentos)")

	var nivel int
	fmt.Print("Nivel (1-3): ")
	fmt.Scanln(&nivel)

	var max, intentos int
	switch nivel {
	case 1:
		max, intentos = 50, 10
	case 2:
		max, intentos = 100, 7
	case 3:
		max, intentos = 200, 5
	default:
		fmt.Println("Nivel inválido, usando nivel medio.")
		max, intentos = 100, 7
	}

	fmt.Printf("Juego configurado: 1-%d con %d intentos\n", max, intentos)
}

// Versión con pistas matemáticas
func darPistaMatematica(numero int) {
	fmt.Print("💡 Pista: El número ")

	if numero%2 == 0 {
		fmt.Print("es par")
	} else {
		fmt.Print("es impar")
	}

	if esPrimo(numero) {
		fmt.Print(" y es primo")
	}

	if numero%5 == 0 {
		fmt.Print(" y es múltiplo de 5")
	}

	fmt.Println()
}

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
