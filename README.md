![e2e tsting CI](https://github.com/MarcoSantonastasi/arex_challenge/actions/workflows/e2etests.yml/badge.svg)

# golang_challenge
see [./REQUIREMENTS.md](./REQUIREMENTS.md) for the challenge prompt.


## Intall, run and demo 
You have to have Docker deamon and dockercompose alreeady installed on your machine, then you also have to make sure the `make` executable is part of your PATH.

***Before running any code, touch a `.env` in the root of the project and insert the following contants:***
```shell
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOSTNAME=localhost
POSTGRES_PROD_DB=postgres
POSTGRES_TESTING_DB=testingdb
POSTGRES_STUB_DB=stubdb
```

You can manage the Docker container with the usual commands.
```shell
docker compose up
```
will start the server and the postgres in separate containers

You have a couple of pre-set commands that can help you explore the module:
- `make e2etests`: Will launch a smple e2e test suite against a __'testing'__ database seeded ad-hoc.
- `make unittests`: Will launch a simple unit testimg suite against a __'stub'__ database seeded ad-hoc.
- `make demo`: Will launch a server and a client connecting to a __'production'__ database seeded ad-hoc.

As of now, all three dbs are seeded with the same migrations, but can be fully customised independently.
Migrations can be found under 'scripts/db/migrations', labled by descriptive file names.
Seeding data is also under the same folder. During seeding, bids and adjudication are run as postgres db functions and may or may not reflect any business logic found in the server code, they are in fact only meant to seed data in the most effective way, making sure all triggers are fired upon insert.
Using 'make testingdbseed' will log output to the console so you can visualise the result of the queries being run on the db.


## CI/CD
There is a CI/CD gate running on GItHub Actions at every push or PR on the main branch. The branch is 'protected' aginst PRs that do not pass the tests.
Look at the ***CI/CD badge*** at the top of this file to have a spy on the staus of the latest run.


## Description of the solution

### General
The core problem of transaction consistency has been approached by having a single source of truth in the __'transactions'__ table. The table records all transactions between Escrow, Cash and user Account(s).
Main use case is for the __'transactions'__ table to record money exchange relative to a __'Bid'__. In this case, the Bid __'id'__ is explicitely recorded in the transaction record as a foreign key. This makes full reconciliation upon bid adjudication trivial, since the Bid 'id' groups any and all money tranfers. The table can also conveniently be used to record generic transaction unrelated to Bids, for example users geting money in and out of their accounts when jopining or leaving (this could be acomplished in future work).

### DB
In order to ensure data consistency, the db service has been built to expose postgres functions that are designed to be the only interaction surface between the server code and the data layer for write operations. Direct manipulation of tables could be easily prevented by implementing RLS rules, and has been left for a further development sprint due to time constraints.
Reads do not impact data consistency and can therefore be achieved through direct SQL query stirngs or by using convenient views available for the purpose. Same considerations about the enforcement of RLS applies here as for the write operations.
__'Invoice'__ and '__Bid'__ carry a __'state'__ field, which is a postgres `TYPE enum`. Transitions across states are controlled by the aforementioned funcitons to guarantee data consistency.
Triggers control the lowest level of data consistency and operate mostly to update the __'Accounts'__ table __'balance'__ filed.
Since each trigger and each fucntion in postgres constitutes a rollable transaction, __"atomicity"__ is implicitely guaranteed (as said strict RLS rules have been postponed to further work).

### Commands

#### server
The `cmd/server` as methods to list all entities. Furthermore it exposes a method `NewBid()` implementing the __"First come, first adjudicated"__ algorithm that was required in the challenge prompt. `NewBid()` takes an Invoice __'id'__ and calls a sequence of specialised methods on the DB to handle the business requirement of checking and adjudicating the fulfilling bids in the order in which they have been recorded in the DB.
I preferred to locate the algo logic in the 'server' module to keep it separate from the data manipualtion methods that I implemented as funcitons on the postgres DB. This choice seemed more in line wiht a "separation of concers" approach. In theory, by having data-only methods on the DB, one can code additional business logic staying within a higher abstraction layer, which is the server code.
    
***cmd/server list of commands***
```golang
GetAllInvestors();
GetAllIssuers();
GetAllInvoices();
GetInvoiceById();
NewInvoice();
GetAllBids();
GetBidById();
GetBidWithInvoiceById();
GetBidsByInvoiceId();
GetBidsByInvestorId();
NewBid();
```

#### client
The `.cmd/client` has the simple purpose to demo the server code. It runs a sequence of methods on the server that are meant to reproduce an UI interaction. It hits listing endpoints to retrieve all invoices, investors, issuers and bids on and it creates a new invoice with and a new bid.

***cmd/client sequence of calls***
```golang
GetAllInvestors();
GetAllIssuers();
GetAllInvoices();
NewInvoice();
GetAllBids();
GetBidWithInvoiceById();
NewBid();
```


## Work Log

An overview of the sprints I covered is available as a GitHub Project at this link [@MarcoSantonastasi's Golang_challenge](https://github.com/users/MarcoSantonastasi/projects/2/views/4)

### Iteration 0: Learing Go

At the end of Iteration 0 I have a working database with basic seeding, a working gRPC server with some very simple endpoints and a client that performs basic RPCs on those. Also I have a number of unit tests and an end-to-end test as a proof of concept.
Time has been spent mostly learning and applying knowledge at a slow pace to make sure I understood correctly how gRPC and the Go language work.

### Iteration 1: Basic DB APIs ad server methods

At the end of Iteration 1 I have basic database APIs with a corresponding storage model in place, comprising all the entities and the triggers and the fucnitons to support the following basic functionality. Also I have the corresponding methosds on the server module that can interact with the database:
    1. new_invoice(issuer_account_id, ref, denom, amount, ask) => creates a new invoice in the corresponding table
    2. new_bid(bidder_id, invoice, amount) => creates a new bid veryfying sufficient funds and reserving funds.
    3. adjudicate_bid(bid_id) => Sets a bid to won and the relative invoice to adjudicated. It pays out the issuer and retains a fee of 10% for us.
    4. set_all_bids_to_lost(invoice_id) => Sets all bids related to an invoice to lost and refunds the money to investors by creating new transactions that reconcile to zero net money.
    5. List endpoints exposed
    6. Create methosd for Invoice and Bid
    7. Post a Bid and receive the results
    
The DB is seeded with just a few entries for each entity, to make testing possible.
