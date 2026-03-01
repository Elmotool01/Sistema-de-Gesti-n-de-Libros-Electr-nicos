package handlers // Paquete handlers: contiene controladores del sistema.

import (
	"database/sql"   // Paquete para trabajar con bases de datos SQL.
	"html/template"  // Paquete para renderizar plantillas HTML.
	"net/http"       // Paquete para servidor web, rutas y cookies.
	"sistema/models" // Importa la estructura Usuario.
	"strings"        // Paquete para limpiar y comparar textos.
)

// AuthHandler agrupa los recursos necesarios para autenticación.
type AuthHandler struct {
	DB        *sql.DB            // Conexión a la base de datos.
	Templates *template.Template // Plantillas HTML cargadas.
}

// NuevoAuthHandler crea una nueva instancia del handler de autenticación.
func NuevoAuthHandler(db *sql.DB, templates *template.Template) *AuthHandler {
	return &AuthHandler{
		DB:        db,        // Guarda la conexión MySQL.
		Templates: templates, // Guarda las plantillas HTML.
	}
}

// MostrarLogin renderiza la plantilla login.html.
// Ruta: GET /login
func (h *AuthHandler) MostrarLogin(w http.ResponseWriter, r *http.Request) {
	// Solo se permite método GET para mostrar el formulario.
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Si el usuario ya está logueado, se redirige al panel principal.
	if EstaLogueado(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Lee mensaje de error/información desde la URL (ej: ?error=Credenciales+inválidas).
	errorMsg := strings.TrimSpace(r.URL.Query().Get("error"))

	// Estructura de datos para la plantilla login.html.
	data := struct {
		Error string // Mensaje de error o aviso.
	}{
		Error: errorMsg,
	}

	// Renderiza la plantilla login.html.
	err := h.Templates.ExecuteTemplate(w, "login.html", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantilla login.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// ProcesarLogin valida credenciales y crea sesión con cookies.
// Ruta: POST /login/procesar
func (h *AuthHandler) ProcesarLogin(w http.ResponseWriter, r *http.Request) {
	// Solo se permite método POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lee el formulario enviado desde login.html.
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al leer formulario", http.StatusBadRequest)
		return
	}

	// Obtiene y limpia los valores enviados por el usuario.
	correo := strings.TrimSpace(r.FormValue("correo"))
	clave := strings.TrimSpace(r.FormValue("clave"))

	// Validación básica: ambos campos son obligatorios.
	if correo == "" || clave == "" {
		http.Redirect(w, r, "/login?error=Debe+ingresar+correo+y+clave", http.StatusSeeOther)
		return
	}

	// Variable para cargar los datos del usuario si las credenciales son válidas.
	var usuario models.Usuario

	// Consulta SQL para validar correo + clave + estado ACTIVO.
	// Nota: En esta fase la clave se compara como texto plano (demo académica).
	query := `
		SELECT 
			u.id_usuario,
			u.nombre,
			u.correo,
			u.id_rol,
			r.nombre_rol,
			u.estado
		FROM usuarios u
		INNER JOIN roles r ON u.id_rol = r.id_rol
		WHERE u.correo = ? AND u.clave = ? AND u.estado = 'ACTIVO'
		LIMIT 1
	`

	// Ejecuta la consulta y llena la estructura usuario.
	err = h.DB.QueryRow(query, correo, clave).Scan(
		&usuario.IDUsuario,
		&usuario.Nombre,
		&usuario.Correo,
		&usuario.IDRol,
		&usuario.NombreRol,
		&usuario.Estado,
	)

	// Si no coincide usuario/clave, redirige al login con error.
	if err != nil {
		if err == sql.ErrNoRows {
			http.Redirect(w, r, "/login?error=Credenciales+inválidas", http.StatusSeeOther)
			return
		}

		// Si ocurre otro error, responde 500.
		http.Error(w, "Error al validar usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// =========================================================
	// CREACIÓN DE SESIÓN SIMPLE CON COOKIES
	// =========================================================

	// Cookie principal: indica que el usuario está autenticado.
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_logueado", // Nombre de la cookie.
		Value:    "true",             // Valor que representa sesión activa.
		Path:     "/",                // Disponible en todo el sitio.
		HttpOnly: true,               // Evita acceso desde JavaScript.
	})

	// Cookie con el nombre del usuario (para mostrar en interfaz).
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_nombre",
		Value:    usuario.Nombre,
		Path:     "/",
		HttpOnly: true,
	})

	// Cookie con el rol del usuario (ADMIN / OPERADOR / CONSULTA).
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_rol",
		Value:    usuario.NombreRol,
		Path:     "/",
		HttpOnly: true,
	})

	// =========================================================
	// REDIRECCIÓN SEGÚN ROL
	// =========================================================

	// Si el usuario tiene rol CONSULTA, se envía directamente al catálogo.
	// Esto permite que el usuario lector vea catálogo, detalle y descarga.
	if strings.ToUpper(strings.TrimSpace(usuario.NombreRol)) == "CONSULTA" {
		http.Redirect(w, r, "/catalogo", http.StatusSeeOther)
		return
	}

	// Si es ADMIN u OPERADOR, se envía al panel principal.
	http.Redirect(w, r, "/?msg=Bienvenido+"+usuario.Nombre, http.StatusSeeOther)
}

