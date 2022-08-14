CREATE OR REPLACE FUNCTION bid (
  OUT _bid_id bigint,
  INOUT _invoice_id uuid,
  INOUT _bidder_account_id uuid,
  INOUT _offer bigint,
  OUT _state type_bid_state
)
AS $$

DECLARE
	_escrow_account_id uuid;
	_invoice_state type_invoice_state;

BEGIN
	SELECT id FROM accounts
	WHERE
		type = 'ESCROW'::type_account_type FETCH FIRST ROW ONLY INTO _escrow_account_id;

	INSERT INTO bids (invoice_id, bidder_account_id, offer)
		VALUES(_invoice_id, _bidder_account_id, _offer)
	RETURNING
		id, invoice_id, bidder_account_id, offer, state
    INTO
        _bid_id, _invoice_id, _bidder_account_id, _offer, _state;
	
  INSERT INTO transactions (bid_id, credit_account_id, debit_account_id, amount)
		VALUES(_bid_id, _escrow_account_id, _bidder_account_id, _offer);
	
  RETURN;
END;
$$
LANGUAGE plpgsql;
