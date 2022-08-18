# arex_challenge
see [./REQUIREMENTS.md](./REQUIREMENTS.md) for the challenge prompt.

## Intall and run 
You have to have Docker deamon and dockercompose alreeady installed on your machine, then you also have to make sure the 'make' executable is part of your PATH.

** Before running any code touch a '.env' in the root of the project and insert the following contants: **

POSTGRES_USER = postgres
POSTGRES_PASSWORD = postgres
POSTGRES_HOSTNAME = localhost
POSTGRES_PROD_DB = postgres
POSTGRES_TESTING_DB = testingdb
POSTGRES_STUB_DB = stubdb


You have a couple of pre-set commands that can help you explore the module:
- 'make demo': Will launch a server and a client connecting to a "production" database
- 'make e2etest': Will launch a smple e2e test suite against a 'testing' database
- 'make unittest': Will launch a simple unit testimg suite against a 'stub' database

As of now, all three dbs are seeded with the same migrations, but can be fully customised independently

## Description
The core problem of transaction consistency has been approached by having a single source of truth in the 'transactions' table.  The table records all transactions between Escrow, Cash and user accounts. There is the possibility to record any generic transaction, for example users geting money in and out of their accounts in case future development calls for exra features.  If a transaction is relative to a 'Bid', the Bid id is associated in the transaction as a foreign key so as to make full reconciliation upon bid closure plain easy. This is normally a best practice in db architecture for double-entry accounting.

The db API makes available a couple of methods thata should be exclusively used to interact with it in order to ensure data consistency. Direct manipulation of tables could be easily prevented by implementing RLS rules in a further development sprint.

The 'cmd/server' exposes a method 'NewBid()' that calls specialised methods on the db to handle the business requirement of checking and adjuducating fulfilling bids in order of "First come, first adjudicated"


## Work Log

### Iteration 0: Learing Go

At the end of Iteration 0 I have a working database with basic seeding, a working gRPC server with some very simple endpoints and a client that performs basic RPCs on those. Also I have a number of unit tests and an end-to-end test as a proof of concept.
Time has been spent mostly learning and applying knowledge at a slow pace to make sure I understood correctly how gRPC and the Go language work.

### Iteration 1: Basic DB APIs

At the end of Iteration 1 I have basic database APIs with a corresponding storage model in place, comprising all the entities and the triggers and the fucnitons to support the following basic functionality:
    1. new_invoice(issuer_account_id, ref, denom, amount, ask) => creates a new invoice in the corresponding table
    2. new_bid(bidder_id, invoice, amount) => creates a new bid veryfying sufficient funds and reserving funds.
    3. adjudicate_bid(bid_id) => Sets a bid to won and the relative invoice to adjudicated. It pays out the issuer and retains a fee of 10% for us.
    4. set_all_bids_to_lost(invoice_id) => Sets all bids related to an invoice to lost and refunds the money to investors by creating new transactions that reconcile to zero net money.
    
The DB is also seeded accordingly with just a few entries to make testing possible
