INSERT INTO accounts (
  id,
  name,
  type,
  balance
)
VALUES
(
  'e0698db1-ca65-4903-94df-a4917f795562'::uuid,
  'Escrow',
  'ESCROW'::type_account_type,
  DEFAULT
),
(
  '93c8ee71-bda7-438c-b261-d9e97f9c5286'::uuid,
  'Cash',
  'CASH'::type_account_type,
  DEFAULT
),
(
  '991842fe-2e97-4481-a560-8d985a82ae74'::uuid,
  'Investor A',
  'INVESTOR'::type_account_type,
  10000000
),
(
  'feac2610-27df-4665-afae-0f536ed06ab5'::uuid,
  'Investor B',
  'INVESTOR'::type_account_type,
  5000000
),
(
  'c5f76419-eb27-4255-86ba-afbbe271114b'::uuid,
  'Investor C',
  'INVESTOR'::type_account_type,
  200000
),
(
  '5af74869-9b16-4ddd-9f0d-4a1df2b980eb'::uuid,
  'Company A',
  'ISSUER'::type_account_type,
  DEFAULT
),
(
  'be95a593-c12a-495d-ae12-6dc45d8d9970'::uuid,
  'Company B',
  'ISSUER'::type_account_type,
  DEFAULT
),
(
  'de16507a-61b3-43f6-b977-4312f52ece1b'::uuid,
  'Company C',
  'ISSUER'::type_account_type,
  DEFAULT
);


INSERT INTO invoices (
    id,
    issuer_account_id,
    reference,
    denom,
    amount,
    asking
)
VALUES
(
    'acb51e7b-2cef-4081-93ad-6b3a97c68b8a'::uuid,
    '5af74869-9b16-4ddd-9f0d-4a1df2b980eb'::uuid,
    '2019/001',
    'EUR',
    150000,
    100000
),
(
    'af80d0ea-78b9-45b1-a7b0-d1ddd0fbd6fe'::uuid,
    'be95a593-c12a-495d-ae12-6dc45d8d9970'::uuid,
    '2022-0002',
    'EUR',
    300000,
    200000
),
(
    'ceeaece4-ca5c-4d31-9fd6-90a90854fed9'::uuid,
    'de16507a-61b3-43f6-b977-4312f52ece1b'::uuid,
    'INV-3-2022',
    'EUR',
    450000,
    300000
);

-- bid(_invoice_id uuid, _bidder_account_id uuid, _offer bigint)
SELECT bid('ceeaece4-ca5c-4d31-9fd6-90a90854fed9', '991842fe-2e97-4481-a560-8d985a82ae74', 160000);
SELECT bid('ceeaece4-ca5c-4d31-9fd6-90a90854fed9', 'feac2610-27df-4665-afae-0f536ed06ab5', 420000);
SELECT bid('af80d0ea-78b9-45b1-a7b0-d1ddd0fbd6fe', 'c5f76419-eb27-4255-86ba-afbbe271114b', 195000);


-- adjudicate(_invoice_id uuid, _bidder_account_id uuid)
SELECT adjudicate('ceeaece4-ca5c-4d31-9fd6-90a90854fed9', 'feac2610-27df-4665-afae-0f536ed06ab5');


