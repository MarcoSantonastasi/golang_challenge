# arex_challenge

## Requirements

### General
In Arex we deal with bank accounts, invoices, transfers, etc... and we want to see how you deal with one of the classic tasks that our users do.
The goal of the tech challenge is to design a system supporting the requirements defined in the next sections, satisfying the following:
- [ ] For each invoice accepted into the market there will be an asking price
- [ ] Each investor will place bids to purchase an invoice
- [ ] Once the invoice has been purchased, the trade is considered locked until it’s approved, in the meantime it can be reversed or commited
- [ ] Reserve money until specific transactions have been approved
- [ ] Transactional consistency should be enforced when registering the cash flows for each use case
- [ ] As a bonus, be mindful of the concurrency when handling invoices

### Invoice Handling
In order to support invoice financing the system should be able to track the balances of the involved parties:
- [ ] Issuer - A company that has an invoice to be financed.
- [ ] Investors - The buyers of the invoice.
For that, the account’s structure of each party should be defined, considering how the ledger will be implemented. Besides modelling parties and the ledger, any required data should be seeded when necessary.

In terms of users interactions, the system should be able to handle and register invoices on the ledger through:
- [ ] An endpoint to sell an invoice
- [ ] An endpoint to retrieve an invoice, including the status: if it has been bought or not. In case it has been bought, by which investor (Example: An issuer issuer_A wants to finance an invoice invoice_1 with value €1000, recording in the ledger that issuer_A has a €1000 invoice invoice_1 to be financed)

### Ledger
Once an invoice has been stored, the bidding process from existing investors with available funds starts, considering the following aspects:
- [ ] An investor cannot place a bid if doesn’t have available funds
- [ ] After a bid has been placed, the available balance of an investor must be reduced

- [ ] For each investor’s bid received, the matching algorithm should:
    - [ ] Process bids by First In First (out) Served order until the invoice has been sold.
    - [ ] If a bid placed fills 100% of the invoice (meaning the price offered is greater or equal to the amount asked by the issuer) there is a match and the invoice is financed
    - [ ] After a match, its results should be recorded and the investor’s unplaced bids restored. At this point the financing is considered locked and the following should be implemented:
     - [ ] A persisted set of investors with their available balance.
     - [ ] An endpoint to list all the investors and their balances.
     - [ ] A Matching algorithm fulfilling the business rules described above

Example: An issuer issuer_A wants to finance invoice invoice_1. The total of the invoice is €1000. The system executes a series of bids for that invoice in the following order:
1. investor_A places a bid of size €500.
2. investor_B places a bid of size €300.
3. investor_C places a bid of size €1200.
4. Once bids 1 and 2 are done, the matching algorithm verifies they are smaller than the amount of the invoice, so they are discarded. When bid 3 is received the algorithm verifies that the total of the bid, €1200, is more than enough to cover the invoice. At this point, there is a match, so the invoice is considered financed and it’s now locked

To keep track of available and reserved balance for different entities, the ledger should:
- [ ] Reserve balance when placing bids.
- [ ] Release balance when a bid is rejected.
- [ ] Revert transactions when financing is reversed.
- [ ] Register cash flows when financing is approved.

Example: After the last bid has been placed and before the matching algorithm runs, the ledger should reflect the following:
- party_A has a €1000 invoice invoice_1 that should be financed.
- investor_C has €1200 reserved for the purchase of invoice_1.
After the matching algorithm executes, it releases to the investor the part of the reserved balance that wasn’t used, then investor_C available balance increases by €200.
At this point the financing is locked and can either be committed or reversed. If reversed, all the ledger entries generated for this specific financing should be rolled-back, restoring investor’s balance and mark the target invoice as available for financing.
On the other hand, if the financing is confirmed, the cash flows should accurately represent the transaction, meaning for this particular example that party_A owes €1000 for invoice_1 to investor_C, and the amount reserved by investor_C, €1000, should be effectively subtracted from their balance and paid to party_A for the purchase.


## Constraints
- [ ] Use Golang and gRPC to implement the functionalities and related endpoints
- [ ] High level of coverage by meaningful tests is expected, including end-to-end scenarios
- [ ] Including linters would be ideal
- [ ] DB is expected to be PostgreSQL
- [ ] Docker Compose for starting the application
- [ ] Provide a collection to explore and query the endpoints in a well-known client, like bloomrpc, or you can craft one
- [ ] Push your implementation to a private GitHub repo and share it with us
- [ ] A descriptive README is expected in order to help us to execute and test the implementation
- [ ] A good commit/PR history that shows how the challenge was approached will be ideal
