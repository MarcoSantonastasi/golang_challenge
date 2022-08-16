
CREATE OR REPLACE VIEW bids_with_invoice
AS
SELECT
    b.*,
    i.created_at AS invoice_created_at,
    i.issuer_account_id AS invoice_issuer_account_id,
    i.reference AS invoice_reference,
    i.denom AS invoice_denom,
    i.amount AS invoice_amount,
    i.asking AS invoice_asking,
    i.state AS invoice_state 
FROM bids AS b
LEFT JOIN invoices AS i
ON i.id = b.invoice_id;