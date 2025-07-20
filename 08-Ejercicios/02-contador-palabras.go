// Ejercicio 2: Contador de Palabras
// Objetivo: Contar la frecuencia de palabras en un texto
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== CONTADOR DE PALABRAS ===")

	// Texto de ejemplo
	texto := `
	Go es un lenguaje de programación desarrollado por Google.
	Go es simple y eficiente.
	Los programadores aman Go por su simplicidad.
	Go tiene una excelente concurrencia.
	`

	fmt.Println("Texto a analizar:")
	fmt.Println(texto)

	// Analizar el texto
	estadisticas := analizarTexto(texto)

	// Mostrar resultados
	mostrarEstadisticas(estadisticas, texto)
}

func analizarTexto(texto string) map[string]int {
	// Convertir a minúsculas y dividir en palabras
	texto = strings.ToLower(texto)
	palabras := strings.Fields(texto)

	// Contador de palabras
	contador := make(map[string]int)

	for _, palabra := range palabras {
		// Limpiar signos de puntuación
		palabra = limpiarPalabra(palabra)
		if palabra != "" {
			contador[palabra]++
		}
	}

	return contador
}

func limpiarPalabra(palabra string) string {
	// Remover signos de puntuación comunes
	signos := ".,!?;:()\""

	for _, signo := range signos {
		palabra = strings.ReplaceAll(palabra, string(signo), "")
	}

	return strings.TrimSpace(palabra)
}

func mostrarEstadisticas(contador map[string]int, textoOriginal string) {
	fmt.Println("\n=== ESTADÍSTICAS ===")

	// Estadísticas generales
	totalPalabras := 0
	for _, count := range contador {
		totalPalabras += count
	}

	fmt.Printf("Total de palabras únicas: %d\n", len(contador))
	fmt.Printf("Total de palabras: %d\n", totalPalabras)
	fmt.Printf("Caracteres totales: %d\n", len(textoOriginal))

	// Crear slice para ordenar las palabras por frecuencia
	type PalabraConteo struct {
		palabra string
		conteo  int
	}

	var palabrasOrdenadas []PalabraConteo
	for palabra, conteo := range contador {
		palabrasOrdenadas = append(palabrasOrdenadas, PalabraConteo{palabra, conteo})
	}

	// Ordenar por conteo descendente
	sort.Slice(palabrasOrdenadas, func(i, j int) bool {
		return palabrasOrdenadas[i].conteo > palabrasOrdenadas[j].conteo
	})

	// Mostrar top 10 palabras más frecuentes
	fmt.Println("\n=== TOP 10 PALABRAS MÁS FRECUENTES ===")
	limite := 10
	if len(palabrasOrdenadas) < limite {
		limite = len(palabrasOrdenadas)
	}

	for i := 0; i < limite; i++ {
		pc := palabrasOrdenadas[i]
		porcentaje := float64(pc.conteo) / float64(totalPalabras) * 100
		fmt.Printf("%2d. %-15s: %2d veces (%.1f%%)\n",
			i+1, pc.palabra, pc.conteo, porcentaje)
	}

	// Mostrar todas las palabras alfabéticamente
	fmt.Println("\n=== TODAS LAS PALABRAS (A-Z) ===")

	// Crear slice para orden alfabético
	var palabrasAlfabeticas []string
	for palabra := range contador {
		palabrasAlfabeticas = append(palabrasAlfabeticas, palabra)
	}
	sort.Strings(palabrasAlfabeticas)

	for _, palabra := range palabrasAlfabeticas {
		fmt.Printf("%-15s: %d\n", palabra, contador[palabra])
	}
}

// EJERCICIO EXTRA: Versión que lee desde archivo
/*
func analizarArchivo(nombreArchivo string) {
	import "io/ioutil"

	data, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
		return
	}

	texto := string(data)
	estadisticas := analizarTexto(texto)
	mostrarEstadisticas(estadisticas, texto)
}
*/

// EJERCICIO EXTRA: Versión interactiva
/*
func versionInteractiva() {
	import "bufio"
	import "os"

	fmt.Println("Ingresa tu texto (termina con una línea vacía):")

	scanner := bufio.NewScanner(os.Stdin)
	var lineas []string

	for scanner.Scan() {
		linea := scanner.Text()
		if linea == "" {
			break
		}
		lineas = append(lineas, linea)
	}

	texto := strings.Join(lineas, " ")
	estadisticas := analizarTexto(texto)
	mostrarEstadisticas(estadisticas, texto)
}
*/
