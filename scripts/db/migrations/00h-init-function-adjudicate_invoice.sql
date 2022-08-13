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
