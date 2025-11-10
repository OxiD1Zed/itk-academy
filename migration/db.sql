--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

-- Started on 2025-11-10 02:39:37

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
-- TOC entry 216 (class 1255 OID 17899)
-- Name: get_balance(uuid); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_balance(wallet_id uuid) RETURNS numeric
    LANGUAGE plpgsql
    AS $$
declare
res numeric;
begin
select balance into res from wallet where uuid = wallet_id;
return res;
end;
$$;


ALTER FUNCTION public.get_balance(wallet_id uuid) OWNER TO postgres;

--
-- TOC entry 228 (class 1255 OID 17884)
-- Name: update_wallet_balance(uuid, numeric); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_wallet_balance(IN wallet_id uuid, IN amount numeric)
    LANGUAGE plpgsql
    AS $$
declare
    current_balance NUMERIC;
begin
	if not (select exists (select 1 from wallet where uuid = wallet_id)) then raise exception sqlstate 'P0002' using message='wallet not found'; 
	end if;
	select balance into current_balance from wallet where uuid = wallet_id for update;
	if current_balance < abs(amount) and amount < 0 then raise exception sqlstate '22P02' using message='insufficient funds';
	end if;
	update wallet set balance = balance + amount where uuid = wallet_id;
	exception when others then
    	rollback;
        raise;
end;
$$;


ALTER PROCEDURE public.update_wallet_balance(IN wallet_id uuid, IN amount numeric) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 17886)
-- Name: strict; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.strict (
    uuid uuid,
    balance numeric
);


ALTER TABLE public.strict OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 17865)
-- Name: wallet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallet (
    uuid uuid DEFAULT gen_random_uuid() NOT NULL,
    balance numeric DEFAULT 0 NOT NULL
);


ALTER TABLE public.wallet OWNER TO postgres;

--
-- TOC entry 3326 (class 0 OID 17886)
-- Dependencies: 215
-- Data for Name: strict; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.strict (uuid, balance) FROM stdin;
\.


--
-- TOC entry 3325 (class 0 OID 17865)
-- Dependencies: 214
-- Data for Name: wallet; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wallet (uuid, balance) FROM stdin;
410a4e9b-45a5-49f4-86d0-dd3af8a1c430	6101
c309875b-4c17-4aab-a6ad-1cb3d0d42430	50
634b0194-cf1c-4457-bdec-b89a3f892468	150
3d100922-c312-417b-97a8-be4e476f57ae	200
b9d869ac-af29-442f-af8e-b0b710ec9c5a	100
\.


--
-- TOC entry 3181 (class 2606 OID 17873)
-- Name: wallet wallet_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallet
    ADD CONSTRAINT wallet_pkey PRIMARY KEY (uuid);


--
-- TOC entry 3182 (class 1259 OID 17885)
-- Name: wallet_uuid_btree; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX wallet_uuid_btree ON public.wallet USING btree (uuid);


-- Completed on 2025-11-10 02:39:37

--
-- PostgreSQL database dump complete
--

