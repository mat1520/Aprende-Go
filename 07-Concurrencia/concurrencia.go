// Concurrencia en Go - Goroutines, Channels y Sincronización
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. GOROUTINE BÁSICA
func tarea(nombre string, duracion time.Duration) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("[%s] Paso %d\n", nombre, i)
		time.Sleep(duracion)
	}
	fmt.Printf("[%s] ¡Completada!\n", nombre)
}

// 2. FUNCIÓN PARA USAR CON CHANNELS
func productor(ch chan int, nombre string) {
	for i := 1; i <= 5; i++ {
		valor := rand.Intn(100)
		fmt.Printf("[%s] Enviando: %d\n", nombre, valor)
		ch <- valor
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumidor(ch chan int, nombre string) {
	for valor := range ch {
		fmt.Printf("[%s] Recibido: %d\n", nombre, valor)
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Printf("[%s] Canal cerrado\n", nombre)
}

// 3. WORKER POOL PATTERN
func worker(id int, trabajos <-chan int, resultados chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for trabajo := range trabajos {
		fmt.Printf("Worker %d procesando trabajo %d\n", id, trabajo)
		time.Sleep(time.Second)
		resultados <- trabajo * 2
	}
}

// 4. FUNCIÓN PARA RACE CONDITION
var contador int
var mutex sync.Mutex

func incrementarSeguro(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		contador++
		mutex.Unlock()
	}
}

func incrementarInseguro(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		contador++
	}
}

// 5. SELECT STATEMENT
func servidor(nombre string, ch chan string) {
	for i := 1; i <= 3; i++ {
		mensaje := fmt.Sprintf("%s-mensaje-%d", nombre, i)
		ch <- mensaje
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	close(ch)
}

func main() {
	fmt.Println("=== CONCURRENCIA EN GO ===")
	rand.Seed(time.Now().UnixNano())

	// 1. GOROUTINES BÁSICAS
	fmt.Println("\n1. Goroutines básicas:")

	fmt.Println("Ejecutando tareas de forma secuencial:")
	start := time.Now()
	tarea("Secuencial-A", 200*time.Millisecond)
	tarea("Secuencial-B", 200*time.Millisecond)
	fmt.Printf("Tiempo secuencial: %v\n", time.Since(start))

	fmt.Println("\nEjecutando tareas de forma concurrente:")
	start = time.Now()
	go tarea("Goroutine-A", 200*time.Millisecond)
	go tarea("Goroutine-B", 200*time.Millisecond)

	// Esperar a que terminen las goroutines
	time.Sleep(2 * time.Second)
	fmt.Printf("Tiempo concurrente: %v\n", time.Since(start))

	// 2. CHANNELS BÁSICOS
	fmt.Println("\n2. Channels básicos:")

	// Channel sin buffer
	ch := make(chan string)

	// Enviar y recibir en goroutines separadas
	go func() {
		ch <- "Hola desde goroutine"
	}()

	mensaje := <-ch
	fmt.Printf("Recibido: %s\n", mensaje)

	// Channel con buffer
	chBuffer := make(chan int, 3)
	chBuffer <- 1
	chBuffer <- 2
	chBuffer <- 3

	fmt.Printf("Channel buffered: %d, %d, %d\n", <-chBuffer, <-chBuffer, <-chBuffer)

	// 3. PRODUCTOR-CONSUMIDOR
	fmt.Println("\n3. Patrón Productor-Consumidor:")

	chProduccion := make(chan int)

	go productor(chProduccion, "Productor-1")
	go consumidor(chProduccion, "Consumidor-1")

	time.Sleep(4 * time.Second)

	// 4. WAITGROUP PARA SINCRONIZACIÓN
	fmt.Println("\n4. Sincronización con WaitGroup:")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d iniciada\n", id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Goroutine %d terminada\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Todas las goroutines han terminado")

	// 5. WORKER POOL
	fmt.Println("\n5. Worker Pool Pattern:")

	const numWorkers = 3
	const numTrabajos = 9

	trabajos := make(chan int, numTrabajos)
	resultados := make(chan int, numTrabajos)

	var wgWorkers sync.WaitGroup

	// Iniciar workers
	for i := 1; i <= numWorkers; i++ {
		wgWorkers.Add(1)
		go worker(i, trabajos, resultados, &wgWorkers)
	}

	// Enviar trabajos
	for i := 1; i <= numTrabajos; i++ {
		trabajos <- i
	}
	close(trabajos)

	// Goroutine para cerrar resultados cuando terminen todos los workers
	go func() {
		wgWorkers.Wait()
		close(resultados)
	}()

	// Recoger resultados
	fmt.Println("Resultados del Worker Pool:")
	for resultado := range resultados {
		fmt.Printf("Resultado: %d\n", resultado)
	}

	// 6. RACE CONDITIONS Y MUTEX
	fmt.Println("\n6. Race Conditions y Mutex:")

	// Ejemplo inseguro (race condition)
	contador = 0
	var wgRace sync.WaitGroup

	fmt.Println("Incremento inseguro (race condition):")
	for i := 0; i < 3; i++ {
		wgRace.Add(1)
		go incrementarInseguro(&wgRace)
	}
	wgRace.Wait()
	fmt.Printf("Contador final (esperado: 3000): %d\n", contador)

	// Ejemplo seguro con mutex
	contador = 0
	fmt.Println("Incremento seguro (con mutex):")
	for i := 0; i < 3; i++ {
		wgRace.Add(1)
		go incrementarSeguro(&wgRace)
	}
	wgRace.Wait()
	fmt.Printf("Contador final (esperado: 3000): %d\n", contador)

	// 7. SELECT STATEMENT
	fmt.Println("\n7. Select Statement:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go servidor("Servidor-A", ch1)
	go servidor("Servidor-B", ch2)

	// Recibir mensajes usando select
	canalesActivos := 2
	for canalesActivos > 0 {
		select {
		case msg, ok := <-ch1:
			if !ok {
				fmt.Println("Canal ch1 cerrado")
				ch1 = nil
				canalesActivos--
			} else {
				fmt.Printf("De ch1: %s\n", msg)
			}
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("Canal ch2 cerrado")
				ch2 = nil
				canalesActivos--
			} else {
				fmt.Printf("De ch2: %s\n", msg)
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout - no se recibió nada")
			canalesActivos = 0
		}
	}

	// 8. CHANNEL DIRECCIONALES
	fmt.Println("\n8. Channels direccionales:")

	// Canal solo para envío
	chSoloEnvio := make(chan<- int)
	// Canal solo para recepción
	chSoloRecepcion := make(<-chan int)

	fmt.Printf("Tipo channel solo envío: %T\n", chSoloEnvio)
	fmt.Printf("Tipo channel solo recepción: %T\n", chSoloRecepcion)

	// 9. BUENAS PRÁCTICAS
	fmt.Println("\n9. Buenas prácticas de concurrencia:")
	fmt.Println("✅ Usa goroutines para tareas independientes")
	fmt.Println("✅ Usa channels para comunicar entre goroutines")
	fmt.Println("✅ Usa WaitGroup para esperar goroutines")
	fmt.Println("✅ Usa mutex para proteger datos compartidos")
	fmt.Println("✅ Cierra channels cuando no los necesites más")
	fmt.Println("✅ Usa select para operaciones no bloqueantes")
	fmt.Println("❌ No compartas memoria, comunica usando channels")
	fmt.Println("❌ Evita race conditions")
	fmt.Println("❌ No olvides cerrar channels para evitar goroutine leaks")

	fmt.Println("\n¡Concurrencia completada!")
}
