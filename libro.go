package main // Paquete principal.

import "errors" // Se usa para devolver errores de validación.

// Libro representa el objeto principal del sistema (un libro electrónico).
// NOTA: Los campos son privados (minúscula) para aplicar encapsulación.
type Libro struct { // Definición de la “clase” en Go (struct).
	id     int    // Campo privado: identificador único del libro.
	titulo string // Campo privado: título del libro.
	autor  string // Campo privado: autor del libro.
	anio   int    // Campo privado: año de publicación.
}

// NewLibro crea un Libro nuevo aplicando validaciones (constructor).
func NewLibro(id int, titulo string, autor string, anio int) (*Libro, error) { // Función constructora con retorno (objeto, error).
	if id <= 0 { // Validación: el ID debe ser mayor que cero.
		return nil, errors.New("id debe ser mayor que 0") // Si falla, se devuelve nil y el error.
	}
	if titulo == "" { // Validación: título no puede ser vacío.
		return nil, errors.New("titulo no puede estar vacio") // Retorna error si es vacío.
	}
	if autor == "" { // Validación: autor no puede ser vacío.
		return nil, errors.New("autor no puede estar vacio") // Retorna error si es vacío.
	}
	if anio < 0 { // Validación: el año no puede ser negativo.
		return nil, errors.New("anio no puede ser negativo") // Retorna error si es negativo.
	}

	l := &Libro{ // Se crea el objeto Libro en memoria (puntero) para poder modificarlo con métodos.
		id:     id,     // Asigna el id validado.
		titulo: titulo, // Asigna el título validado.
		autor:  autor,  // Asigna el autor validado.
		anio:   anio,   // Asigna el año validado.
	}

	return l, nil // Devuelve el puntero al libro creado y nil (sin error).
}

// ID devuelve el identificador (getter) sin exponer el campo directamente.
func (l Libro) ID() int { // Método público que retorna el id.
	return l.id // Retorna el campo privado id.
}

// Titulo devuelve el título (getter).
func (l Libro) Titulo() string { // Método público que retorna el título.
	return l.titulo // Retorna el campo privado titulo.
}

// Autor devuelve el autor (getter).
func (l Libro) Autor() string { // Método público que retorna el autor.
	return l.autor // Retorna el campo privado autor.
}

// Anio devuelve el año (getter).
func (l Libro) Anio() int { // Método público que retorna el año.
	return l.anio // Retorna el campo privado anio.
}

// SetTitulo modifica el título con validación (setter).
func (l *Libro) SetTitulo(nuevo string) error { // Método con puntero porque modifica el objeto.
	if nuevo == "" { // Valida que el nuevo título no sea vacío.
		return errors.New("titulo no puede estar vacio") // Retorna error si el nuevo valor es inválido.
	}
	l.titulo = nuevo // Asigna el nuevo título al campo privado.
	return nil       // Retorna nil indicando que no hubo error.
}

// SetAutor modifica el autor con validación (setter).
func (l *Libro) SetAutor(nuevo string) error { // Método con puntero para modificar.
	if nuevo == "" { // Valida autor no vacío.
		return errors.New("autor no puede estar vacio") // Retorna error si es inválido.
	}
	l.autor = nuevo // Asigna el nuevo autor.
	return nil      // Retorna nil (ok).
}

// SetAnio modifica el año con validación (setter).
func (l *Libro) SetAnio(nuevo int) error { // Método con puntero para modificar.
	if nuevo < 0 { // Valida que el año no sea negativo.
		return errors.New("anio no puede ser negativo") // Retorna error si es inválido.
	}
	l.anio = nuevo // Asigna el nuevo año.
	return nil     // Retorna nil (ok).
}
