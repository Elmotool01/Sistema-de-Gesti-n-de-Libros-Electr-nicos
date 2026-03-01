package handlers // Paquete handlers: contiene controladores del sistema.

import (
	"database/sql"   // Paquete para trabajar con SQL.
	"html/template"  // Paquete para renderizar plantillas HTML.
	"net/http"       // Paquete para rutas, respuestas y descarga de archivos.
	"os"             // Paquete para validar existencia de archivos.
	"path/filepath"  // Paquete para manejar rutas de archivos.
	"sistema/models" // Estructuras del sistema (Libro).
	"strconv"        // Paquete para convertir string a int.
	"strings"        // Paquete para limpiar texto.
)

// CatalogoHandler maneja las vistas del catálogo para usuario lector.
type CatalogoHandler struct {
	DB        *sql.DB            // Conexión a la base de datos.
	Templates *template.Template // Plantillas HTML cargadas.
}

// NuevoCatalogoHandler crea una nueva instancia del handler de catálogo.
func NuevoCatalogoHandler(db *sql.DB, templates *template.Template) *CatalogoHandler {
	return &CatalogoHandler{
		DB:        db,
		Templates: templates,
	}
}

// VerCatalogo muestra el catálogo de libros para el usuario lector.
// Ruta: GET /catalogo
func (h *CatalogoHandler) VerCatalogo(w http.ResponseWriter, r *http.Request) {
	// Solo permitir GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtiene texto de búsqueda opcional.
	busqueda := strings.TrimSpace(r.URL.Query().Get("buscar"))

	// Obtiene datos del usuario desde cookies.
	nombreUsuario := ObtenerNombreUsuario(r)
	rolUsuario := ObtenerRolUsuario(r)

	// Variables para la consulta.
	var (
		rows *sql.Rows
		err  error
	)

	// Si hay búsqueda, filtra por título o categoría.
	if busqueda != "" {
		query := `
			SELECT id, titulo, autor, categoria, anio_publicacion, formato, stock_licencias
			FROM libros
			WHERE titulo LIKE ? OR categoria LIKE ?
			ORDER BY titulo ASC
		`
		filtro := "%" + busqueda + "%"
		rows, err = h.DB.Query(query, filtro, filtro)
	} else {
		// Si no hay búsqueda, lista todos los libros.
		query := `
			SELECT id, titulo, autor, categoria, anio_publicacion, formato, stock_licencias
			FROM libros
			ORDER BY titulo ASC
		`
		rows, err = h.DB.Query(query)
	}

	// Manejo de error en consulta.
	if err != nil {
		http.Error(w, "Error al consultar catálogo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice de libros del catálogo.
	var libros []models.Libro

	// Recorre resultados de la consulta.
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
			http.Error(w, "Error al leer datos del catálogo: "+err.Error(), http.StatusInternalServerError)
			return
		}

		libros = append(libros, libro)
	}

	// Data para la plantilla catalogo.html.
	data := struct {
		Libros        []models.Libro // Lista de libros para mostrar.
		Buscar        string         // Texto del buscador.
		UsuarioNombre string         // Nombre del usuario logueado.
		UsuarioRol    string         // Rol del usuario logueado.
	}{
		Libros:        libros,
		Buscar:        busqueda,
		UsuarioNombre: nombreUsuario,
		UsuarioRol:    rolUsuario,
	}

	// Renderiza la plantilla catalogo.html.
	err = h.Templates.ExecuteTemplate(w, "catalogo.html", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla catalogo.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// VerDetalleLibro muestra el detalle de un libro seleccionado del catálogo.
// Ruta: GET /catalogo/detalle?id=...
func (h *CatalogoHandler) VerDetalleLibro(w http.ResponseWriter, r *http.Request) {
	// Solo permitir GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtiene ID desde URL.
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de libro inválido", http.StatusBadRequest)
		return
	}

	// Variable para cargar libro.
	var libro models.Libro

	// Consulta libro por ID.
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
		http.Error(w, "Error al consultar detalle del libro: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Data para detalle_libro.html.
	data := struct {
		Libro         models.Libro // Libro seleccionado.
		UsuarioNombre string       // Usuario actual.
		UsuarioRol    string       // Rol actual.
	}{
		Libro:         libro,
		UsuarioNombre: ObtenerNombreUsuario(r),
		UsuarioRol:    ObtenerRolUsuario(r),
	}

	// Renderiza detalle_libro.html.
	err = h.Templates.ExecuteTemplate(w, "detalle_libro.html", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla detalle_libro.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// DescargarLibroDemo descarga un PDF de demostración para simular la descarga del libro.
// Ruta: GET /catalogo/descargar?id=...
func (h *CatalogoHandler) DescargarLibroDemo(w http.ResponseWriter, r *http.Request) {
	// Solo permitir GET para la descarga.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtiene ID del libro (solo para validar que el libro existe y para personalizar nombre).
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de libro inválido", http.StatusBadRequest)
		return
	}

	// Consulta el libro para validar existencia y obtener título.
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
		http.Error(w, "Error al validar libro para descarga: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ruta del archivo PDF de demostración (misma para todos los libros en esta fase).
	archivoDemo := filepath.Join("static", "demo", "demo.pdf")

	// Verifica que el archivo exista.
	if _, err := os.Stat(archivoDemo); err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Archivo PDF de demostración no encontrado. Coloque demo.pdf en static/demo/", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al verificar archivo de descarga: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Genera un nombre sugerido para la descarga (solo demostrativo).
	// Se reemplazan espacios por guiones bajos.
	nombreDescarga := strings.ReplaceAll(libro.Titulo, " ", "_") + "_demo.pdf"

	// Configura encabezados para forzar descarga en navegador.
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="`+nombreDescarga+`"`)

	// Envía el archivo al cliente.
	http.ServeFile(w, r, archivoDemo)
}
