package main // Paquete principal: punto de entrada de la aplicación.

import (
	"html/template"    // Paquete para cargar y renderizar plantillas HTML.
	"log"              // Paquete para imprimir mensajes en consola.
	"net/http"         // Paquete para crear servidor web y manejar rutas HTTP.
	"sistema/db"       // Paquete local para la conexión con MySQL.
	"sistema/handlers" // Paquete local con handlers de libros, auth y catálogo.
)

func main() {
	// =========================================================
	// 1) CONEXIÓN A LA BASE DE DATOS
	// =========================================================

	// Se crea la conexión a MySQL usando la función del paquete db.
	conexion := db.ConectarDB()

	// Se asegura que la conexión se cierre cuando termine la aplicación.
	defer conexion.Close()

	// =========================================================
	// 2) CARGA DE PLANTILLAS HTML
	// =========================================================

	// Se cargan todas las plantillas HTML de la carpeta "templates".
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		// Si falla la carga de plantillas, se detiene la ejecución.
		log.Fatal("❌ Error al cargar plantillas: ", err)
	}

	// =========================================================
	// 3) CREACIÓN DE HANDLERS (CONTROLADORES)
	// =========================================================

	// Handler del módulo de libros (CRUD + dashboard + búsqueda).
	libroHandler := handlers.NuevoLibroHandler(conexion, templates)

	// Handler del módulo de autenticación (login / logout / cookies).
	authHandler := handlers.NuevoAuthHandler(conexion, templates)

	// Handler del módulo catálogo (usuario lector).
	catalogoHandler := handlers.NuevoCatalogoHandler(conexion, templates)

	// =========================================================
	// 4) ARCHIVOS ESTÁTICOS (CSS, PDF demo, etc.)
	// =========================================================

	// Esta ruta permite servir archivos de la carpeta "static".
	// Ejemplo:
	//   URL:  /static/style.css
	//   File: static/style.css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// =========================================================
	// 5) RUTAS PÚBLICAS (AUTENTICACIÓN)
	// =========================================================

	// Ruta GET: muestra formulario de login.
	http.HandleFunc("/login", authHandler.MostrarLogin)

	// Ruta POST: procesa login (valida usuario/clave y crea cookies).
	http.HandleFunc("/login/procesar", authHandler.ProcesarLogin)

	// Ruta GET: cierra sesión y elimina cookies.
	http.HandleFunc("/logout", authHandler.Logout)

	// =========================================================
	// 6) RUTAS DEL CATÁLOGO (USUARIO LECTOR)
	//    Requieren login, pero no rol específico.
	// =========================================================

	// Ruta GET: muestra catálogo de libros.
	http.HandleFunc("/catalogo", RequiereLogin(catalogoHandler.VerCatalogo))

	// Ruta GET: muestra detalle de un libro.
	http.HandleFunc("/catalogo/detalle", RequiereLogin(catalogoHandler.VerDetalleLibro))

	// Ruta GET: descarga PDF demo del libro (flujo de demostración).
	http.HandleFunc("/catalogo/descargar", RequiereLogin(catalogoHandler.DescargarLibroDemo))

	// =========================================================
	// 7) RUTAS DEL PANEL ADMINISTRATIVO / CRUD DE LIBROS
	//    Requieren login + control por roles.
	// =========================================================

	// Ruta GET: panel principal (dashboard + listado + búsqueda).
	http.HandleFunc("/", RequiereLogin(libroHandler.Index))

	// Rutas CREATE (solo ADMIN y OPERADOR).
	http.HandleFunc("/libros/nuevo", RequiereLoginYRol(libroHandler.NuevoLibroForm, "ADMIN", "OPERADOR"))
	http.HandleFunc("/libros/crear", RequiereLoginYRol(libroHandler.CrearLibro, "ADMIN", "OPERADOR"))

	// Rutas UPDATE (solo ADMIN y OPERADOR).
	http.HandleFunc("/libros/editar", RequiereLoginYRol(libroHandler.EditarLibroForm, "ADMIN", "OPERADOR"))
	http.HandleFunc("/libros/actualizar", RequiereLoginYRol(libroHandler.ActualizarLibro, "ADMIN", "OPERADOR"))

	// Ruta DELETE (solo ADMIN).
	http.HandleFunc("/libros/eliminar", RequiereLoginYRol(libroHandler.EliminarLibro, "ADMIN"))

	// =========================================================
	// 8) INICIO DEL SERVIDOR WEB
	// =========================================================

	// Mensaje en consola con la URL del sistema.
	log.Println("🚀 Servidor iniciado en http://localhost:8082")

	// Inicia servidor HTTP en puerto 8082.
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		// Si falla el servidor, se muestra error y se detiene la app.
		log.Fatal("❌ Error al iniciar servidor: ", err)
	}
}

// =========================================================
// MIDDLEWARE: REQUIERE LOGIN
// =========================================================

// RequiereLogin protege rutas que exigen sesión iniciada.
// Recibe un handler y devuelve un handler protegido.
func RequiereLogin(next http.HandlerFunc) http.HandlerFunc {
	// Retorna una función wrapper.
	return func(w http.ResponseWriter, r *http.Request) {
		// Verifica si el usuario está autenticado usando cookies.
		if !handlers.EstaLogueado(r) {
			// Si no está logueado, redirige al login.
			http.Redirect(w, r, "/login?error=Debe+iniciar+sesión", http.StatusSeeOther)
			return
		}

		// Si está autenticado, ejecuta el handler original.
		next(w, r)
	}
}

// =========================================================
// MIDDLEWARE: REQUIERE LOGIN + ROL
// =========================================================

// RequiereLoginYRol protege rutas que exigen:
// 1) sesión iniciada
// 2) rol permitido (ej. ADMIN, OPERADOR)
func RequiereLoginYRol(next http.HandlerFunc, rolesPermitidos ...string) http.HandlerFunc {
	// Retorna una función wrapper.
	return func(w http.ResponseWriter, r *http.Request) {
		// -----------------------------------------------------
		// 1) VALIDAR SESIÓN
		// -----------------------------------------------------
		if !handlers.EstaLogueado(r) {
			// Si no hay sesión, redirige al login.
			http.Redirect(w, r, "/login?error=Debe+iniciar+sesión", http.StatusSeeOther)
			return
		}

		// -----------------------------------------------------
		// 2) VALIDAR ROL
		// -----------------------------------------------------
		if !handlers.TieneRol(r, rolesPermitidos...) {
			// Si no tiene permisos, redirige al panel con mensaje.
			http.Redirect(w, r, "/?msg=No+tiene+permisos+para+esa+acción", http.StatusSeeOther)
			return
		}

		// -----------------------------------------------------
		// 3) EJECUTAR HANDLER ORIGINAL
		// -----------------------------------------------------
		next(w, r)
	}
}