// Logout elimina cookies de sesión y redirige al login.
// Ruta: GET /logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Elimina cookie de estado de login.
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_logueado",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // MaxAge negativo elimina la cookie.
		HttpOnly: true,
	})

	// Elimina cookie de nombre de usuario.
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_nombre",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Elimina cookie de rol de usuario.
	http.SetCookie(w, &http.Cookie{
		Name:     "usuario_rol",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Redirige al login con mensaje informativo.
	http.Redirect(w, r, "/login?error=Sesión+cerrada+correctamente", http.StatusSeeOther)
}

// =========================================================
// FUNCIONES AUXILIARES DE SESIÓN Y ROLES
// =========================================================

// EstaLogueado verifica si existe la cookie de sesión activa.
func EstaLogueado(r *http.Request) bool {
	// Busca la cookie de login.
	cookie, err := r.Cookie("usuario_logueado")
	if err != nil {
		// Si no existe cookie, no hay sesión.
		return false
	}

	// Si existe y su valor es "true", se considera logueado.
	return cookie.Value == "true"
}

// ObtenerNombreUsuario obtiene el nombre del usuario desde cookie.
// Si no existe cookie, devuelve cadena vacía.
func ObtenerNombreUsuario(r *http.Request) string {
	cookie, err := r.Cookie("usuario_nombre")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// ObtenerRolUsuario obtiene el rol del usuario desde cookie.
// Si no existe cookie, devuelve cadena vacía.
func ObtenerRolUsuario(r *http.Request) string {
	cookie, err := r.Cookie("usuario_rol")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// TieneRol verifica si el usuario tiene alguno de los roles permitidos.
// Ejemplo de uso: TieneRol(r, "ADMIN", "OPERADOR")
func TieneRol(r *http.Request, rolesPermitidos ...string) bool {
	// Obtiene el rol actual desde cookie y lo normaliza a mayúsculas.
	rolActual := strings.ToUpper(strings.TrimSpace(ObtenerRolUsuario(r)))

	// Si no existe rol, no tiene permisos.
	if rolActual == "" {
		return false
	}

	// Recorre la lista de roles permitidos.
	for _, rolPermitido := range rolesPermitidos {
		// Compara normalizando mayúsculas/espacios.
		if rolActual == strings.ToUpper(strings.TrimSpace(rolPermitido)) {
			return true
		}
	}

	// Si no coincidió con ninguno, no tiene permiso.
	return false
}
