# Sistema de Gestión de Libros Electrónicos (Go + JSON)

## Descripción del proyecto

Este proyecto consiste en un **Sistema de Gestión de Libros Electrónicos** desarrollado en **Go (Golang)**, orientado a la administración y consulta de libros digitales.

El sistema permite trabajar con libros electrónicos usando **persistencia en memoria** y **persistencia en archivo JSON**, cumpliendo con el requerimiento de que **la serialización se realiza mediante JSON**.

---

## Objetivo general

Desarrollar una aplicación en Go para gestionar libros electrónicos, aplicando estructuras de datos, separación de responsabilidades y serialización/deserialización en formato JSON.

---

## Objetivos específicos

* Implementar un sistema para registrar, consultar y administrar libros electrónicos.
* Aplicar programación estructurada/modular en Go.
* Manejar persistencia en:

  * **memoria**
  * **archivo JSON**
* Implementar serialización y deserialización de datos con JSON.
* Organizar el código en componentes reutilizables (modelo, repositorios, interfaz y lógica principal).
* Documentar el proyecto para su ejecución y evaluación.

---

## Tecnologías utilizadas

* **Go (Golang)**
* **JSON** (serialización y almacenamiento de datos)
* **Git / GitHub**

---

## Funcionalidades del sistema

* Registro de libros electrónicos
* Consulta/listado de libros
* Gestión de datos desde repositorio en memoria
* Gestión de datos desde repositorio JSON
* Carga de libros desde archivo `libros.json`
* Manejo de errores (archivo `errores.go`)
* Interfaz/flujo de interacción del sistema (archivo `ui.go`)

> Las funcionalidades exactas pueden variar según la versión del código cargada en el repositorio.

---

## Estructura real del proyecto

Según el repositorio actual, el proyecto contiene archivos principales como: `main.go`, `libro.go`, `repo_json.go`, `repo_memoria.go`, `ui.go`, `errores.go` y `libros.json`, además de diagramas e imágenes. ([GitHub][1])

### Árbol de archivos (referencial)

```text
.
├── main.go
├── libro.go
├── repo_json.go
├── repo_memoria.go
├── ui.go
├── errores.go
├── libros.json
├── README.md
├── Diagrama inicial.png
└── Diagrama casos de uso.png
```

---

## Explicación de archivos principales

### `main.go`

Punto de entrada del programa.
Coordina la ejecución del sistema y el flujo principal.

### `libro.go`

Define la estructura de datos del libro (modelo), por ejemplo atributos como título, autor, categoría, etc.

### `repo_memoria.go`

Implementa el repositorio en memoria para manejar libros temporalmente durante la ejecución.

### `repo_json.go`

Implementa el repositorio con persistencia en JSON:

* lectura del archivo JSON
* escritura del archivo JSON
* serialización y deserialización de datos

### `libros.json`

Archivo de almacenamiento de datos en formato JSON.

### `ui.go`

Maneja la interacción con el usuario (menús, mensajes, entradas, salidas).

### `errores.go`

Centraliza o define el manejo de errores del sistema.

---

## Serialización con JSON (requisito)

### ¿Qué significa serializar con JSON?

Serializar significa convertir los datos del programa (por ejemplo, structs y listas en Go) a un formato de texto estructurado como JSON.

### ¿Qué significa deserializar?

Es el proceso inverso: convertir un JSON a estructuras de Go para poder trabajar con los datos en el programa.

### En Go se utiliza:

* `json.Marshal()` / `json.MarshalIndent()` → serializar
* `json.Unmarshal()` → deserializar

### Aplicación en este proyecto

En este sistema, la serialización JSON se utiliza para almacenar y recuperar la información de los libros desde el archivo `libros.json`, cumpliendo el requerimiento de persistencia mediante JSON.

---

## Requisitos para ejecutar el proyecto

* Tener instalado **Go** (recomendado Go 1.20 o superior)
* Tener **Git** instalado (opcional, para clonar el repositorio)

---

## Instalación y ejecución

### 1) Clonar el repositorio

```bash
git clone "https://github.com/Elmotool01/Sistema-de-Gesti-n-de-Libros-Electr-nicos.git"
cd "Sistema-de-Gesti-n-de-Libros-Electr-nicos"
```

> Si tu sistema tiene problemas con caracteres especiales/codificación del nombre de carpeta, también puedes descargar el proyecto como ZIP desde GitHub.

### 2) Verificar dependencias del módulo

```bash
go mod tidy
```

### 3) Ejecutar el programa

```bash
go run .
```

> También puede funcionar:

```bash
go run main.go
```

> dependiendo de cómo estén distribuidos los archivos y paquetes en tu proyecto.

---

## Ejemplo de uso (flujo general)

1. Ejecutar el programa
2. Mostrar menú/interfaz
3. Seleccionar opción (registrar/listar/consultar libros)
4. Guardar cambios en memoria o en archivo JSON
5. Cerrar el programa

---

## Diagramas del proyecto

El repositorio incluye diagramas de apoyo para el análisis y diseño del sistema, como:

* `Diagrama inicial.png`
* `Diagrama casos de uso.png` ([GitHub][1])

Estos diagramas ayudan a explicar el funcionamiento general del sistema y la interacción esperada del usuario.

---

## Cumplimiento académico (rúbrica) – resumen

Este proyecto evidencia:

* ✅ Desarrollo en **Go**
* ✅ Organización del código por responsabilidades
* ✅ Implementación de lógica de gestión de libros
* ✅ Persistencia en **JSON**
* ✅ Serialización / deserialización
* ✅ Documentación del proyecto (README)
* ✅ Material de apoyo (diagramas)

---

## Dificultades encontradas (ejemplo)

Durante el desarrollo se presentaron retos relacionados con:

* estructuración del proyecto
* compatibilidad de cambios
* integración de persistencia JSON sin afectar el flujo funcional

### Solución aplicada

Se trabajó de forma incremental, separando la lógica en archivos específicos (modelo, repositorio, UI y errores), permitiendo mantener una base más clara y escalable.

---

## Posibles mejoras futuras

* Interfaz gráfica/web completa
* CRUD más robusto con validaciones
* Búsqueda por autor/categoría/título
* Login y roles de usuario
* Conexión a base de datos (MySQL/PostgreSQL)
* API REST en Go
* Pruebas unitarias

---

## Autor

**Aldebarán Centurión** *(editar si deseas otro nombre)*
Proyecto académico – Gestión de Libros Electrónicos

---

## Repositorio

GitHub: `Elmotool01/Sistema-de-Gesti-n-de-Libros-Electr-nicos` ([GitHub][1])

---

[1]: https://github.com/Elmotool01/Sistema-de-Gesti-n-de-Libros-Electr-nicos "GitHub - Elmotool01/Sistema-de-Gesti-n-de-Libros-Electr-nicos"
