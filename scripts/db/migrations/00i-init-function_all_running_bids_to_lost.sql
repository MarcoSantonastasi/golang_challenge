CREATE OR REPLACE FUNCTION all_running_bids_to_lost (
  INOUT _invoice_id uuid,
  OUT _bid_id bigint
)
AS $$
BEGIN
  WITH running_bids AS (
	  UPDATE
		  bids
	  SET
		  state = 'LOST'::type_bid_state
    WHERE
		  state = 'RUNNING'::type_bid_state AND
		  invoice_id = _invoice_id
	  RETURNING
		  id AS bid_id,
		  bidder_account_id AS credit_account_id,
		  offer AS amount
  ),
  escrow_account AS (
	  SELECT
		  id
	  FROM
		  accounts
	  WHERE
		  TYPE = 'ESCROW'::type_account_type FETCH FIRST ROW ONLY
  )
  INSERT INTO transactions (bid_id, credit_account_id, debit_account_id, amount)
  SELECT
	  bid_id, credit_account_id, escrow_account.id AS debit_account_id, amount
  FROM
	  running_bids,
	  escrow_account;

RETURN;
END;
$$ LANGUAGE plpgsql;
