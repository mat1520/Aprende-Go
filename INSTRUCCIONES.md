# 📋 Guía de Ejecución de Ejemplos

## 🚀 Cómo ejecutar los ejemplos

### Prerrequisitos
- Go 1.19 o superior instalado
- Terminal/PowerShell configurado
- VS Code con extensión de Go (recomendado)

### Comandos básicos

#### Ejecutar un archivo Go:
```bash
go run nombre-archivo.go
```

#### Compilar y generar ejecutable:
```bash
go build nombre-archivo.go
```

## 📂 Orden recomendado de estudio

1. **01-Fundamentos/**
   - `go run 01-variables.go`
   - `go run 02-control-flujo.go`

2. **02-Funciones/**
   - `go run funciones.go`

3. **03-Estructuras/**
   - `go run estructuras.go`

4. **04-Punteros/**
   - `go run punteros.go`

5. **05-Errores/**
   - `go run errores.go`

6. **06-Interfaces/**
   - `go run interfaces.go`

7. **07-Concurrencia/**
   - `go run concurrencia.go`

8. **08-Ejercicios/**
   - `go run 01-calculadora.go 10 + 5`
   - `go run 02-contador-palabras.go`
   - `go run 03-adivina-numero.go`

9. **09-Proyectos/**
   - `cd api-libros && go run main.go`

## ⚡ Ejemplos de comandos rápidos

```powershell
# Ejecutar todos los ejemplos básicos
cd "c:\Users\Ariel\Desktop\Aprende-Go"
go run 01-Fundamentos\01-variables.go
go run 02-Funciones\funciones.go
go run 03-Estructuras\estructuras.go

# Probar la API de libros
cd 09-Proyectos\api-libros
go run main.go
```

¡Ahora ya tienes una estructura completa para aprender Go! 🎉
