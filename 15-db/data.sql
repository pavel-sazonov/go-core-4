-- Первичное наполнение БД данными.

INSERT INTO studios (title) VALUES ('Columbia Pictures');

INSERT INTO 
    actors (first_name, last_name, date_of_birth) 
VALUES 
    ('Joe', 'First', '1975-12-03'),
    ('Mike', 'Second', '1980-12-03'),
    ('Chloe', 'Abrams', '2012-07-03'),
    ('Joel', 'Numberone', '2008-07-09'),
    ('Kate', 'Juddieh', '1995-07-03');

INSERT INTO 
    directors (first_name, last_name, date_of_birth) 
VALUES 
    ('Lou', 'Gileo', '1958-04-03'),
    ('Tafu', 'Kerutto', '1975-04-12'),
    ('Hanna', 'Okreo', '2001-04-03');

INSERT INTO 
    movies (title, year_of_release, box_office, rating, studio_id) 
VALUES 
    ('First movie', 1979, 1000000, 'PG-10', 1),
    ('Second movie', 1980, 2000000, 'PG-13', 1),
    ('Third movie', 1995, 100000, 'PG-18', 1),
    ('Fourth movie', 1979, 1000000, 'PG-13', 1);

INSERT INTO 
    movies_actors (movie_id, actor_id) 
VALUES 
    (1, 2),
    (1, 3),
    (1, 4),
    (2, 1),
    (2, 5),
    (3, 5),
    (3, 2),
    (4, 3),
    (4, 4);

INSERT INTO 
    movies_directors (movie_id, director_id) 
VALUES 
    (1, 2),
    (2, 3),
    (3, 1),
    (4, 2);