# arex_challenge
see (./REQUIREMENTS.md)[REQUIREMENTS.md] for the challenge prompt.

## Intall and run 
You have to have Docker deamon and dockercompose alreeady installed on your machine, then you also have to make sure the 'make' executable is part of your PATH.

You have a couple of pre-set commands that can help you explore the module:
- 'make demo': Will launche a server and a client connecting to a "production" database
- 'make e2etest': Will launch a smple e2e test suite against a 'production' database

## Description

## Work Log

### Iteration 0: Learing Go

At the end of Iteration 0 I have a working database with basic seeding, a working gRPC server with some very simple endpoints and a client that performs basic RPCs on those. Also I have a number of unit tests and an end-to-end test as a proof of concept.
Time has been spent mostly learning and applying knowledge at a slow pace to make sure I understood correctly how gRPC and the Go language work.

### Iteration 1: Basic DB APIs

At the end of Iteration 1 I have basic database APIs with a corresponding storage model in place, comprising all the entities and the triggers and the fucnitons to support the following basic functionality:
    1. new_invoice(issuer_account_id, ref, denom, amount, ask) => creates a new invoice in the corresponding table
    1. new_bid(bidder_id, invoice, amount) => creates a new bid veryfying sufficient funds and reserving funds
The DB is also seeded accordingly with just a few entries to make testing possible
