--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.3 (Debian 16.3-1.pgdg120+1)

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
-- Name: blocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.blocks (
    id integer NOT NULL,
    title character varying(100) NOT NULL
);


ALTER TABLE public.blocks OWNER TO postgres;

--
-- Name: blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.blocks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blocks_id_seq OWNER TO postgres;

--
-- Name: blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.blocks_id_seq OWNED BY public.blocks.id;


--
-- Name: departments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.departments (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    block_id integer NOT NULL,
    office_id integer NOT NULL
);


ALTER TABLE public.departments OWNER TO postgres;

--
-- Name: departaments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.departaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.departaments_id_seq OWNER TO postgres;

--
-- Name: departaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.departaments_id_seq OWNED BY public.departments.id;


--
-- Name: divisions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.divisions (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    department_id integer NOT NULL
);


ALTER TABLE public.divisions OWNER TO postgres;

--
-- Name: divisions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.divisions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.divisions_id_seq OWNER TO postgres;

--
-- Name: divisions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.divisions_id_seq OWNED BY public.divisions.id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    id integer NOT NULL,
    first_name character varying(50) NOT NULL,
    middle_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    division_id integer NOT NULL,
    role_id integer NOT NULL,
    position_id integer NOT NULL,
    project_id integer NOT NULL,
    head_id integer
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.employees_id_seq OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- Name: offices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.offices (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    city character varying(255) NOT NULL
);


ALTER TABLE public.offices OWNER TO postgres;

--
-- Name: officies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.officies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.officies_id_seq OWNER TO postgres;

--
-- Name: officies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.officies_id_seq OWNED BY public.offices.id;


--
-- Name: positions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.positions (
    id integer NOT NULL,
    title character varying(100) NOT NULL
);


ALTER TABLE public.positions OWNER TO postgres;

--
-- Name: positions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.positions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.positions_id_seq OWNER TO postgres;

--
-- Name: positions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.positions_id_seq OWNED BY public.positions.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id integer NOT NULL,
    title character varying(255) NOT NULL
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_id_seq OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    title character varying(100) NOT NULL
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: blocks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blocks ALTER COLUMN id SET DEFAULT nextval('public.blocks_id_seq'::regclass);


--
-- Name: departments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments ALTER COLUMN id SET DEFAULT nextval('public.departaments_id_seq'::regclass);


--
-- Name: divisions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisions ALTER COLUMN id SET DEFAULT nextval('public.divisions_id_seq'::regclass);


--
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- Name: offices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offices ALTER COLUMN id SET DEFAULT nextval('public.officies_id_seq'::regclass);


--
-- Name: positions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions ALTER COLUMN id SET DEFAULT nextval('public.positions_id_seq'::regclass);


--
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Data for Name: blocks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.blocks (id, title) FROM stdin;
1	Корпоративный блок
2	Розничный блок
\.


--
-- Data for Name: departments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.departments (id, title, block_id, office_id) FROM stdin;
1	Департамент информационных технологий	1	1
2	Департамент управления персоналом	1	2
3	Финансовый департамент	1	3
4	Департамент продаж	2	1
5	Департамент клиентской поддержки	2	2
6	Маркетинговый департамент	2	3
\.


