DROP TABLE IF EXISTS invoices CASCADE;

DROP TYPE IF EXISTS type_invoice_state;

CREATE TYPE type_invoice_state AS ENUM ('FORSALE', 'ADJUDICATED', 'FINANCED');

CREATE TABLE invoices
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    issuer_account_id uuid NOT NULL,
    reference varchar NOT NULL,
    denom character varying(3) NOT NULL DEFAULT 'EUR',
    amount bigint,
    asking bigint NOT NULL,
    state type_invoice_state NOT NULL DEFAULT 'FORSALE'::type_invoice_state,
    CONSTRAINT invoices_pkey PRIMARY KEY (id),
    CONSTRAINT invoices_issuer_account_id_fkey FOREIGN KEY (issuer_account_id)
      REFERENCES accounts(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION,
    CONSTRAINT invoices_issuer_account_id_reference_unique UNIQUE (issuer_account_id, reference)
);

COMMENT ON TABLE invoices IS
  'Invoices are the documents that are auctioned, the asking field representing the minimum value the issuer is willing to acept. Numeric fields representing currency are always in cents of the denomination (i.e. a value of 100 = 1 currency unit)';
