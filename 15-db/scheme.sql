/*
- фильмы (название, год выхода, актёры, режиссёры, сборы, рейтинг);

- актёр (имя, дата рождения);

- режиссёры (имя, дата рождения);

- студии (название).

У каждого фильма может быть много актёров и режиссёров, но одна студия.

Рейтинг – значение из набора PG-10, PG-13, PG-18.

Год выхода не может быть меньше 1800.

В один год не может быть двух фильмов с одинаковым названием.
*/

DROP TABLE IF EXISTS movies_directors;
DROP TABLE IF EXISTS movies_actors;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS studios;

-- студии
CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL DEFAULT ''
);

-- актеры
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    date_of_birth DATE NOT NULL
);

-- режиссеры
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    date_of_birth DATE NOT NULL
);

-- фильмы
CREATE TABLE movies (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL DEFAULT '',
    year_of_release INTEGER CHECK (year_of_release >= 1800) NOT NULL DEFAULT 0,
    box_office INTEGER NOT NULL DEFAULT 0,
    rating VARCHAR(5) NOT NULL DEFAULT 'PG-10', 
    studio_id INTEGER REFERENCES studios(id)
);

CREATE INDEX IF NOT EXISTS movies_title_idx ON movies USING btree (lower(title));

-- связь между фильмами и актерами
CREATE TABLE movies_actors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    movie_id BIGINT NOT NULL REFERENCES movies(id),
    actor_id INTEGER NOT NULL REFERENCES actors(id),
    UNIQUE(movie_id, actor_id)
);

-- связь между фильмами и режиссерами
CREATE TABLE movies_directors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    movie_id BIGINT NOT NULL REFERENCES movies(id),
    director_id INTEGER NOT NULL REFERENCES directors(id),
    UNIQUE(movie_id, director_id)
);

-- функция-триггер для проверки: 
-- в один год не может быть двух фильмов с одинаковым назнванием
CREATE OR REPLACE FUNCTION check_moview_year_and_title()
  RETURNS TRIGGER AS $$
BEGIN
    IF NEW.year_of_release = (SELECT year_of_release FROM movies WHERE year_of_release = NEW.year_of_release) AND
        NEW.title = (SELECT title FROM movies WHERE year_of_release = NEW.year_of_release)
        THEN RAISE EXCEPTION 'Invalid movie title'; --RETURN NULL;
        ELSE RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;
-- регистрация тригера для таблицы
CREATE OR REPLACE TRIGGER check_moview_year_and_title BEFORE INSERT OR UPDATE ON movies 
FOR EACH ROW EXECUTE PROCEDURE check_moview_year_and_title();
