DROP TABLE IF EXISTS accounts CASCADE;

DROP TYPE IF EXISTS type_account_type CASCADE;

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

DROP TYPE IF EXISTS type_invoice_state CASCADE;

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

DROP TYPE IF EXISTS type_bid_state CASCADE;

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



DROP FUNCTION IF EXISTS check_sufficient_balance_on_transactions_insert;

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

DROP TRIGGER IF EXISTS check_sufficient_balance_on_transactions_insert;

CREATE OR REPLACE TRIGGER check_sufficient_balance_on_transactions_insert
BEFORE INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION check_sufficient_balance_on_transactions_insert();


DROP FUNCTION IF EXISTS update_account_balance_on_transactions_insert;

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

DROP TRIGGER IF EXISTS update_account_balance_on_transactions_insert;

CREATE OR REPLACE TRIGGER update_account_balance_on_transactions_insert
AFTER INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION update_account_balance_on_transactions_insert();



DROP FUNCTION IF EXISTS check_invoice_forsale_on_bids_insert;

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

DROP TRIGGER IF EXISTS check_invoice_forsale_on_bids_insert;

CREATE OR REPLACE TRIGGER check_invoice_forsale_on_bids_insert
BEFORE INSERT ON bids
  FOR EACH ROW EXECUTE FUNCTION check_invoice_forsale_on_bids_insert();



DROP FUNCTION IF EXISTS bid;

CREATE OR REPLACE FUNCTION bid (
  OUT new_bid_id bigint,
  INOUT invoice_id uuid,
  INOUT bidder_account_id uuid,
  INOUT new_bid_offer bigint,
  OUT new_bid_state type_bid_state
)
AS $$

DECLARE
	_escrow_account_id uuid;
	_invoice_state type_invoice_state;

BEGIN
	SELECT id FROM accounts AS a
	WHERE
		a.type = 'ESCROW'::type_account_type FETCH FIRST ROW ONLY INTO _escrow_account_id;

	INSERT INTO bids AS b (invoice_id, bidder_account_id, offer)
		VALUES(invoice_id, bidder_account_id, new_bid_offer)
	RETURNING
		b.id, b.invoice_id, b.bidder_account_id, b.offer, b.state
    INTO
        new_bid_id, invoice_id, bidder_account_id, new_bid_offer, new_bid_state;
	
  INSERT INTO transactions (bid_id, credit_account_id, debit_account_id, amount)
		VALUES(new_bid_id, _escrow_account_id, bidder_account_id, new_bid_offer);
	
  RETURN;
END;
$$
LANGUAGE plpgsql;




DROP FUNCTION IF EXISTS adjudicate_bid;

CREATE OR REPLACE FUNCTION adjudicate_bid (INOUT bid_id bigint, OUT paid_amount bigint)
AS $$

DECLARE

_escrow_account_id uuid;
_cash_account_id uuid;
_bidder_account_id uuid;
_bid_offer bigint;
_invoice_id uuid;
_invoice_issuer_account_id uuid;
_invoice_asking bigint;

BEGIN

SELECT id FROM accounts AS e
WHERE e.type = 'ESCROW'::type_account_type
FETCH FIRST ROW ONLY INTO _escrow_account_id ;

SELECT id FROM accounts AS c
WHERE c.type = 'CASH'::type_account_type
FETCH FIRST ROW ONLY INTO _cash_account_id;

SELECT
    b.bidder_account_id,
    b.offer,
    b.invoice_id,
    i.issuer_account_id,
    i.asking
  FROM
    bids AS b
    LEFT JOIN invoices AS i
    ON i.id = b.invoice_id
  WHERE
    b.id = bid_id
    AND b.state = 'RUNNING'::type_bid_state
    AND i.state = 'FORSALE'::type_invoice_state
  FETCH FIRST ROW ONLY
  INTO
  	_bidder_account_id,
  	_bid_offer,
  	_invoice_id,
  	_invoice_issuer_account_id,
    _invoice_asking;

INSERT INTO transactions (bid_id, credit_account_id, debit_account_id, amount)
	VALUES
		-- Pay the issuer 90% the asking amount on the invoice debiting the escrow account
		(bid_id, _invoice_issuer_account_id, _escrow_account_id, (_invoice_asking * 0.90)),
		-- Retain 10% as our fee debiting the escrow account
		(bid_id, _cash_account_id, _escrow_account_id, (_invoice_asking * 0.10)),
		-- Reinburse the difference between the asking and the offer to the bidder account debiting the escrow account
		(bid_id, _bidder_account_id, _escrow_account_id, (_bid_offer - _invoice_asking));
    

UPDATE
	bids AS ub
SET
	state = 'WON'::type_bid_state
WHERE
	ub.id = bid_id;

UPDATE
	invoices AS ui
SET
	state = 'ADJUDICATED'::type_invoice_state
WHERE
	ui.id = _invoice_id;

-- Can be different than asking when adjudicating for a lower price is implemented
paid_amount = _invoice_asking;

RETURN;
END;
$$
LANGUAGE plpgsql;




DROP FUNCTION IF EXISTS all_running_bids_to_lost;

CREATE OR REPLACE FUNCTION all_running_bids_to_lost ( _invoice_id uuid )
RETURNS TABLE (
	id bigint,
  created_at timestamp with time zone,
  invoice_id uuid,
  bidder_account_id uuid,
  offer bigint,
  state type_bid_state
)
AS $$
BEGIN
  WITH running_bids AS (
	  UPDATE
		  bids AS rb
	  SET
		  state = 'LOST'::type_bid_state
    WHERE
		  rb.state = 'RUNNING'::type_bid_state AND
		  rb.invoice_id = _invoice_id
	  RETURNING
		  rb.id AS bid_id,
		  rb.bidder_account_id AS credit_account_id,
		  rb.offer AS amount
  ),
  escrow_account AS (
	  SELECT
		  a.id
	  FROM
		  accounts as a
	  WHERE
		  a.type = 'ESCROW'::type_account_type FETCH FIRST ROW ONLY
  )
  INSERT INTO transactions AS t (bid_id, credit_account_id, debit_account_id, amount)
  SELECT
	  running_bids.bid_id, running_bids.credit_account_id, escrow_account.id AS debit_account_id, running_bids.amount
  FROM
	  running_bids,
	  escrow_account;

  SELECT
    ub.*
  FROM
    bids AS ub
  WHERE
    ub.state = 'LOST'::type_bid_state
    AND ub.invoice_id = _invoice_id;
RETRUN;
END;
$$ LANGUAGE plpgsql;





DROP VIEW IF EXISTS issuers;

CREATE OR REPLACE VIEW issuers AS SELECT * FROM accounts WHERE type = 'ISSUER'::type_account_type;


DROP VIEW IF EXISTS investors;

CREATE OR REPLACE VIEW investors AS SELECT * FROM accounts WHERE type = 'INVESTOR'::type_account_type;


DROP VIEW IF EXISTS bids_with_invoice;

CREATE OR REPLACE VIEW bids_with_invoice
AS
SELECT
    b.*,
    i.created_at AS invoice_created_at,
    i.issuer_account_id AS invoice_issuer_account_id,
    i.reference AS invoice_reference,
    i.denom AS invoice_denom,
    i.amount AS invoice_amount,
    i.asking AS invoice_asking,
    i.state AS invoice_state 
FROM bids AS b
LEFT JOIN invoices AS i
ON i.id = b.invoice_id;