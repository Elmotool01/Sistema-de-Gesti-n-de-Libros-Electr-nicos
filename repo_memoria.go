package main // Paquete principal.

// RepoMemoria implementa RepositorioLibros usando almacenamiento en memoria.
// Usamos:
// - map[int]Libro para búsqueda rápida por ID.
// - []int para mantener el orden de inserción al listar.
type RepoMemoria struct { // Estructura del repositorio en memoria.
	porID map[int]Libro // Mapa: clave = ID, valor = Libro.
	orden []int         // Slice con los IDs en el orden de inserción.
}

// NewRepoMemoria crea e inicializa un repositorio en memoria.
func NewRepoMemoria() *RepoMemoria { // Retorna un puntero a RepoMemoria listo para usar.
	return &RepoMemoria{ // Crea el objeto repositorio.
		porID: make(map[int]Libro), // Inicializa el map (evita nil map).
		orden: make([]int, 0),      // Inicializa el slice vacío.
	} // Fin de la creación del repositorio.
}

// Agregar inserta un libro si su ID no existe.
func (r *RepoMemoria) Agregar(libro Libro) error { // Método de la interfaz RepositorioLibros.
	id := libro.ID()              // Obtiene el id mediante getter (encapsulación).
	if _, ok := r.porID[id]; ok { // Verifica si ya existe un libro con ese id.
		return ErrDuplicado // Retorna error si es duplicado.
	}
	r.porID[id] = libro           // Guarda el libro en el map.
	r.orden = append(r.orden, id) // Guarda el id en el slice para mantener el orden.
	return nil                    // Retorna nil si todo salió bien.
}

// Actualizar reemplaza un libro existente con el mismo ID.
func (r *RepoMemoria) Actualizar(libro Libro) error { // Método de la interfaz.
	id := libro.ID()               // Obtiene el id con getter.
	if _, ok := r.porID[id]; !ok { // Si no existe el id, no se puede actualizar.
		return ErrNoExiste // Retorna error “no existe”.
	}
	r.porID[id] = libro // Reemplaza el libro en el map (actualización).
	return nil          // Retorna nil si todo bien.
}

// Eliminar borra un libro del map y del slice de orden.
func (r *RepoMemoria) Eliminar(id int) error { // Método de la interfaz.
	if _, ok := r.porID[id]; !ok { // Verifica existencia antes de borrar.
		return ErrNoExiste // Si no existe, retorna error.
	}

	delete(r.porID, id) // Elimina el libro del map.

	// --- Parte compleja: eliminar el ID del slice "orden" sin romper el arreglo ---
	// Recorremos el slice para encontrar el índice del ID a eliminar.
	for i, v := range r.orden { // i = índice, v = valor (id guardado).
		if v == id { // Si encontramos el id a eliminar...
			// Removemos el elemento i uniendo:
			// - r.orden[:i]  (todo antes de i)
			// - r.orden[i+1:] (todo después de i)
			// Esto evita dejar huecos y mantiene el orden restante.
			r.orden = append(r.orden[:i], r.orden[i+1:]...) // Elimina el elemento del slice.
			break                                           // Salimos porque el ID es único.
		}
	}
	// ---------------------------------------------------------------------------

	return nil // Retorna nil si la eliminación fue exitosa.
}

// BuscarPorID retorna un libro por su ID.
func (r *RepoMemoria) BuscarPorID(id int) (Libro, error) { // Método de la interfaz.
	libro, ok := r.porID[id] // Busca en el map.
	if !ok {                 // Si no se encontró...
		return Libro{}, ErrNoExiste // Retorna libro vacío + error.
	}
	return libro, nil // Retorna el libro encontrado + nil error.
}

// Listar devuelve libros respetando el orden de inserción.
func (r *RepoMemoria) Listar() []Libro { // Método de la interfaz.
	out := make([]Libro, 0, len(r.orden)) // Crea slice de salida con capacidad esperada.
	for _, id := range r.orden {          // Recorre IDs en orden.
		out = append(out, r.porID[id]) // Agrega el libro correspondiente al slice de salida.
	}
	return out // Devuelve la lista final.
}
