package main // Paquete principal.

import ( // Bloque de imports.
	"bufio"   // Para leer entrada del usuario por consola.
	"fmt"     // Para imprimir mensajes al usuario.
	"strconv" // Para convertir string a entero.
	"strings" // Para limpiar espacios/saltos de línea.
)

// leerLinea lee una línea y devuelve el texto limpio (sin espacios extras).
func leerLinea(r *bufio.Reader) string { // Recibe un lector y retorna string.
	s, _ := r.ReadString('\n')  // Lee hasta Enter (incluye '\n'); ignoramos error por simplicidad en consola.
	return strings.TrimSpace(s) // Quita espacios y saltos de línea.
}

// leerEntero solicita un entero y repite hasta que el usuario ingrese uno válido.
func leerEntero(r *bufio.Reader, mensaje string) int { // Recibe lector y mensaje, retorna int válido.
	for { // Bucle infinito hasta recibir un entero válido.
		fmt.Print(mensaje)        // Muestra el mensaje (ej: "ID: ").
		s := leerLinea(r)         // Lee la entrada del usuario.
		n, err := strconv.Atoi(s) // Intenta convertir a entero.
		if err == nil {           // Si no hubo error de conversión...
			return n // Devuelve el número válido.
		}
		fmt.Println("❌ Ingresa un número válido.") // Si falla, se informa y se repite el ciclo.
	}
}
