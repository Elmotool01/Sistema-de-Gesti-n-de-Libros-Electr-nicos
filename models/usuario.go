package models // Paquete models: contiene estructuras de datos del sistema.

// Usuario representa a un usuario del sistema que puede iniciar sesión.
type Usuario struct {
	// IDUsuario guarda el identificador único del usuario en la base de datos.
	IDUsuario int

	// Nombre guarda el nombre del usuario.
	Nombre string

	// Correo guarda el correo electrónico usado para iniciar sesión.
	Correo string

	// Clave guarda la contraseña (en esta fase se usa texto simple para pruebas).
	Clave string

	// IDRol guarda el identificador del rol del usuario.
	IDRol int

	// NombreRol guarda el nombre del rol (ADMIN, OPERADOR, CONSULTA).
	NombreRol string

	// Estado guarda si el usuario está ACTIVO o INACTIVO.
	Estado string
}
