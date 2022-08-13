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