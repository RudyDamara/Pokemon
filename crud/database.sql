CREATE DATABASE pokemon;
-- Drop table

-- DROP TABLE public.my_catch;

CREATE TABLE public.my_catch (
	id serial NOT NULL,
	name varchar NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	url_pokemon varchar NULL,
	name_pokemon varchar NOT NULL
);
