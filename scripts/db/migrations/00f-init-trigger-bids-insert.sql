CREATE OR REPLACE FUNCTION check_invoice_forsale_on_bids_insert()
RETURNS TRIGGER AS $$

DECLARE
  _invoice_state type_invoice_state;

BEGIN

  SELECT state FROM invoices
  WHERE id = new.invoice_id
  INTO _invoice_state;
  
  IF _invoice_state <> 'FORSALE'::type_invoice_state THEN
    RAISE EXCEPTION 'cannot bid on invoice that is not for sale';
  END IF;

  RETURN new;

END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER check_invoice_forsale_on_bids_insert
BEFORE INSERT ON bids
  FOR EACH ROW EXECUTE FUNCTION check_invoice_forsale_on_bids_insert();
