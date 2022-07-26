DROP TABLE IF EXISTS invoices;

CREATE TABLE invoices
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    issuer uuid NOT NULL,
    denom character varying(3) NOT NULL DEFAULT 'EUR',
    amount bigint,
    asking bigint NOT NULL,
    is_biddable boolean NOT NULL DEFAULT true,
    CONSTRAINT invoices_pkey PRIMARY KEY (id),
    CONSTRAINT invoices_issuer_fkey FOREIGN KEY (issuer)
      REFERENCES accounts(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION
);

COMMENT ON TABLE invoices
    IS 'Invoices are the document that is traded. Numeric fields representing currency are always in cents of the denomination (i.e. a value of 100 = 1 currency unit)';
