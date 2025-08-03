// Programa básico en Go: Diccionarios (Maps)
// Este ejemplo enseña cómo trabajar con mapas en Go
package main

import "fmt"

func main() {
	fmt.Println("=== DICCIONARIOS (MAPS) EN GO ===")

	// 1. 🗂️  CREAR Y DEFINIR UN MAPA
	fmt.Println("\n1. 🗂️  Creando nuestro diccionario:")

	// Definir un mapa con clave de tipo string y valores de tipo string
	lenguajes := make(map[string]string)

	// 2. ➕ AGREGAR PARES CLAVE-VALOR
	fmt.Println("\n2. ➕ Agregando elementos al diccionario:")

	// Agregar al menos tres pares clave-valor
	lenguajes["Go"] = "Lenguaje de programación desarrollado por Google"
	lenguajes["Python"] = "Lenguaje de programación interpretado y fácil de aprender"
	lenguajes["JavaScript"] = "Lenguaje de programación para desarrollo web"
	lenguajes["Java"] = "Lenguaje de programación orientado a objetos"

	fmt.Println("✅ Se agregaron 4 lenguajes de programación al diccionario")

	// 3. 📋 IMPRIMIR EL CONTENIDO DEL MAPA
	fmt.Println("\n3. 📋 Contenido completo del diccionario:")

	// Mostrar todo el mapa
	fmt.Printf("Diccionario completo: %v\n", lenguajes)

	// Mostrar cada elemento de forma más legible
	fmt.Println("\nDetalle de cada lenguaje:")
	for clave, valor := range lenguajes {
		fmt.Printf("🔹 %s: %s\n", clave, valor)
	}

	// 4. 🔍 BUSCAR UNA CLAVE ESPECÍFICA
	fmt.Println("\n4. 🔍 Buscando una clave específica:")

	// Buscar la clave "Go"
	claveABuscar := "Go"
	if descripcion, existe := lenguajes[claveABuscar]; existe {
		fmt.Printf("✅ Encontrado! %s: %s\n", claveABuscar, descripcion)
	} else {
		fmt.Printf("❌ No se encontró la clave '%s'\n", claveABuscar)
	}

	// Ejemplo adicional: buscar una clave que no existe
	claveInexistente := "Rust"
	if descripcion, existe := lenguajes[claveInexistente]; existe {
		fmt.Printf("✅ Encontrado! %s: %s\n", claveInexistente, descripcion)
	} else {
		fmt.Printf("❌ No se encontró la clave '%s' en nuestro diccionario\n", claveInexistente)
	}

	// 5. 📊 INFORMACIÓN ADICIONAL DEL MAPA
	fmt.Println("\n5. 📊 Información del diccionario:")
	fmt.Printf("Número total de elementos: %d\n", len(lenguajes))

	// 6. 🎯 ACCESO DIRECTO A VALORES
	fmt.Println("\n6. 🎯 Acceso directo a valores:")
	fmt.Printf("Descripción de Python: %s\n", lenguajes["Python"])
	fmt.Printf("Descripción de JavaScript: %s\n", lenguajes["JavaScript"])

	fmt.Println("\n🎉 ¡Programa completado! Has aprendido lo básico sobre mapas en Go.")
}
