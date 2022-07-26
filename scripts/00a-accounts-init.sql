DROP TABLE IF EXISTS accounts CASCADE;

DROP TYPE IF EXISTS type_account_type;

CREATE TYPE type_account_type AS ENUM ('INVESTOR', 'ISSUER', 'ESCROW', 'CASH');

CREATE TABLE accounts
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT now(),
    name character varying COLLATE pg_catalog."default",
    type type_account_type NOT NULL,
    balance bigint NOT NULL DEFAULT 0::bigint,
    CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE accounts
	IS 'Accounts are an accounting object necessary to keep the ledger';
