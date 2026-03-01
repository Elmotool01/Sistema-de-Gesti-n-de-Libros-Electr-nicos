package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ConectarDB crea y valida la conexión con MySQL.
func ConectarDB() *sql.DB {
	usuario := os.Getenv("DB_USER")
	clave := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	puerto := os.Getenv("DB_PORT")
	nombreBD := os.Getenv("DB_NAME")

	// Valores por defecto (tu configuración actual)
	if usuario == "" {
		usuario = "root"
	}
	if clave == "" {
		clave = "Fevoutvx14@"
	}
	if host == "" {
		host = "127.0.0.1"
	}
	if puerto == "" {
		puerto = "3306"
	}
	if nombreBD == "" {
		nombreBD = "biblioteca_ebooks"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		usuario, clave, host, puerto, nombreBD)

	conexion, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error al abrir la conexión con MySQL: ", err)
	}

	if err = conexion.Ping(); err != nil {
		log.Fatal("❌ Error al conectar con MySQL (Ping): ", err)
	}

	log.Println("✅ Conexión exitosa a MySQL")
	return conexion
}
