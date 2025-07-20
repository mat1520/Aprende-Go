// Manejo de Errores en Go - Guía Completa
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Error personalizado usando struct
type ErrorDivision struct {
	Dividendo float64
	Divisor   float64
	Mensaje   string
}

// Implementar interface error
func (e ErrorDivision) Error() string {
	return fmt.Sprintf("error división: %s (%.2f / %.2f)", e.Mensaje, e.Dividendo, e.Divisor)
}

// Función que puede devolver error simple
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// Función con error personalizado
func dividirConErrorPersonalizado(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorDivision{
			Dividendo: a,
			Divisor:   b,
			Mensaje:   "divisor no puede ser cero",
		}
	}
	return a / b, nil
}

// Función que convierte string a entero con manejo de errores
func convertirAEntero(str string) (int, error) {
	numero, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir '%s' a entero: %v", str, err)
	}
	return numero, nil
}

// Función que puede tener múltiples tipos de errores
func procesarArchivo(nombreArchivo string) error {
	if nombreArchivo == "" {
		return errors.New("nombre de archivo vacío")
	}

	if nombreArchivo == "archivo_protegido.txt" {
		return errors.New("acceso denegado al archivo")
	}

	if nombreArchivo == "archivo_inexistente.txt" {
		return errors.New("archivo no encontrado")
	}

	// Simulamos que todo salió bien
	return nil
}

// Función con defer para limpieza en caso de error
func operacionConLimpieza() error {
	fmt.Println("Iniciando operación...")

	defer func() {
		fmt.Println("Limpieza ejecutada (defer)")
	}()

	// Simulamos un error
	if true {
		return errors.New("algo salió mal en la operación")
	}

	fmt.Println("Operación completada exitosamente")
	return nil
}

// Función que maneja panic y recover
func operacionPeligrosa() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperado de panic: %v\n", r)
		}
	}()

	fmt.Println("Ejecutando operación peligrosa...")
	panic("¡Algo salió terriblemente mal!")
	fmt.Println("Esta línea nunca se ejecutará")
}

func main() {
	fmt.Println("=== MANEJO DE ERRORES EN GO ===")

	// 1. ERROR BÁSICO
	fmt.Println("\n1. Error básico:")

	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// División exitosa
	resultado, err = dividir(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// 2. ERROR PERSONALIZADO
	fmt.Println("\n2. Error personalizado:")

	_, err = dividirConErrorPersonalizado(15, 0)
	if err != nil {
		fmt.Printf("Error personalizado: %v\n", err)

		// Verificar tipo específico de error
		if errorDiv, ok := err.(ErrorDivision); ok {
			fmt.Printf("Detalles: Dividendo=%.2f, Divisor=%.2f\n",
				errorDiv.Dividendo, errorDiv.Divisor)
		}
	}

	// 3. WRAPPING DE ERRORES CON fmt.Errorf
	fmt.Println("\n3. Wrapping de errores:")

	numero, err := convertirAEntero("abc")
	if err != nil {
		fmt.Printf("Error wrapped: %v\n", err)
	}

	numero, err = convertirAEntero("42")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Número convertido: %d\n", numero)
	}

	// 4. MÚLTIPLES ERRORES POSIBLES
	fmt.Println("\n4. Múltiples errores posibles:")

	archivos := []string{
		"documento.txt",
		"",
		"archivo_protegido.txt",
		"archivo_inexistente.txt",
	}

	for _, archivo := range archivos {
		err := procesarArchivo(archivo)
		if err != nil {
			fmt.Printf("Error procesando '%s': %v\n", archivo, err)
		} else {
			fmt.Printf("Archivo '%s' procesado exitosamente\n", archivo)
		}
	}

	// 5. DEFER PARA LIMPIEZA
	fmt.Println("\n5. Defer para limpieza:")

	err = operacionConLimpieza()
	if err != nil {
		fmt.Printf("La operación falló: %v\n", err)
	}

	// 6. PANIC Y RECOVER
	fmt.Println("\n6. Panic y recover:")

	operacionPeligrosa()
	fmt.Println("El programa continúa después del panic")

	// 7. PATRONES COMUNES DE MANEJO DE ERRORES
	fmt.Println("\n7. Patrones comunes:")

	// Patrón 1: Error como parte del flujo normal
	if resultado, err := dividir(20, 4); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("División exitosa: %.2f\n", resultado)
	}

	// Patrón 2: Early return
	fmt.Println("\nFunción con early return:")
	miFuncionConEarlyReturn()

	// 8. BUENAS PRÁCTICAS
	fmt.Println("\n8. Buenas prácticas:")
	fmt.Println("✅ Siempre verifica los errores")
	fmt.Println("✅ Usa errores específicos y descriptivos")
	fmt.Println("✅ No ignores errores con _")
	fmt.Println("✅ Usa defer para limpieza")
	fmt.Println("✅ Panic solo para errores irrecuperables")
	fmt.Println("❌ No uses panic como control de flujo normal")
	fmt.Println("❌ No captures todos los panics sin razón")
}

func miFuncionConEarlyReturn() {
	fmt.Println("  Iniciando función...")

	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Printf("  Error encontrado, saliendo: %v\n", err)
		return // Early return - salir inmediatamente
	}

	// Esta línea solo se ejecuta si no hay error
	fmt.Printf("  Resultado calculado: %.2f\n", resultado)
	fmt.Println("  Función completada exitosamente")
}
