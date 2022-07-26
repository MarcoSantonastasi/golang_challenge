CREATE OR REPLACE FUNCTION revert_outstanding_bids (
  INOUT _invoice uuid,
  OUT _offer bigint
)
AS $$
DECLARE
  _tx bigint;
  _esc uuid;
  _account RECORD;
  _credit bigint;
  _debit bigint;

BEGIN

-- SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;

SELECT id FROM transactions
WHERE
  invoice = _invoice AND
  is_active = true
INTO _tx;

SELECT id FROM accounts
WHERE
  type='ESCROW'::type_account_type
INTO _esc;

FOR _account IN (
  SELECT debit AS _id FROM ledger
    WHERE transaction = _tx
  UNION
  SELECT credit AS _id FROM ledger
    WHERE transaction = _tx
)
LOOP

  SELECT SUM(amount) FROM ledger
  WHERE
    transaction=_transaction AND
    credit= _esc AND
    debit= _account
  INTO _credit;

  SELECT SUM(amount) FROM ledger
  WHERE
    transaction=_transaction AND
    credit= _account AND
    debit=  _esc
  INTO _debit;

  INSERT INTO ledger( transaction, credit, debit, amount)
    VALUES (_tx, _esc, _account, (coalesce(_credit,0::int) - coalesce(_debit,0::int)));
  
END LOOP;

UPDATE transactions SET is_active = false
  WHERE invoice = _invoice;

RETURN;
END; 
$$ LANGUAGE plpgsql;
