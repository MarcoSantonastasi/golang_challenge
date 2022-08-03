CREATE OR REPLACE VIEW bids (
invoice, issuer, bidder_account)
AS (
  with escrow as (select id from accounts where type = 'ESCROW' fetch first row only)
  SELECT
    tx.invoice as invoice,
    inv.issuer as issuer,
    lg.debit as bidder_account
  FROM transactions as tx
  JOIN invoices as inv ON tx.invoice = inv.id
  JOIN ledger as lg ON lg.transaction = tx.id
  -- add JOIN LATERAL (select ...) as debit
);
