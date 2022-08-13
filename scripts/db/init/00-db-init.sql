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

COMMENT ON TABLE accounts IS
    'Accounts are abstract identifiers of parties in a monetary transaction. They are of four types: INVESTOR, ISSUER, ESCROW, CASH';


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


DROP TABLE IF EXISTS bids CASCADE;

DROP TYPE IF EXISTS type_bid_state;

CREATE TYPE type_bid_state AS ENUM ('RUNNING', 'WIDTHDRAWN', 'WON', 'LOST');

CREATE TABLE bids
(
    id bigint GENERATED BY DEFAULT AS IDENTITY,
    created_at timestamp with time zone DEFAULT now(),
    invoice_id uuid NOT NULL,
    bidder_account_id uuid NOT NULL,
    offer bigint NOT NULL,
    state type_bid_state NOT NULL DEFAULT 'RUNNING'::type_bid_state,
    CONSTRAINT bids_pkey PRIMARY KEY (id),
    CONSTRAINT bids_invoice_id_fkey FOREIGN KEY (invoice_id)
      REFERENCES invoices(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION,
    CONSTRAINT bids_bidder_account_id_fkey FOREIGN KEY (bidder_account_id)
      REFERENCES accounts(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION
);

CREATE UNIQUE INDEX bids_invoice_id_bidder_account_id_state_running_unique
ON bids (invoice_id, bidder_account_id, state)
  WHERE state = 'RUNNING'::type_bid_state;

COMMENT ON TABLE bids IS
  'Bids represent Investor offers on invoices. They are the grouping abstraction for reconciling money transactions related to auctioning activity';


DROP TABLE IF EXISTS transactions CASCADE;

CREATE TABLE transactions
(
    id bigint GENERATED BY DEFAULT AS IDENTITY,
    created_at timestamp with time zone DEFAULT now(),
    bid_id bigint,
    credit_account_id uuid NOT NULL,
    debit_account_id uuid NOT NULL,
    amount bigint NOT NULL,
    CONSTRAINT transactions_pkey PRIMARY KEY (id),
    CONSTRAINT transactions_bid_id_fkey FOREIGN KEY (bid_id)
      REFERENCES bids(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION,
    CONSTRAINT transactions_credit_account_id_fkey FOREIGN KEY (credit_account_id)
      REFERENCES accounts(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION,
    CONSTRAINT transactions_debit_account_id_fkey FOREIGN KEY (debit_account_id)
      REFERENCES accounts(id) MATCH SIMPLE
      ON UPDATE NO ACTION
      ON DELETE NO ACTION
);

COMMENT ON TABLE transactions IS
  'Transactions record all tranfers of money. Only if the bid_id field is populated it means that the transaction is relative to a bid and reconciliation can be done within a bididng round.';


CREATE OR REPLACE FUNCTION check_sufficient_balance_on_transactions_insert()
RETURNS TRIGGER AS $$

DECLARE
  _debit_account_balance bigint;

BEGIN

  SELECT balance FROM accounts
  WHERE id = new.debit_account_id
  INTO _debit_account_balance;
  
  IF new.amount > _debit_account_balance THEN
    RAISE EXCEPTION 'insufficient funds';
  END IF;

  RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER check_sufficient_balance_on_transactions_insert
BEFORE INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION check_sufficient_balance_on_transactions_insert();


CREATE OR REPLACE FUNCTION update_account_balance_on_transactions_insert() 
RETURNS TRIGGER AS $$
BEGIN

  UPDATE accounts
  SET balance = balance + new.amount
  WHERE id = new.credit_account_id;

  UPDATE accounts
  SET balance = balance - new.amount
  WHERE id = new.debit_account_id;

RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_account_balance_on_transactions_insert
AFTER INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION update_account_balance_on_transactions_insert();


CREATE OR REPLACE FUNCTION check_invoice_forsale_on_bids_insert()
RETURNS TRIGGER AS $$

DECLARE
  _invoice_state type_invoice_state;

BEGIN

  SELECT state FROM invoices
  WHERE id = new.invoice_id
  INTO _invoice_state;
  
  IF _invoice_state <> 'FORSALE'::type_invoice_state THEN
    RAISE EXCEPTION 'cannot bid on invoice that is not for sale';
  END IF;

  RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER check_invoice_forsale_on_bids_insert
BEFORE INSERT ON bids
  FOR EACH ROW EXECUTE FUNCTION check_invoice_forsale_on_bids_insert();


CREATE OR REPLACE FUNCTION bid (
  INOUT _invoice_id uuid,
  INOUT _bidder_account_id uuid,
  INOUT _offer bigint
)
AS $$
DECLARE
  _escrow_account_id uuid;
  _invoice_state type_invoice_state;
  _bid_id bigint;

BEGIN
  SELECT id FROM accounts
  WHERE
    type = 'ESCROW'::type_account_type
  FETCH FIRST ROW ONLY
  INTO _escrow_account_id;

  INSERT INTO bids (
    invoice_id,
    bidder_account_id,
    offer
  )
  VALUES (
    _invoice_id,
    _bidder_account_id,
    _offer
  )
  RETURNING id INTO _bid_id;

  INSERT INTO transactions(bid_id, credit_account_id, debit_account_id, amount)
  VALUES (_bid_id, _escrow_account_id, _bidder_account_id, _offer);

  RETURN;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION adjudicate (
  INOUT _invoice_id uuid,
  INOUT _bidder_account_id uuid,
  OUT _amount bigint
)
AS $$
DECLARE
  _issuer_account_id uuid;
  _escrow_account_id uuid;
  _cash_account_id uuid;
  _invoice_asking bigint;
  _invoice_state type_invoice_state;
  _bid_id bigint;
  _bid_offer bigint;
  _bid_state type_bid_state;

BEGIN

SELECT issuer_account_id, asking, state FROM invoices
WHERE
  id = _invoice_id
FETCH FIRST ROW ONLY
INTO _issuer_account_id, _invoice_asking, _invoice_state;

IF _invoice_state <> 'FORSALE'::type_invoice_state THEN
  RAISE EXCEPTION 'can only adjudicate an invoice that is for sale';
END IF;

SELECT id, offer, state FROM bids
WHERE
  invoice_id = _invoice_id AND
  bidder_account_id = _bidder_account_id AND
  state = 'RUNNING'::type_bid_state
FETCH FIRST ROW ONLY
INTO _bid_id, _bid_offer, _bid_state;

IF _bid_id IS NULL THEN
  RAISE EXCEPTION 'cannot find a running bid for that invoice and bidder';
END IF;

SELECT id FROM accounts
WHERE
  type = 'ESCROW'::type_account_type
FETCH FIRST ROW ONLY
INTO _escrow_account_id;

SELECT id FROM accounts
WHERE
  type = 'CASH'::type_account_type
FETCH FIRST ROW ONLY
INTO _cash_account_id;

-- Pay the issuer 90% the asking amount on the invoice debiting the escrow account
-- Retain 10% as our fee debiting the escrow account
-- Reinburse the difference between the asking and the offer to the bidder account debiting the escrow account
INSERT INTO transactions(bid_id, credit_account_id, debit_account_id, amount)
VALUES
(_bid_id, _issuer_account_id, _escrow_account_id, (_invoice_asking * 0.90)),
(_bid_id, _cash_account_id, _escrow_account_id, (_invoice_asking * 0.10)),
(_bid_id, _bidder_account_id, _escrow_account_id, (_bid_offer - _invoice_asking));

-- For future cases where we could adjudicate below asking price
_amount := _invoice_asking;

UPDATE bids
SET state = 'WON'::type_bid_state
WHERE id = _bid_id;

UPDATE invoices
SET state = 'ADJUDICATED'::type_invoice_state
WHERE id = _invoice_id;

RETURN;
END;
$$ LANGUAGE plpgsql;