--
-- Data for Name: divisions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.divisions (id, title, department_id) FROM stdin;
1	Отдел разработки	1
2	Отдел кадров	2
3	Финансовый отдел	3
4	Отдел продаж	4
5	Отдел поддержки	5
6	Маркетинговый отдел	6
\.


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employees (id, first_name, middle_name, last_name, created_at, division_id, role_id, position_id, project_id, head_id) FROM stdin;
301	Анна	Сергеевна	Морозова	2024-12-01 10:50:31.890211	1	1	1	1	\N
302	Георгий	Алексеевич	Новиков	2024-12-01 10:50:31.890211	2	1	1	2	\N
303	Екатерина	Ивановна	Фомина	2024-12-01 10:50:31.890211	3	1	1	3	\N
304	Владислав	Юрьевич	Тимофеев	2024-12-01 10:50:31.890211	4	1	1	1	\N
305	Марина	Николаевна	Шестакова	2024-12-01 10:50:31.890211	5	1	1	2	\N
306	Денис	Аркадьевич	Руднев	2024-12-01 10:50:31.890211	6	1	1	3	\N
7	Иван	Петрович	Иванов	2024-12-01 00:13:23.002207	1	1	1	1	\N
8	Мария	Александровна	Смирнова	2024-12-01 00:13:23.002207	2	2	2	2	7
9	Алексей	Дмитриевич	Кузнецов	2024-12-01 00:13:23.002207	3	3	3	3	7
10	Ольга	Игоревна	Лебедева	2024-12-01 00:13:23.002207	4	4	4	1	8
11	Петр	Анатольевич	Сидоров	2024-12-01 00:13:23.002207	5	3	2	2	8
12	Елена	Викторовна	Григорьева	2024-12-01 00:13:23.002207	6	2	5	3	8
307	Иван	Михайлович	Кузьмин	2024-12-01 10:50:31.890211	1	2	2	1	301
308	Татьяна	Сергеевна	Зайцева	2024-12-01 10:50:31.890211	2	2	2	2	302
309	Максим	Денисович	Головин	2024-12-01 10:50:31.890211	3	2	2	3	303
310	Юлия	Алексеевна	Павлова	2024-12-01 10:50:31.890211	4	2	2	1	304
311	Виктор	Павлович	Беляев	2024-12-01 10:50:31.890211	5	2	2	2	305
312	Дарья	Николаевна	Григорьева	2024-12-01 10:50:31.890211	6	2	2	3	306
313	Евгений	Иванович	Семенов	2024-12-01 10:50:31.890211	1	2	3	1	307
314	Александра	Дмитриевна	Мартынова	2024-12-01 10:50:31.890211	2	2	3	2	308
315	Сергей	Валерьевич	Соловьев	2024-12-01 10:50:31.890211	3	2	3	3	309
316	Кристина	Олеговна	Полякова	2024-12-01 10:50:31.890211	4	2	3	1	310
317	Галина	Юрьевна	Лазарева	2024-12-01 10:50:31.890211	5	2	3	2	311
318	Павел	Станиславович	Гаврилов	2024-12-01 10:50:31.890211	6	2	3	3	312
319	Анастасия	Викторовна	Миронова	2024-12-01 10:50:31.890211	1	3	4	1	313
320	Илья	Дмитриевич	Тихонов	2024-12-01 10:50:31.890211	2	3	4	2	314
321	Дмитрий	Николаевич	Князев	2024-12-01 10:50:31.890211	3	3	4	3	315
322	Елена	Сергеевна	Ильина	2024-12-01 10:50:31.890211	4	3	4	1	316
323	Роман	Петрович	Фролов	2024-12-01 10:50:31.890211	5	3	4	2	317
324	Оксана	Игоревна	Никитина	2024-12-01 10:50:31.890211	6	3	4	3	318
325	Кирилл	Алексеевич	Попов	2024-12-01 10:50:31.890211	1	4	6	1	319
326	Маргарита	Андреевна	Орлова	2024-12-01 10:50:31.890211	2	4	6	2	320
327	Владимир	Евгеньевич	Мартынов	2024-12-01 10:50:31.890211	3	4	6	3	321
328	Полина	Геннадьевна	Герасимова	2024-12-01 10:50:31.890211	4	4	6	1	322
329	Антон	Игоревич	Трофимов	2024-12-01 10:50:31.890211	5	4	6	2	323
330	Надежда	Данииловна	Кудрявцева	2024-12-01 10:50:31.890211	6	4	6	3	324
159	Иван	Петрович	Иванов	2024-12-01 00:36:40.780307	1	1	1	1	\N
160	Мария	Александровна	Смирнова	2024-12-01 00:36:40.780307	1	1	2	2	\N
161	Алексей	Иванович	Кузнецов	2024-12-01 00:36:40.780307	1	2	3	1	7
162	Ольга	Николаевна	Лебедева	2024-12-01 00:36:40.780307	2	2	4	2	8
163	Петр	Васильевич	Сидоров	2024-12-01 00:36:40.780307	1	3	5	1	9
164	Елена	Сергеевна	Григорьева	2024-12-01 00:36:40.780307	2	3	6	2	9
165	Светлана	Петровна	Куликова	2024-12-01 00:36:40.780307	2	3	1	1	10
166	Татьяна	Юрьевна	Васильева	2024-12-01 00:36:40.780307	1	3	2	1	11
167	Игорь	Александрович	Романов	2024-12-01 00:36:40.780307	2	3	3	2	12
168	Анастасия	Дмитриевна	Королева	2024-12-01 00:36:40.780307	2	3	4	1	\N
169	Максим	Игоревич	Кузнецов	2024-12-01 00:36:40.780307	1	3	5	1	\N
170	Андрей	Петрович	Белый	2024-12-01 00:36:40.780307	2	3	6	2	\N
171	Дарина	Сергеевна	Смирнова	2024-12-01 00:36:40.780307	2	3	1	1	\N
172	Владимир	Викторович	Соколов	2024-12-01 00:36:40.780307	1	4	2	1	7
173	Марина	Игоревна	Петрова	2024-12-01 00:36:40.780307	1	4	3	2	8
174	Роман	Михайлович	Дмитриев	2024-12-01 00:36:40.780307	2	4	4	1	9
175	Юлия	Федоровна	Лапшина	2024-12-01 00:36:40.780307	2	4	5	2	10
176	Евгений	Артурович	Смирнов	2024-12-01 00:36:40.780307	1	4	6	1	11
177	Ирина	Юрьевна	Коваленко	2024-12-01 00:36:40.780307	2	4	1	2	12
178	Александра	Анатольевна	Шевченко	2024-12-01 00:36:40.780307	2	4	2	1	\N
179	Иван	Петрович	Сидоров	2024-12-01 08:20:01.996555	1	2	1	1	\N
180	Елена	Ивановна	Смирнова	2024-12-01 08:20:01.996555	2	2	1	2	\N
181	Ольга	Александровна	Попова	2024-12-01 08:20:01.996555	3	2	1	3	\N
182	Сергей	Васильевич	Кузнецов	2024-12-01 08:20:01.996555	4	2	1	1	\N
183	Наталья	Владимировна	Морозова	2024-12-01 08:20:01.996555	5	2	1	2	\N
184	Дмитрий	Алексеевич	Лебедев	2024-12-01 08:20:01.996555	6	2	1	3	\N
201	Роман	Андреевич	Жуков	2024-12-01 08:20:02.013699	1	3	4	1	\N
202	Нина	Станиславовна	Щербакова	2024-12-01 08:20:02.013699	2	3	2	2	\N
203	Оксана	Игоревна	Павлова	2024-12-01 08:20:02.013699	3	3	5	3	\N
204	Леонид	Валерьевич	Кириллов	2024-12-01 08:20:02.013699	4	3	3	1	\N
205	Эльвира	Сергеевна	Агапова	2024-12-01 08:20:02.013699	5	3	4	2	\N
206	Артур	Георгиевич	Селезнёв	2024-12-01 08:20:02.013699	6	3	2	3	\N
\.


