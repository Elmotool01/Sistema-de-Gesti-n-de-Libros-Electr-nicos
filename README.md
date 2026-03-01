# Sistema de Gestión de Libros Electrónicos (Go + Web)

## Descripción del proyecto

Este proyecto consiste en un **Sistema de Gestión de Libros Electrónicos** desarrollado en **Go (Golang)** con una **interfaz web**.
Su propósito es permitir la organización, consulta y acceso a libros digitales mediante un catálogo, visualización de detalles, historial de acciones y gestión de usuarios.

El sistema fue diseñado con una estructura modular para facilitar su mantenimiento, ampliación e integración con nuevas funcionalidades (por ejemplo: autenticación, módulo administrativo completo y persistencia avanzada).

---

## Objetivo general

Desarrollar una aplicación web en Go que permita gestionar libros electrónicos de forma organizada, ofreciendo al usuario una interfaz para consultar catálogo, revisar detalles, acceder a lectura/descarga y registrar historial de uso.

---

## Objetivos específicos

* Implementar una **interfaz web** para visualizar libros electrónicos.
* Organizar el proyecto usando una **estructura modular** (handlers, models, templates, static).
* Definir **modelos de datos** para representar libros, historial y usuarios.
* Implementar rutas HTTP para la navegación del sistema.
* Permitir el acceso a funcionalidades como:

  * catálogo de libros
  * detalle de libro
  * lectura/descarga
  * historial de acciones
  * perfil de usuario
* Incorporar el concepto de **serialización con JSON** para manejo/intercambio de datos.
* Dejar una base para futuras mejoras como CRUD administrativo y base de datos persistente.

---

## Tecnologías utilizadas

* **Go (Golang)** – Lógica principal del sistema
* **HTML** – Estructura de vistas
* **CSS** – Estilos de interfaz
* **JSON** – Serialización/deserialización de datos (según requerimiento)
* **Git / GitHub** – Control de versiones y repositorio

---

## Funcionalidades implementadas

### Módulo de usuario

* Visualización de **catálogo de libros**
* **Búsqueda / navegación** de libros (según implementación)
* **Detalle de libro** (información completa)
* **Lectura o descarga** de archivo digital (según disponibilidad)
* **Historial** de acciones (lectura/descarga)
* **Perfil** de usuario

### Base para módulo administrativo (estructura preparada)

* Gestión de libros (CRUD) *(propuesta / ampliable)*
* Gestión de usuarios *(propuesta / ampliable)*
* Control de disponibilidad *(propuesta / ampliable)*

---

## Estructura del proyecto

La estructura del proyecto está organizada por responsabilidades:

* `main.go` → punto de entrada del sistema, configuración del servidor y rutas
* `handlers/` → lógica de atención de solicitudes HTTP
* `models/` → definición de estructuras de datos (Book, History, etc.)
* `templates/` → vistas HTML renderizadas en el navegador
* `static/` → archivos estáticos (CSS, imágenes, recursos)
* `data/` → archivos de datos (por ejemplo JSON)
* `db/` → lógica de acceso/persistencia (si aplica)
* `go.mod` → definición del módulo y dependencias

### Ejemplo de árbol de carpetas

```text
.
├── data/
├── db/
├── handlers/
├── models/
├── static/
├── templates/
├── main.go
├── go.mod
└── README.md
```

---

## Modelado de datos (ejemplo)

El sistema utiliza estructuras (structs) en Go para representar la información.
Ejemplo conceptual:

* **Book**: id, título, autor, categoría, descripción, portada, archivo, disponibilidad
* **History**: registro de acciones del usuario (lectura/descarga)
* **User**: información del usuario y rol *(si aplica en la versión actual)*

---

## Serialización con JSON (requisito)

### ¿Qué es serializar con JSON?

Serializar significa convertir los datos del programa (structs, slices, etc.) a un formato de texto estructurado como JSON, para poder:

* guardar datos en archivos
* intercambiar información entre sistemas
* preparar respuestas de API
* respaldar información del sistema

### En Go se usa:

