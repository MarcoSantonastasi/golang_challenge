CREATE OR REPLACE FUNCTION update_account_balance_on_ledger_insert() 
RETURNS TRIGGER AS $$
BEGIN
  UPDATE accounts
  SET balance = balance + new.amount
  WHERE id = new.credit;
  UPDATE accounts
  SET balance = balance - new.amount
  WHERE id = new.debit;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_account_balance_on_ledger_insert
AFTER INSERT ON ledger
  FOR EACH ROW EXECUTE FUNCTION update_account_balance_on_ledger_insert();


CREATE OR REPLACE FUNCTION check_sufficient_balance_on_ledger_insert()
RETURNS TRIGGER AS $$

DECLARE
  _balance bigint;

BEGIN
  SELECT balance FROM accounts
  WHERE id = new.debit
  INTO _balance;
  
  IF new.amount > _balance THEN
    RAISE EXCEPTION 'insufficient funds';
   END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER check_sufficient_balance_on_ledger_insert
BEFORE INSERT ON ledger
  FOR EACH ROW EXECUTE FUNCTION check_sufficient_balance_on_ledger_insert();
