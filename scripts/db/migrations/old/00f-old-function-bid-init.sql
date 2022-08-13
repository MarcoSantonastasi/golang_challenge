CREATE OR REPLACE FUNCTION bid (
  INOUT _invoice_id uuid,
  INOUT _bidder_account_id uuid,
  INOUT _offer bigint
)
AS $$
DECLARE
  _escrow_account_id uuid;
  _bid_id bigint;

BEGIN

SELECT id FROM accounts
WHERE
  type='ESCROW'::type_account_type
FETCH FIRST ROW ONLY
INTO _escrow_account_id;

LOOP
  SELECT id FROM bids
  WHERE
    invoice_id = _invoice_id AND
    state = 'RUNNING'::type_bid_state
  INTO _bid_id;

  EXIT WHEN FOUND;

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
  ON CONFLICT DO NOTHING
  RETURNING id INTO _bid_id;
  
  EXIT WHEN FOUND;
END LOOP;

INSERT INTO transactions( bid_id, credit_account_id, debit_account_id, amount)
VALUES (_bid_id, _escrow_account_id, _bidder_account_id, _offer);

RETURN;
END;
$$ LANGUAGE plpgsql;
