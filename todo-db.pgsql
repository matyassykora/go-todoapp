--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: todo-db; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "todo-db" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';


ALTER DATABASE "todo-db" OWNER TO postgres;

\connect -reuse-previous=on "dbname='todo-db'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: todos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.todos (
    name text,
    id uuid,
    done boolean,
    pk integer NOT NULL,
    pos integer
);


ALTER TABLE public.todos OWNER TO postgres;

--
-- Name: todos_pk_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.todos ALTER COLUMN pk ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.todos_pk_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: todos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.todos (name, id, done, pk, pos) FROM stdin;
Watch a movie	a5b8e513-6ab2-4b48-b19c-ffb34bcc2659	t	95	\N
Do a thing	d7badef5-7c0a-4816-861f-ba4e60e4cecf	t	56	\N
Finish a task	d67796da-a32c-42c3-a524-ce67155ee346	f	64	\N
\.


--
-- Name: todos_pk_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.todos_pk_seq', 104, true);


--
-- Name: todos todos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.todos
    ADD CONSTRAINT todos_pkey PRIMARY KEY (pk);


--
-- PostgreSQL database dump complete
--

