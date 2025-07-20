// API REST para Gestión de Libros
// Este proyecto demuestra cómo crear una API REST completa en Go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Estructura de datos para un libro
type Libro struct {
	ID          int       `json:"id"`
	Titulo      string    `json:"titulo"`
	Autor       string    `json:"autor"`
	Año         int       `json:"año"`
	Genero      string    `json:"genero"`
	Disponible  bool      `json:"disponible"`
	FechaCreado time.Time `json:"fecha_creado"`
}

// Base de datos en memoria
var libros []Libro
var contadorID = 1

// Middleware para logging
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Iniciado %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Completado %s %s en %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// Middleware para CORS
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Helper para respuestas JSON
func responderJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Helper para respuestas de error
func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, map[string]string{"error": mensaje})
}

// GET /api/libros - Obtener todos los libros
func obtenerLibros(w http.ResponseWriter, r *http.Request) {
	// Parámetros de consulta opcionales
	genero := r.URL.Query().Get("genero")
	disponible := r.URL.Query().Get("disponible")

	librosResultado := libros

	// Filtrar por género si se especifica
	if genero != "" {
		var filtrados []Libro
		for _, libro := range libros {
			if strings.EqualFold(libro.Genero, genero) {
				filtrados = append(filtrados, libro)
			}
		}
		librosResultado = filtrados
	}

	// Filtrar por disponibilidad si se especifica
	if disponible != "" {
		esDisponible, _ := strconv.ParseBool(disponible)
		var filtrados []Libro
		for _, libro := range librosResultado {
			if libro.Disponible == esDisponible {
				filtrados = append(filtrados, libro)
			}
		}
		librosResultado = filtrados
	}

	responderJSON(w, http.StatusOK, map[string]interface{}{
		"libros": librosResultado,
		"total":  len(librosResultado),
	})
}

// GET /api/libros/{id} - Obtener un libro específico
func obtenerLibroPorID(w http.ResponseWriter, r *http.Request) {
	// Extraer ID de la URL
	path := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(path)
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Buscar el libro
	for _, libro := range libros {
		if libro.ID == id {
			responderJSON(w, http.StatusOK, libro)
			return
		}
	}

	responderError(w, http.StatusNotFound, "Libro no encontrado")
}

// POST /api/libros - Crear un nuevo libro
func crearLibro(w http.ResponseWriter, r *http.Request) {
	var nuevoLibro Libro

	// Decodificar JSON del body
	if err := json.NewDecoder(r.Body).Decode(&nuevoLibro); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	// Validaciones
	if nuevoLibro.Titulo == "" {
		responderError(w, http.StatusBadRequest, "El título es requerido")
		return
	}
	if nuevoLibro.Autor == "" {
		responderError(w, http.StatusBadRequest, "El autor es requerido")
		return
	}
	if nuevoLibro.Año < 1000 || nuevoLibro.Año > time.Now().Year() {
		responderError(w, http.StatusBadRequest, "Año inválido")
		return
	}

	// Asignar ID y fecha
	nuevoLibro.ID = contadorID
	contadorID++
	nuevoLibro.FechaCreado = time.Now()
	nuevoLibro.Disponible = true // Por defecto disponible

	// Agregar a la base de datos
	libros = append(libros, nuevoLibro)

	responderJSON(w, http.StatusCreated, nuevoLibro)
}

// PUT /api/libros/{id} - Actualizar un libro
func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	// Extraer ID
	path := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(path)
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Buscar el libro
	indice := -1
	for i, libro := range libros {
		if libro.ID == id {
			indice = i
			break
		}
	}

	if indice == -1 {
		responderError(w, http.StatusNotFound, "Libro no encontrado")
		return
	}

	// Decodificar datos actualizados
	var libroActualizado Libro
	if err := json.NewDecoder(r.Body).Decode(&libroActualizado); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	// Mantener ID y fecha de creación original
	libroActualizado.ID = libros[indice].ID
	libroActualizado.FechaCreado = libros[indice].FechaCreado

	// Validaciones
	if libroActualizado.Titulo == "" {
		responderError(w, http.StatusBadRequest, "El título es requerido")
		return
	}
	if libroActualizado.Autor == "" {
		responderError(w, http.StatusBadRequest, "El autor es requerido")
		return
	}

	// Actualizar en la base de datos
	libros[indice] = libroActualizado

	responderJSON(w, http.StatusOK, libroActualizado)
}

