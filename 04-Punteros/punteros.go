// Punteros en Go - Conceptos Fundamentales
package main

import "fmt"

// Struct para ejemplos
type Cuenta struct {
	Titular string
	Saldo   float64
}

// Método que NO modifica (value receiver)
func (c Cuenta) consultarSaldo() float64 {
	return c.Saldo
}

// Método que SÍ modifica (pointer receiver)
func (c *Cuenta) depositar(cantidad float64) {
	c.Saldo += cantidad
}

func (c *Cuenta) retirar(cantidad float64) bool {
	if c.Saldo >= cantidad {
		c.Saldo -= cantidad
		return true
	}
	return false
}

// Función que recibe valor (copia)
func incrementarPorValor(numero int) {
	numero = numero + 10
	fmt.Printf("  Dentro de la función: %d\n", numero)
}

// Función que recibe puntero (referencia)
func incrementarPorReferencia(numero *int) {
	*numero = *numero + 10
	fmt.Printf("  Dentro de la función: %d\n", *numero)
}

// Función que intercambia dos valores usando punteros
func intercambiar(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("=== PUNTEROS EN GO ===")

	// 1. CONCEPTOS BÁSICOS
	fmt.Println("\n1. Conceptos básicos:")

	numero := 42
	fmt.Printf("Valor de numero: %d\n", numero)
	fmt.Printf("Dirección de memoria de numero: %p\n", &numero)

	// Crear un puntero
	var punteroNumero *int = &numero
	fmt.Printf("Valor del puntero: %p\n", punteroNumero)
	fmt.Printf("Valor al que apunta el puntero: %d\n", *punteroNumero)

	// 2. MODIFICAR VALORES A TRAVÉS DE PUNTEROS
	fmt.Println("\n2. Modificar valores con punteros:")

	*punteroNumero = 100
	fmt.Printf("Después de modificar con puntero: %d\n", numero)

	// 3. DIFERENCIA: PASAR POR VALOR vs PASAR POR REFERENCIA
	fmt.Println("\n3. Pasar por valor vs referencia:")

	x := 5
	fmt.Printf("Valor inicial de x: %d\n", x)

	fmt.Println("Llamando incrementarPorValor:")
	incrementarPorValor(x)
	fmt.Printf("Valor de x después: %d (no cambió)\n", x)

	fmt.Println("Llamando incrementarPorReferencia:")
	incrementarPorReferencia(&x)
	fmt.Printf("Valor de x después: %d (sí cambió)\n", x)

	// 4. PUNTEROS CON STRUCTS
	fmt.Println("\n4. Punteros con structs:")

	cuenta := Cuenta{
		Titular: "María González",
		Saldo:   1000.0,
	}

	fmt.Printf("Cuenta inicial: %+v\n", cuenta)
	fmt.Printf("Saldo: %.2f\n", cuenta.consultarSaldo())

	// Depositar usando método con pointer receiver
	cuenta.depositar(500.0)
	fmt.Printf("Después de depositar 500: %.2f\n", cuenta.Saldo)

	// Retirar dinero
	if cuenta.retirar(300.0) {
		fmt.Printf("Retiro exitoso. Saldo actual: %.2f\n", cuenta.Saldo)
	}

	// 5. PUNTERO A STRUCT
	fmt.Println("\n5. Puntero a struct:")

	ptrCuenta := &cuenta
	fmt.Printf("Acceso directo - Titular: %s\n", ptrCuenta.Titular)
	fmt.Printf("Con (*ptrCuenta) - Saldo: %.2f\n", (*ptrCuenta).Saldo)

	// Modificar a través del puntero
	ptrCuenta.Titular = "María González Pérez"
	fmt.Printf("Titular modificado: %s\n", cuenta.Titular)

	// 6. NEW() - CREAR PUNTEROS
	fmt.Println("\n6. Función new():")

	// new() crea un puntero a un tipo y lo inicializa con zero value
	ptrEntero := new(int)
	fmt.Printf("Valor inicial con new: %d\n", *ptrEntero)

	*ptrEntero = 25
	fmt.Printf("Después de asignar: %d\n", *ptrEntero)

	// Con struct
	ptrCuentaNueva := new(Cuenta)
	ptrCuentaNueva.Titular = "Carlos Ruiz"
	ptrCuentaNueva.Saldo = 2000.0
	fmt.Printf("Nueva cuenta: %+v\n", *ptrCuentaNueva)

	// 7. PUNTEROS NIL
	fmt.Println("\n7. Punteros nil:")

	var ptrVacio *int
	fmt.Printf("Puntero vacío: %v\n", ptrVacio)

	if ptrVacio == nil {
		fmt.Println("El puntero es nil")
		ptrVacio = &numero
		fmt.Printf("Ahora apunta a: %d\n", *ptrVacio)
	}

	// 8. INTERCAMBIAR VALORES
	fmt.Println("\n8. Intercambiar valores:")

	a, b := 10, 20
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	intercambiar(&a, &b)
	fmt.Printf("Después: a=%d, b=%d\n", a, b)

	// 9. SLICE DE PUNTEROS
	fmt.Println("\n9. Slice de punteros:")

	numeros := []int{1, 2, 3, 4, 5}
	var punteros []*int

	for i := range numeros {
		punteros = append(punteros, &numeros[i])
	}

	fmt.Println("Modificando valores a través de punteros:")
	for i, ptr := range punteros {
		*ptr = *ptr * 10
		fmt.Printf("numeros[%d] = %d\n", i, numeros[i])
	}

	// 10. CUÁNDO USAR PUNTEROS
	fmt.Println("\n10. Cuándo usar punteros:")
	fmt.Println("✅ Cuando necesites modificar el valor original")
	fmt.Println("✅ Para evitar copiar structs grandes (mejor rendimiento)")
	fmt.Println("✅ Para implementar estructuras de datos como listas enlazadas")
	fmt.Println("❌ No los uses si solo necesitas leer el valor")
	fmt.Println("❌ Evítalos con tipos pequeños como int, bool")
}
