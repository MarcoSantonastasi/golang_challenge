CREATE OR REPLACE FUNCTION bid (
  INOUT _bidder_account uuid,
  INOUT _invoice uuid,
  INOUT _offer bigint
) 
AS $$
DECLARE
  _esc uuid;
  _tx uuid;

BEGIN
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;

SELECT id FROM accounts
WHERE
  type='ESCROW'::type_account_type
FETCH FIRST ROW ONLY
INTO _esc;

INSERT INTO transactions ( invoice, is_bid )
VALUES (_invoice, true)
RETURNING id INTO _tx;

EXCEPTION WHEN unique_violation THEN
SELECT id FROM transactions
WHERE
  invoice = _invoice AND
  is_active = true
INTO _tx;

INSERT INTO ledger( transaction, credit, debit, amount)
VALUES (_tx, _esc, _bidder_account, _offer);

RETURN;
END;
$$ LANGUAGE plpgsql;
