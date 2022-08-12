# arex_challenge

## Intall and run


## Description

## Work Log

### Iteration 0: Learing Go

At the end of Iteration 0 I have a working database with basic seeding, a working gRPC server with some very simple endpoints and a client that performs basic RPCs on those. Also I have a number of unit tests and an end-to-end test as a proof of concept.
Time has been spent mostly learning and applying knowledge at a slow pace to make sure I understood correctly how gRPC and the Go language work.

### Iteration 1: Basic DB APIs

At the end of Iteration 1 I have basic database APIs with a corresponding storage model in place, comprising all the entities and the triggers and the fucnitons to support the following basic functionality:
    1. new_invoice(issuer_id, ref, denom, amount, ask) => creates a new invoice in the corresponding table
    1. new_bid(bidder_id, invoice, amount) => creates a new bid veryfying sufficient funds and reserving funds
The DB is also seeded accordingly with just a few entries to make testing possible