// DELETE /api/libros/{id} - Eliminar un libro
func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	// Extraer ID
	path := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(path)
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Buscar y eliminar el libro
	for i, libro := range libros {
		if libro.ID == id {
			// Eliminar del slice
			libros = append(libros[:i], libros[i+1:]...)
			responderJSON(w, http.StatusOK, map[string]string{
				"mensaje": "Libro eliminado correctamente",
			})
			return
		}
	}

	responderError(w, http.StatusNotFound, "Libro no encontrado")
}

// Router principal
func manejarRuta(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/api/libros" && r.Method == "GET":
		obtenerLibros(w, r)
	case path == "/api/libros" && r.Method == "POST":
		crearLibro(w, r)
	case strings.HasPrefix(path, "/api/libros/") && r.Method == "GET":
		obtenerLibroPorID(w, r)
	case strings.HasPrefix(path, "/api/libros/") && r.Method == "PUT":
		actualizarLibro(w, r)
	case strings.HasPrefix(path, "/api/libros/") && r.Method == "DELETE":
		eliminarLibro(w, r)
	default:
		responderError(w, http.StatusNotFound, "Endpoint no encontrado")
	}
}

func inicializarDatos() {
	// Datos de ejemplo
	libros = []Libro{
		{
			ID:          1,
			Titulo:      "Cien años de soledad",
			Autor:       "Gabriel García Márquez",
			Año:         1967,
			Genero:      "Realismo mágico",
			Disponible:  true,
			FechaCreado: time.Now().AddDate(0, 0, -10),
		},
		{
			ID:          2,
			Titulo:      "1984",
			Autor:       "George Orwell",
			Año:         1949,
			Genero:      "Distopía",
			Disponible:  true,
			FechaCreado: time.Now().AddDate(0, 0, -8),
		},
		{
			ID:          3,
			Titulo:      "El Quijote",
			Autor:       "Miguel de Cervantes",
			Año:         1605,
			Genero:      "Clásico",
			Disponible:  false,
			FechaCreado: time.Now().AddDate(0, 0, -5),
		},
	}
	contadorID = 4
}

func main() {
	// Inicializar datos de ejemplo
	inicializarDatos()

	// Configurar rutas
	handler := corsMiddleware(loggingMiddleware(manejarRuta))
	http.HandleFunc("/", handler)

	// Información de inicio
	puerto := ":8080"
	fmt.Printf("🚀 Servidor API de Libros iniciado en http://localhost%s\n", puerto)
	fmt.Println("📚 Endpoints disponibles:")
	fmt.Println("  GET    /api/libros           - Obtener todos los libros")
	fmt.Println("  GET    /api/libros/{id}      - Obtener libro por ID")
	fmt.Println("  POST   /api/libros           - Crear nuevo libro")
	fmt.Println("  PUT    /api/libros/{id}      - Actualizar libro")
	fmt.Println("  DELETE /api/libros/{id}      - Eliminar libro")
	fmt.Println("\n💡 Ejemplos de uso con curl:")
	fmt.Println("  curl http://localhost:8080/api/libros")
	fmt.Println("  curl -X POST -H 'Content-Type: application/json' -d '{\"titulo\":\"Mi Libro\",\"autor\":\"Mi Autor\",\"año\":2023,\"genero\":\"Ficción\"}' http://localhost:8080/api/libros")

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(puerto, nil))
}
