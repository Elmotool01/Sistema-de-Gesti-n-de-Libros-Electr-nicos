package main // Indica que este archivo pertenece al paquete principal del programa.

import "errors" // Importa el paquete errors para crear errores personalizados.

// ErrNoExiste se usa cuando se busca/actualiza/elimina un libro que no está registrado.
var ErrNoExiste = errors.New("libro no existe") // Error específico para “no encontrado”.

// ErrDuplicado se usa cuando se intenta agregar un libro con un ID ya existente.
var ErrDuplicado = errors.New("ya existe un libro con ese id") // Error específico para “duplicado”.

// ErrEntradaInvalida se usa cuando el usuario ingresa datos que no cumplen lo esperado.
var ErrEntradaInvalida = errors.New("entrada invalida") // Error genérico para entradas incorrectas.
