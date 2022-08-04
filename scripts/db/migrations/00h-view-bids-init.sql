CREATE OR REPLACE VIEW bids (
  transaction,
  invoice,
  issuer,
  bidder_account,
  amount
)
AS (
  SELECT
    tx.id as transaction,
    tx.invoice as invoice,
    inv.issuer as issuer,
    bid_up.debit as bidder_account,
    coalesce(bid_up.amount, 0::bigint) - coalesce(bid_down.amount, 0::bigint) as amount
  FROM transactions as tx
  JOIN invoices as inv ON tx.invoice = inv.id
  JOIN (SELECT id FROM accounts WHERE type = 'ESCROW'::type_account_type FETCH FIRST ROW ONLY) as escrow ON true
  LEFT JOIN LATERAL (
    SELECT
      transaction,
      debit,
      sum(amount) as amount
    FROM ledger l
    WHERE l.credit = escrow.id
    GROUP BY transaction, debit
    ) AS bid_up ON bid_up.transaction = tx.id
  LEFT JOIN LATERAL (
    SELECT
      transaction,
      credit,
      sum(amount)as amount
    FROM ledger l
    WHERE l.debit = escrow.id
    GROUP BY transaction, credit
    ) AS bid_down ON bid_down.transaction = tx.id
  WHERE tx.is_active = true
);