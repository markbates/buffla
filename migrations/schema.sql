--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.3
-- Dumped by pg_dump version 9.6.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: click_activity(uuid); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION click_activity(_link_id uuid) RETURNS TABLE(click_activity bigint, click_date date)
    LANGUAGE sql
    AS $_$
SELECT
COUNT(s.*),
d. DAY

FROM
(
  SELECT
  DAY :: DATE
  FROM
  generate_series(
    (
      SELECT
      created_at
      FROM
      links
      WHERE
      id = $1
    ) :: DATE,
    CURRENT_DATE :: DATE
    , INTERVAL '1 day'
  ) DAY
  ) d LEFT JOIN clicks s ON(
  s.created_at :: DATE = d.DAY AND s.link_id = $1
) GROUP BY d.DAY ORDER BY d.DAY
$_$;


ALTER FUNCTION public.click_activity(_link_id uuid) OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: clicks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE clicks (
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    id uuid NOT NULL,
    link_id uuid NOT NULL
);


ALTER TABLE clicks OWNER TO postgres;

--
-- Name: links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE links (
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    link character varying(255) NOT NULL,
    code character varying(255) NOT NULL
);


ALTER TABLE links OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE users (
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255),
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: clicks clicks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY clicks
    ADD CONSTRAINT clicks_pkey PRIMARY KEY (id);


--
-- Name: links links_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY links
    ADD CONSTRAINT links_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

