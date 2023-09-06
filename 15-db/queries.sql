/*


- выборка фильмов для нескольких режиссёров из списка (подзапрос);

- подсчёт количества фильмов для актёра;

- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;

- подсчёт количества фильмов со сборами больше 1000;

- подсчитать количество режиссёров, фильмы которых собрали больше 1000;

- выборка различных фамилий актёров;

- подсчёт количества фильмов, имеющих дубли по названию.
*/

-- выборка фильмов с названием студии;
select movies.title, studios.title 
from movies 
join studios 
on movies.studio_id = studios.id;

-- выборка фильмов для некоторого актёра
select movies.title , actors.first_name , actors.last_name 
from movies 
join movies_actors 
on movies_actors.movie_id = movies.id
join actors 
on actors.id = movies_actors.actor_id
where actors.first_name = 'Mike' and actors.last_name = 'Second';

-- подсчёт фильмов для некоторого режиссёра
select  
    directors.first_name, 
    directors.last_name,
    count (movies.id)
from movies 
join movies_directors 
on movies_directors.movie_id = movies.id
join directors 
on directors.id = movies_directors.director_id
where directors.first_name = 'Tafu' and directors.last_name = 'Kerutto'
group by directors.first_name, directors.last_name;

