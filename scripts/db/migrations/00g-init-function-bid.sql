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