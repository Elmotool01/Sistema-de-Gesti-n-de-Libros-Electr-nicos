USE biblioteca_ebooks;

-- OPCIONAL: limpiar la tabla antes de insertar (evita duplicados)
-- TRUNCATE TABLE libros;

INSERT INTO libros (titulo, autor, categoria, anio_publicacion, formato, stock_licencias)
VALUES
('Clean Code', 'Robert C. Martin', 'Programación', 2008, 'PDF', 5),
('The Pragmatic Programmer', 'Andrew Hunt', 'Programación', 1999, 'EPUB', 3),
('Refactoring', 'Martin Fowler', 'Programación', 2018, 'PDF', 4),
('Code Complete', 'Steve McConnell', 'Programación', 2004, 'PDF', 2),
('Design Patterns', 'Erich Gamma', 'Programación', 1994, 'EPUB', 3),
('Eloquent JavaScript', 'Marijn Haverbeke', 'Programación', 2018, 'MOBI', 6),
('JavaScript: The Good Parts', 'Douglas Crockford', 'Programación', 2008, 'PDF', 2),
('You Don''t Know JS Yet', 'Kyle Simpson', 'Programación', 2020, 'EPUB', 4),

('Introducción a Algoritmos', 'Thomas H. Cormen', 'Informática', 2009, 'MOBI', 2),
('Algoritmos', 'Robert Sedgewick', 'Informática', 2011, 'PDF', 3),
('Computer Networks', 'Andrew S. Tanenbaum', 'Informática', 2010, 'EPUB', 2),
('Operating System Concepts', 'Abraham Silberschatz', 'Informática', 2018, 'PDF', 2),
('Compilers: Principles, Techniques, and Tools', 'Alfred V. Aho', 'Informática', 2006, 'MOBI', 1),

('Sistemas de Bases de Datos', 'Abraham Silberschatz', 'Bases de Datos', 2011, 'PDF', 3),
('Database System Concepts', 'Henry F. Korth', 'Bases de Datos', 2010, 'EPUB', 2),
('Learning SQL', 'Alan Beaulieu', 'Bases de Datos', 2020, 'PDF', 5),
('SQL Antipatterns', 'Bill Karwin', 'Bases de Datos', 2010, 'MOBI', 2),

('Don Quijote de la Mancha', 'Miguel de Cervantes', 'Literatura', 1605, 'PDF', 2),
('Cien Años de Soledad', 'Gabriel García Márquez', 'Novela', 1967, 'EPUB', 4),
('La Odisea', 'Homero', 'Literatura', 800, 'PDF', 1),
('1984', 'George Orwell', 'Novela', 1949, 'MOBI', 5),
('Fahrenheit 451', 'Ray Bradbury', 'Novela', 1953, 'EPUB', 3),
('Crimen y Castigo', 'Fiódor Dostoyevski', 'Novela', 1866, 'PDF', 2),
('Orgullo y Prejuicio', 'Jane Austen', 'Novela', 1813, 'EPUB', 4),

('El Arte de la Guerra', 'Sun Tzu', 'Historia', 500, 'PDF', 3),
('Sapiens', 'Yuval Noah Harari', 'Historia', 2011, 'EPUB', 6),
('Breve Historia del Tiempo', 'Stephen Hawking', 'Ciencia', 1988, 'PDF', 4),
('Cosmos', 'Carl Sagan', 'Ciencia', 1980, 'MOBI', 3),
('El Gen Egoísta', 'Richard Dawkins', 'Ciencia', 1976, 'EPUB', 2),
('Thinking, Fast and Slow', 'Daniel Kahneman', 'Psicología', 2011, 'PDF', 5);

-- Verificación
SELECT COUNT(*) AS total_libros FROM libros;
SELECT * FROM libros ORDER BY id;