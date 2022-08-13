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