CREATE TABLE IF NOT EXISTS issuers
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    name character varying COLLATE pg_catalog."default",
    CONSTRAINT issuers_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE issuers
    IS 'Issuer is company that has an invoice to be financed';


CREATE TABLE IF NOT EXISTS investors
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    name character varying COLLATE pg_catalog."default",
    CONSTRAINT investors_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE investors
    IS 'Investors are the buyers of the invoice';


CREATE TABLE IF NOT EXISTS invoices
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    denom character varying(3),
    amount integer,
    asking integer,
    CONSTRAINT invoices_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE invoices
    IS 'Invoices are the document that is traded. Numeric fields representing currency are in always in cents of the denomination (i.e. a value of 100 = 1 currency unit)';

