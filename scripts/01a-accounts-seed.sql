INSERT INTO accounts (
  id,
  name,
  type,
  balance
)
VALUES
(
  'e0698db1-ca65-4903-94df-a4917f795562'::uuid,
  'ESCROW',
  'ESCROW'::type_account_type,
  DEFAULT
),
(
  '93c8ee71-bda7-438c-b261-d9e97f9c5286'::uuid,
  'CASH',
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
  DEFAULT
),
(
  'c5f76419-eb27-4255-86ba-afbbe271114b'::uuid,
  'Investor C',
  'INVESTOR'::type_account_type,
  DEFAULT
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