* `json.Marshal()` / `json.MarshalIndent()` → **serialización**
* `json.Unmarshal()` → **deserialización**

### Uso en el proyecto

Se planteó el uso de JSON para almacenar o intercambiar datos del sistema (por ejemplo catálogo de libros o historial), cumpliendo con el requerimiento de que **la serialización se realiza mediante JSON**.

---

## Requisitos para ejecutar el proyecto

* Tener instalado **Go** (versión 1.20+ recomendada)
* Tener configurado **Git** (opcional, para clonar)
* Navegador web (Chrome, Edge, Firefox, etc.)

---

## Instalación y ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/TU-USUARIO/TU-REPOSITORIO.git
cd TU-REPOSITORIO
```

### 2. Verificar dependencias

```bash
go mod tidy
```

### 3. Ejecutar el proyecto

```bash
go run main.go
```

### 4. Abrir en el navegador

Ingresa a la URL que se muestre en consola (por ejemplo):

```text
http://localhost:8080
```

> **Nota:** El puerto puede variar según tu configuración en `main.go`.

---

## Flujo básico de uso

1. Iniciar el servidor con `go run main.go`
2. Abrir el sistema en el navegador
3. Ingresar al catálogo de libros
4. Seleccionar un libro para ver su detalle
5. Ejecutar acciones como lectura o descarga
6. Consultar el historial de uso
7. Revisar el perfil del usuario (si está habilitado)

---

## Arquitectura y organización del código

El proyecto se organizó con una estructura modular para mejorar la mantenibilidad:

* **Separación de responsabilidades**

  * Vistas (`templates`)
  * Lógica HTTP (`handlers`)
  * Datos (`models`)
  * Recursos (`static`)
* **Escalabilidad**

  * Permite añadir autenticación, base de datos y panel administrativo sin rehacer todo el sistema
* **Legibilidad**

  * Facilita la comprensión del proyecto por parte de docentes/evaluadores y otros desarrolladores

---

## Evidencias / demostración (para rúbrica)

Durante la presentación del proyecto se muestra:

* Estructura de carpetas y archivos
* Código principal (`main.go`)
* Rutas y handlers
* Catálogo de libros
* Detalle de libro
* Historial
* Navegación de la interfaz web

> Se recomienda complementar este repositorio con capturas de pantalla o un video demostrativo.

---

## Dificultades encontradas y solución

Una de las principales dificultades fue mantener compatibilidad entre la estructura existente del proyecto y nuevas mejoras (como serialización JSON o ampliación de modelos), sin afectar el funcionamiento actual.

### Solución aplicada

* Se trabajó con una estrategia de **integración incremental**
* Se evitó reemplazar archivos funcionales
* Se añadieron componentes de forma modular para no romper el sistema

---

## Resultados obtenidos

* Sistema web funcional para gestión de libros electrónicos
* Estructura organizada y escalable en Go
* Implementación de vistas y navegación básica
* Base para integrar persistencia adicional y módulo administrativo
* Cumplimiento del requerimiento conceptual de serialización JSON

---

## Posibles mejoras futuras

* Autenticación completa (login/registro)
* Roles de usuario (admin/usuario)
* CRUD completo de libros y usuarios
* Conexión a base de datos (MySQL/PostgreSQL)
* Panel administrativo con interfaz propia
* Búsqueda avanzada y filtros por categoría/autor
* Registro persistente de historial
* API REST para consumo externo
* Validaciones y manejo de errores más robusto

---

## Conclusión

Este proyecto permitió aplicar conocimientos de programación en Go, estructura modular de aplicaciones web, modelado de datos y organización de un sistema funcional orientado a la gestión de libros electrónicos. Además, deja una base sólida para futuras mejoras y para la incorporación de funcionalidades más avanzadas.

---

## Autor

**Aldebaran**
**por cuestion de seguridad**

---

## Licencia

Este proyecto se desarrolló con fines **académicos/educativos**.

*(Opcional: si deseas, puedes cambiar esta sección por una licencia MIT u otra.)*
