// Arrays, Slices, Maps y Structs en Go
package main

import "fmt"

// Definición de struct (fuera de main)
type Persona struct {
	Nombre  string
	Edad    int
	Email   string
	Salario float64
	Activo  bool
}

// Método para struct
func (p Persona) presentarse() string {
	return fmt.Sprintf("Hola, soy %s, tengo %d años y trabajo en tecnología", p.Nombre, p.Edad)
}

// Método que modifica el struct (pointer receiver)
func (p *Persona) cumplirAnos() {
	p.Edad++
}

func main() {
	fmt.Println("=== ESTRUCTURAS DE DATOS EN GO ===")

	// 1. ARRAYS (tamaño fijo)
	fmt.Println("\n1. ARRAYS:")

	// Declaración de array
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Printf("Array: %v\n", numeros)
	fmt.Printf("Primer elemento: %d\n", numeros[0])
	fmt.Printf("Longitud del array: %d\n", len(numeros))

	// Array con valores iniciales
	dias := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	fmt.Printf("Días de la semana: %v\n", dias)

	// 2. SLICES (arrays dinámicos)
	fmt.Println("\n2. SLICES:")

	// Crear slice vacío
	var frutas []string
	fmt.Printf("Slice vacío: %v (longitud: %d)\n", frutas, len(frutas))

	// Añadir elementos con append
	frutas = append(frutas, "manzana", "banana", "naranja")
	fmt.Printf("Después de append: %v (longitud: %d)\n", frutas, len(frutas))

	// Slice con make
	numeros2 := make([]int, 3, 5) // longitud 3, capacidad 5
	numeros2[0] = 100
	numeros2[1] = 200
	numeros2[2] = 300
	fmt.Printf("Slice con make: %v (len: %d, cap: %d)\n", numeros2, len(numeros2), cap(numeros2))

	// Slicing (obtener sub-slice)
	subFrutas := frutas[1:3] // desde índice 1 hasta 2 (3 no incluido)
	fmt.Printf("Sub-slice frutas[1:3]: %v\n", subFrutas)

	// Iterar slice
	fmt.Println("Iterando frutas:")
	for i, fruta := range frutas {
		fmt.Printf("  %d: %s\n", i, fruta)
	}

	// 3. MAPS (diccionarios/hash tables)
	fmt.Println("\n3. MAPS:")

	// Crear map vacío
	edades := make(map[string]int)
	edades["Ana"] = 25
	edades["Carlos"] = 30
	edades["María"] = 28

	fmt.Printf("Map de edades: %v\n", edades)

	// Map con valores iniciales
	capitales := map[string]string{
		"España":   "Madrid",
		"Francia":  "París",
		"Italia":   "Roma",
		"Alemania": "Berlín",
	}

	// Acceder a valores
	fmt.Printf("Capital de España: %s\n", capitales["España"])

	// Verificar si existe una clave
	if capital, existe := capitales["Portugal"]; existe {
		fmt.Printf("Capital de Portugal: %s\n", capital)
	} else {
		fmt.Println("Portugal no está en nuestro map")
	}

	// Iterar map
	fmt.Println("Capitales europeas:")
	for pais, capital := range capitales {
		fmt.Printf("  %s: %s\n", pais, capital)
	}

	// Eliminar elemento
	delete(capitales, "Francia")
	fmt.Printf("Después de eliminar Francia: %v\n", capitales)

	// 4. STRUCTS (estructuras)
	fmt.Println("\n4. STRUCTS:")

	// Crear struct con valores
	empleado1 := Persona{
		Nombre:  "Juan Pérez",
		Edad:    32,
		Email:   "juan@empresa.com",
		Salario: 45000.0,
		Activo:  true,
	}

	fmt.Printf("Empleado 1: %+v\n", empleado1)
	fmt.Printf("Nombre: %s, Edad: %d\n", empleado1.Nombre, empleado1.Edad)

	// Usar método del struct
	fmt.Println(empleado1.presentarse())

	// Crear struct sin nombres de campo
	empleado2 := Persona{"Ana García", 28, "ana@empresa.com", 48000.0, true}
	fmt.Printf("Empleado 2: %+v\n", empleado2)

	// Modificar struct con método pointer
	fmt.Printf("Edad antes: %d\n", empleado1.Edad)
	empleado1.cumplirAnos()
	fmt.Printf("Edad después: %d\n", empleado1.Edad)

	// Struct anónimo
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Configuración: %+v\n", config)

	// 5. SLICE DE STRUCTS
	fmt.Println("\n5. SLICE DE STRUCTS:")

	empleados := []Persona{
		{"Luis", 35, "luis@empresa.com", 50000, true},
		{"Carmen", 29, "carmen@empresa.com", 46000, true},
		{"Roberto", 41, "roberto@empresa.com", 55000, false},
	}

	fmt.Println("Lista de empleados:")
	for i, emp := range empleados {
		fmt.Printf("  %d. %s (%d años) - Activo: %t\n", i+1, emp.Nombre, emp.Edad, emp.Activo)
	}

	// 6. MAP DE STRUCTS
	fmt.Println("\n6. MAP DE STRUCTS:")

	empleadosPorID := map[int]Persona{
		101: {"David", 33, "david@empresa.com", 47000, true},
		102: {"Elena", 27, "elena@empresa.com", 44000, true},
		103: {"Miguel", 39, "miguel@empresa.com", 52000, true},
	}

	if emp, existe := empleadosPorID[101]; existe {
		fmt.Printf("Empleado ID 101: %s\n", emp.Nombre)
	}
}
