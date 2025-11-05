--
-- PostgreSQL database dump
--

\restrict lWksdMx4dlwur1jjUTZqwvAmOwGsd065g8fbW6Aik7EmNvu3rrtqvkbi2PhROhO

-- Dumped from database version 17.6
-- Dumped by pg_dump version 17.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    title text NOT NULL,
    done boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tasks_id_seq OWNER TO postgres;

--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (id, title, done, created_at) FROM stdin;
1	Первая задача из psql	f	2025-11-05 20:18:24.567892+03
2	Сделать ПЗ №5	f	2025-11-05 20:22:06.41897+03
3	Купить кофе	f	2025-11-05 20:22:06.420432+03
4	Проверить отчёты	f	2025-11-05 20:22:06.42085+03
5	Сделать ПЗ №5	f	2025-11-05 20:33:25.315853+03
6	Купить кофе	f	2025-11-05 20:33:25.323093+03
7	Проверить отчёты	f	2025-11-05 20:33:25.323745+03
8	Задача из транзакции 1	f	2025-11-05 20:33:25.331841+03
9	Задача из транзакции 2	f	2025-11-05 20:33:25.331841+03
10	Завершенная задача (тест)	t	2025-11-05 20:33:25.332856+03
\.


--
-- Name: tasks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tasks_id_seq', 10, true);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\unrestrict lWksdMx4dlwur1jjUTZqwvAmOwGsd065g8fbW6Aik7EmNvu3rrtqvkbi2PhROhO

