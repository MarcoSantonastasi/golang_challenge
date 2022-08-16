CREATE OR REPLACE FUNCTION adjudicate_bid (INOUT _bid_id bigint, OUT _paid_amount bigint)
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

SELECT id FROM accounts
WHERE type = 'ESCROW'::type_account_type
FETCH FIRST ROW ONLY INTO _escrow_account_id ;

SELECT id FROM accounts
WHERE type = 'CASH'::type_account_type
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
    b.id = _bid_id
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
		(_bid_id, _invoice_issuer_account_id, _escrow_account_id, (_invoice_asking * 0.90)),
		-- Retain 10% as our fee debiting the escrow account
		(_bid_id, _cash_account_id, _escrow_account_id, (_invoice_asking * 0.10)),
		-- Reinburse the difference between the asking and the offer to the bidder account debiting the escrow account
		(_bid_id, _bidder_account_id, _escrow_account_id, (_bid_offer - _invoice_asking));
    

UPDATE
	bids
SET
	state = 'WON'::type_bid_state
WHERE
	id = _bid_id;

UPDATE
	invoices
SET
	state = 'ADJUDICATED'::type_invoice_state
WHERE
	id = _invoice_id;

-- Can be different than asking when adjudicating for a lower price is implemented
_paid_amount = _invoice_asking;

RETURN;
END;
$$
LANGUAGE plpgsql;
