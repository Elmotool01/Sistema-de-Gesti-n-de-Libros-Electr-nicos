package handlers // Paquete handlers: contiene la lógica de las rutas/controladores.

import (
	"database/sql"
	"html/template"
	"net/http"
	"sistema/models"
	"strconv"
	"strings"
)

// Stats representa estadísticas básicas del sistema para el dashboard.
type Stats struct {
	TotalLibros int
	TotalPDF    int
	TotalEPUB   int
	TotalMOBI   int
}

// LibroHandler agrupa recursos que usan los handlers de libros.
type LibroHandler struct {
	DB        *sql.DB
	Templates *template.Template
}

// NuevoLibroHandler crea una nueva instancia de LibroHandler.
func NuevoLibroHandler(db *sql.DB, templates *template.Template) *LibroHandler {
	return &LibroHandler{
		DB:        db,
		Templates: templates,
	}
}

// Index muestra el panel principal y envía permisos según rol.
func (h *LibroHandler) Index(w http.ResponseWriter, r *http.Request) {
	busqueda := strings.TrimSpace(r.URL.Query().Get("buscar"))
	mensaje := strings.TrimSpace(r.URL.Query().Get("msg"))

	nombreUsuario := ObtenerNombreUsuario(r)
	rolUsuario := ObtenerRolUsuario(r)

	// Banderas de permisos para el template (más seguro y simple que comparar en HTML).
	puedeCrear := TieneRol(r, "ADMIN", "OPERADOR")
	puedeEditar := TieneRol(r, "ADMIN", "OPERADOR")
	puedeEliminar := TieneRol(r, "ADMIN")

	// Estadísticas dashboard.
	var stats Stats
	queryStats := `
		SELECT
			(SELECT COUNT(*) FROM libros) AS total_libros,
			(SELECT COUNT(*) FROM libros WHERE formato = 'PDF') AS total_pdf,
			(SELECT COUNT(*) FROM libros WHERE formato = 'EPUB') AS total_epub,
			(SELECT COUNT(*) FROM libros WHERE formato = 'MOBI') AS total_mobi
	`
	err := h.DB.QueryRow(queryStats).Scan(
		&stats.TotalLibros,
		&stats.TotalPDF,
		&stats.TotalEPUB,
		&stats.TotalMOBI,
	)
	if err != nil {
		http.Error(w, "Error al consultar estadísticas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Consulta de libros con/sin búsqueda.
	var rows *sql.Rows
	if busqueda != "" {
		query := `
			SELECT id, titulo, autor, categoria, anio_publicacion, formato, stock_licencias
			FROM libros
			WHERE titulo LIKE ?
			ORDER BY id DESC
		`
		rows, err = h.DB.Query(query, "%"+busqueda+"%")
	} else {
		query := `
			SELECT id, titulo, autor, categoria, anio_publicacion, formato, stock_licencias
			FROM libros
			ORDER BY id DESC
		`
		rows, err = h.DB.Query(query)
	}
	if err != nil {
		http.Error(w, "Error al consultar libros: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var libros []models.Libro
	for rows.Next() {
		var libro models.Libro
		err := rows.Scan(
			&libro.ID,
			&libro.Titulo,
			&libro.Autor,
			&libro.Categoria,
			&libro.AnioPublicacion,
			&libro.Formato,
			&libro.StockLicencias,
		)
		if err != nil {
			http.Error(w, "Error al leer datos de libros: "+err.Error(), http.StatusInternalServerError)
			return
		}
		libros = append(libros, libro)
	}

	// Data para index.
	data := struct {
		Libros        []models.Libro
		Buscar        string
		Mensaje       string
		Stats         Stats
		UsuarioNombre string
		UsuarioRol    string
		PuedeCrear    bool
		PuedeEditar   bool
		PuedeEliminar bool
	}{
		Libros:        libros,
		Buscar:        busqueda,
		Mensaje:       mensaje,
		Stats:         stats,
		UsuarioNombre: nombreUsuario,
		UsuarioRol:    rolUsuario,
		PuedeCrear:    puedeCrear,
		PuedeEditar:   puedeEditar,
		PuedeEliminar: puedeEliminar,
	}

	err = h.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// NuevoLibroForm muestra el formulario para registrar un libro.
func (h *LibroHandler) NuevoLibroForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	err := h.Templates.ExecuteTemplate(w, "nuevo.html", nil)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla nuevo.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// CrearLibro guarda un nuevo libro.
func (h *LibroHandler) CrearLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al leer formulario", http.StatusBadRequest)
		return
	}

	titulo := strings.TrimSpace(r.FormValue("titulo"))
	autor := strings.TrimSpace(r.FormValue("autor"))
	categoria := strings.TrimSpace(r.FormValue("categoria"))
	formato := strings.TrimSpace(r.FormValue("formato"))

	anio, err := strconv.Atoi(r.FormValue("anio_publicacion"))
	if err != nil {
		http.Error(w, "Año de publicación inválido", http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("stock_licencias"))
	if err != nil {
		http.Error(w, "Stock de licencias inválido", http.StatusBadRequest)
		return
	}

	if titulo == "" || autor == "" || categoria == "" || formato == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO libros (titulo, autor, categoria, anio_publicacion, formato, stock_licencias)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = h.DB.Exec(query, titulo, autor, categoria, anio, formato, stock)
	if err != nil {
		http.Error(w, "Error al guardar libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?msg=Libro+creado+correctamente", http.StatusSeeOther)
}

// EditarLibroForm muestra formulario de edición.
func (h *LibroHandler) EditarLibroForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var libro models.Libro
	query := `
		SELECT id, titulo, autor, categoria, anio_publicacion, formato, stock_licencias
		FROM libros
		WHERE id = ?
	`
	err = h.DB.QueryRow(query, id).Scan(
		&libro.ID,
		&libro.Titulo,
		&libro.Autor,
		&libro.Categoria,
		&libro.AnioPublicacion,
		&libro.Formato,
		&libro.StockLicencias,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al consultar libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Templates.ExecuteTemplate(w, "editar.html", libro)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla editar.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// ActualizarLibro actualiza un libro existente.
func (h *LibroHandler) ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al leer formulario", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	titulo := strings.TrimSpace(r.FormValue("titulo"))
	autor := strings.TrimSpace(r.FormValue("autor"))
	categoria := strings.TrimSpace(r.FormValue("categoria"))
	formato := strings.TrimSpace(r.FormValue("formato"))

	anio, err := strconv.Atoi(r.FormValue("anio_publicacion"))
	if err != nil {
		http.Error(w, "Año de publicación inválido", http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("stock_licencias"))
	if err != nil {
		http.Error(w, "Stock de licencias inválido", http.StatusBadRequest)
		return
	}

	if titulo == "" || autor == "" || categoria == "" || formato == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE libros
		SET titulo = ?, autor = ?, categoria = ?, anio_publicacion = ?, formato = ?, stock_licencias = ?
		WHERE id = ?
	`
	_, err = h.DB.Exec(query, titulo, autor, categoria, anio, formato, stock, id)
	if err != nil {
		http.Error(w, "Error al actualizar libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?msg=Libro+actualizado+correctamente", http.StatusSeeOther)
}

// EliminarLibro elimina un libro por ID.
func (h *LibroHandler) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al leer formulario", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM libros WHERE id = ?`
	_, err = h.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Error al eliminar libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/?msg=Libro+eliminado+correctamente", http.StatusSeeOther)
}
