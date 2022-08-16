CREATE OR REPLACE FUNCTION adjudicate_bid (
  IN _bid_id bigint,
  OUT _paid_amount bigint
)
AS $$
BEGIN

-- Pay the issuer 90% the asking amount on the invoice debiting the escrow account
-- Retain 10% as our fee debiting the escrow account
-- Reinburse the difference between the asking and the offer to the bidder account debiting the escrow account
INSERT INTO transactions(bid_id, credit_account_id, debit_account_id, amount)
VALUES
(_bid_id, invoice_issuer_account_id, escrow_account.id, (invoice_asking * 0.90)),
(_bid_id, cash_account.id, escrow_account.id, (invoice_asking * 0.10)),
(_bid_id, bid_bidder_account_id, escrow_account.id, (bid_offer - invoice_asking))
SELECT
    b.bidder_account_id AS bid_bidder_account_id,
    b.offer AS bid_offer,
    b.state AS bid_state,
    b.invoice_id,
    i.issuer_account_id AS invoice_issuer_account_id,
    i.asking AS invoice_asking,
    i.state AS invoice_state 
FROM bids AS b
LEFT JOIN invoices AS i
ON i.id = b.invoice_id
WHERE
  b.id = _bid_id AND
  b.state = 'RUNNING'::type_bid_state AND
  i.state = 'FORSALE'::type_invoice_state
FETCH FIRST ROW ONLY
RETURNING invoice_asking INTO _paid_amount;

WITH ub AS (
  UPDATE bids
  SET state = 'WON'::type_bid_state
  WHERE id = _bid_id
  FETCH FIRST ROW ONLY;
)
UPDATE invoices
SET state = 'ADJUDICATED'::type_invoice_state
WHERE id = ub.invoice_id;



















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


RETURN;
END;
$$ LANGUAGE plpgsql;
