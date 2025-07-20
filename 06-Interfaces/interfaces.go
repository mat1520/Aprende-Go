// Interfaces en Go - Conceptos y Ejemplos Prácticos
package main

import (
	"fmt"
	"math"
)

// 1. INTERFACE BÁSICA
type Animal interface {
	Sonido() string
	Moverse() string
}

// Implementaciones de la interface Animal
type Perro struct {
	Nombre string
	Raza   string
}

func (p Perro) Sonido() string {
	return "Guau guau"
}

func (p Perro) Moverse() string {
	return "Corre con cuatro patas"
}

type Gato struct {
	Nombre string
	Color  string
}

func (g Gato) Sonido() string {
	return "Miau"
}

func (g Gato) Moverse() string {
	return "Camina sigilosamente"
}

type Pájaro struct {
	Nombre   string
	Especies string
}

func (p Pájaro) Sonido() string {
	return "Pío pío"
}

func (p Pájaro) Moverse() string {
	return "Vuela por el aire"
}

// 2. INTERFACE PARA GEOMETRÍA
type Forma interface {
	Area() float64
	Perimetro() float64
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// 3. INTERFACE VACÍA (interface{})
func imprimirTipo(valor interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", valor, valor)
}

// 4. TYPE ASSERTION Y TYPE SWITCH
func procesarValor(valor interface{}) {
	switch v := valor.(type) {
	case int:
		fmt.Printf("Es un entero: %d\n", v)
	case string:
		fmt.Printf("Es un string: %s\n", v)
	case float64:
		fmt.Printf("Es un float64: %.2f\n", v)
	case Animal:
		fmt.Printf("Es un animal que hace: %s\n", v.Sonido())
	default:
		fmt.Printf("Tipo desconocido: %T\n", v)
	}
}

// 5. INTERFACE COMPOSITION (EMBEDDING)
type Volador interface {
	Volar() string
}

type Nadador interface {
	Nadar() string
}

// Interface que combina otras interfaces
type AnimalAcuático interface {
	Animal
	Nadador
}

type Pato struct {
	Nombre string
}

func (p Pato) Sonido() string {
	return "Cuac cuac"
}

func (p Pato) Moverse() string {
	return "Camina y nada"
}

func (p Pato) Nadar() string {
	return "Nada en el agua"
}

func (p Pato) Volar() string {
	return "Vuela bajo sobre el agua"
}

// 6. INTERFACE PARA COMPARACIÓN
type Comparable interface {
	EsIgual(otro Comparable) bool
}

type Punto struct {
	X, Y int
}

func (p Punto) EsIgual(otro Comparable) bool {
	if otroPunto, ok := otro.(Punto); ok {
		return p.X == otroPunto.X && p.Y == otroPunto.Y
	}
	return false
}

// 7. FUNCIÓN QUE ACEPTA INTERFACE
func hacerSonidoAnimal(a Animal) {
	fmt.Printf("El animal hace: %s y se mueve: %s\n", a.Sonido(), a.Moverse())
}

func calcularAreaTotal(formas []Forma) float64 {
	total := 0.0
	for _, forma := range formas {
		total += forma.Area()
	}
	return total
}

func main() {
	fmt.Println("=== INTERFACES EN GO ===")

	// 1. USO BÁSICO DE INTERFACES
	fmt.Println("\n1. Uso básico de interfaces:")

	var animal Animal

	animal = Perro{Nombre: "Rex", Raza: "Pastor Alemán"}
	hacerSonidoAnimal(animal)

	animal = Gato{Nombre: "Whiskers", Color: "Naranja"}
	hacerSonidoAnimal(animal)

	animal = Pájaro{Nombre: "Tweety", Especies: "Canario"}
	hacerSonidoAnimal(animal)

	// 2. SLICE DE INTERFACES
	fmt.Println("\n2. Slice de interfaces:")

	animales := []Animal{
		Perro{Nombre: "Bobby", Raza: "Golden"},
		Gato{Nombre: "Mimi", Color: "Blanco"},
		Pájaro{Nombre: "Piolín", Especies: "Canario"},
	}

	fmt.Println("Todos los animales:")
	for i, a := range animales {
		fmt.Printf("  %d. %s - %s\n", i+1, a.Sonido(), a.Moverse())
	}

	// 3. INTERFACES CON FORMAS GEOMÉTRICAS
	fmt.Println("\n3. Interfaces con geometría:")

	rectangulo := Rectangulo{Ancho: 5, Alto: 3}
	circulo := Circulo{Radio: 4}

	fmt.Printf("Rectángulo - Área: %.2f, Perímetro: %.2f\n",
		rectangulo.Area(), rectangulo.Perimetro())
	fmt.Printf("Círculo - Área: %.2f, Perímetro: %.2f\n",
		circulo.Area(), circulo.Perimetro())

	// Calcular área total usando interface
	formas := []Forma{rectangulo, circulo}
	areaTotal := calcularAreaTotal(formas)
	fmt.Printf("Área total de todas las formas: %.2f\n", areaTotal)

	// 4. INTERFACE VACÍA
	fmt.Println("\n4. Interface vacía (interface{}):")

	imprimirTipo(42)
	imprimirTipo("Hola Go")
	imprimirTipo(3.14)
	imprimirTipo([]int{1, 2, 3})
	imprimirTipo(Perro{Nombre: "Max"})

	// 5. TYPE SWITCH
	fmt.Println("\n5. Type switch:")

	valores := []interface{}{
		42,
		"Hola",
		3.14159,
		Gato{Nombre: "Luna"},
		true,
	}

	for _, valor := range valores {
		procesarValor(valor)
	}

	// 6. TYPE ASSERTION
	fmt.Println("\n6. Type assertion:")

	var animal2 interface{} = Perro{Nombre: "Buddy", Raza: "Labrador"}

	// Type assertion segura
	if perro, ok := animal2.(Perro); ok {
		fmt.Printf("Es un perro: %s de raza %s\n", perro.Nombre, perro.Raza)
	} else {
		fmt.Println("No es un perro")
	}

	// Type assertion a interface
	if a, ok := animal2.(Animal); ok {
		fmt.Printf("Implementa Animal: %s\n", a.Sonido())
	}

	// 7. INTERFACE COMPOSITION
	fmt.Println("\n7. Interface composition:")

	var pato Animal = Pato{Nombre: "Donald"}
	fmt.Printf("Pato como Animal: %s\n", pato.Sonido())

	// El pato también implementa AnimalAcuático
	var patoAcuatico AnimalAcuático = Pato{Nombre: "Daffy"}
	fmt.Printf("Pato acuático: %s y %s\n",
		patoAcuatico.Sonido(), patoAcuatico.Nadar())

	// 8. INTERFACE COMPARABLE
	fmt.Println("\n8. Interface comparable:")

	punto1 := Punto{X: 5, Y: 3}
	punto2 := Punto{X: 5, Y: 3}
	punto3 := Punto{X: 2, Y: 8}

	fmt.Printf("Punto1 == Punto2: %t\n", punto1.EsIgual(punto2))
	fmt.Printf("Punto1 == Punto3: %t\n", punto1.EsIgual(punto3))

	// 9. INTERFACE NIL
	fmt.Println("\n9. Interface nil:")

	var animalNil Animal
	fmt.Printf("Animal nil: %v\n", animalNil)

	if animalNil == nil {
		fmt.Println("La interface es nil")
	}

	// 10. BUENAS PRÁCTICAS
	fmt.Println("\n10. Buenas prácticas con interfaces:")
	fmt.Println("✅ Define interfaces pequeñas y específicas")
	fmt.Println("✅ Usa la interface vacía para tipos desconocidos")
	fmt.Println("✅ Implementa interfaces implícitamente")
	fmt.Println("✅ Usa type assertion de forma segura")
	fmt.Println("✅ Prefiere composición de interfaces")
	fmt.Println("❌ No definas interfaces muy grandes")
	fmt.Println("❌ No uses type assertion sin verificar")
}