--
-- Data for Name: offices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.offices (id, title, city) FROM stdin;
1	Головной офис	Москва
2	Региональный офис	Санкт-Петербург
3	Филиал	Новосибирск
\.


--
-- Data for Name: positions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.positions (id, title) FROM stdin;
1	Менеджер
2	Аналитик
3	Специалист по продажам
4	Технический специалист
5	Бухгалтер
6	Стажёр
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, title) FROM stdin;
1	Проект А
2	Проект Б
3	Проект В
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, title) FROM stdin;
1	Администратор
2	Менеджер
3	Сотрудник
4	Гостевой доступ
\.


--
-- Name: blocks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.blocks_id_seq', 2, true);


--
-- Name: departaments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.departaments_id_seq', 6, true);


--
-- Name: divisions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.divisions_id_seq', 6, true);


--
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_id_seq', 206, true);


--
-- Name: officies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.officies_id_seq', 3, true);


--
-- Name: positions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.positions_id_seq', 6, true);


--
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.projects_id_seq', 3, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 4, true);


--
-- Name: blocks blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.blocks
    ADD CONSTRAINT blocks_pkey PRIMARY KEY (id);


--
-- Name: departments departaments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departaments_pkey PRIMARY KEY (id);


--
-- Name: divisions divisions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisions
    ADD CONSTRAINT divisions_pkey PRIMARY KEY (id);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);


--
-- Name: offices officies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offices
    ADD CONSTRAINT officies_pkey PRIMARY KEY (id);


--
-- Name: positions positions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: projects projects_title_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_title_key UNIQUE (title);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: roles roles_title_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_title_key UNIQUE (title);


--
-- Name: departments departaments_block_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departaments_block_id_fkey FOREIGN KEY (block_id) REFERENCES public.blocks(id);


--
-- Name: departments departaments_office_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departaments_office_id_fkey FOREIGN KEY (office_id) REFERENCES public.offices(id);


--
-- Name: divisions divisions_departament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisions
    ADD CONSTRAINT divisions_departament_id_fkey FOREIGN KEY (department_id) REFERENCES public.departments(id);


--
-- Name: employees employees_division_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_division_id_fkey FOREIGN KEY (division_id) REFERENCES public.divisions(id);


--
-- Name: employees employees_position_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_position_id_fkey FOREIGN KEY (position_id) REFERENCES public.positions(id);


--
-- Name: employees employees_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: employees employees_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: divisions fk_departament; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.divisions
    ADD CONSTRAINT fk_departament FOREIGN KEY (department_id) REFERENCES public.departments(id);


--
-- Name: employees fk_head; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_head FOREIGN KEY (head_id) REFERENCES public.employees(id);


--
-- Name: employees fk_position; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_position FOREIGN KEY (position_id) REFERENCES public.positions(id);


--
-- Name: employees fk_project; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: employees fk_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

