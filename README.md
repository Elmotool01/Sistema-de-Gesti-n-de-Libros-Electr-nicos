# Sistema de Gestión de Libros Electrónicos (Go)

## Descripción
Sistema para administrar un repositorio digital de libros electrónicos.
Incluye: usuarios (roles), catálogo, búsqueda y filtros, biblioteca personal, pago simulado y descarga segura, auditoría y reportes básicos.

## Módulos
- A: Autenticación y usuarios (registro/login/roles)
- B: Catálogo (CRUD admin)
- C: Búsqueda y filtros
- D: Biblioteca personal (favoritos, progreso)
- E: Operaciones (descargas/préstamos, trazabilidad)
- F: Archivos (PDF/EPUB, entrega segura)
- G: Reportes y auditoría
- H: Pago simulado y descarga

## Alcance
✅ Incluye pago simulado (estados: pendiente/aprobado/rechazado)  
❌ No incluye pasarela real, DRM avanzado, frontend complejo.

## Cómo ejecutar
```bash
go run ./cmd/app
