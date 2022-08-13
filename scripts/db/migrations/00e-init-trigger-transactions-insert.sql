CREATE OR REPLACE FUNCTION check_sufficient_balance_on_transactions_insert()
RETURNS TRIGGER AS $$

DECLARE
  _debit_account_balance bigint;

BEGIN

  SELECT balance FROM accounts
  WHERE id = new.debit_account_id
  INTO _debit_account_balance;
  
  IF new.amount > _debit_account_balance THEN
    RAISE EXCEPTION 'insufficient funds';
  END IF;

  RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER check_sufficient_balance_on_transactions_insert
BEFORE INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION check_sufficient_balance_on_transactions_insert();


CREATE OR REPLACE FUNCTION update_account_balance_on_transactions_insert() 
RETURNS TRIGGER AS $$
BEGIN

  UPDATE accounts
  SET balance = balance + new.amount
  WHERE id = new.credit_account_id;

  UPDATE accounts
  SET balance = balance - new.amount
  WHERE id = new.debit_account_id;

RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_account_balance_on_transactions_insert
AFTER INSERT ON transactions
  FOR EACH ROW EXECUTE FUNCTION update_account_balance_on_transactions_insert();
