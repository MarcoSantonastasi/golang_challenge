CREATE OR REPLACE FUNCTION bid (
  INOUT _bidder_account uuid,
  INOUT _invoice uuid,
  INOUT _offer bigint
) 
AS $$
DECLARE
  _esc uuid;
  _tx bigint;

BEGIN

SELECT id FROM accounts
WHERE
  type='ESCROW'::type_account_type
FETCH FIRST ROW ONLY
INTO _esc;

LOOP
  SELECT id FROM transactions
  WHERE
    invoice = _invoice AND
    is_active = true
  INTO _tx;

  EXIT WHEN FOUND;

  INSERT INTO transactions (invoice)
  VALUES (_invoice)
  RETURNING id INTO _tx;

  EXIT WHEN FOUND;
END LOOP;

INSERT INTO ledger( transaction, credit, debit, amount)
VALUES (_tx, _esc, _bidder_account, _offer);

RETURN;
END;
$$ LANGUAGE plpgsql;
