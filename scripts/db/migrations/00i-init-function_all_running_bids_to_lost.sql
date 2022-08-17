CREATE OR REPLACE FUNCTION all_running_bids_to_lost ( _invoice_id uuid)
RETURNS SETOF bids
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

RETURN QUERY
SELECT
    ub.*
  FROM
    bids AS ub
  WHERE
    ub.state = 'LOST'::type_bid_state
    AND ub.invoice_id = _invoice_id;

END;
$$ LANGUAGE plpgsql;