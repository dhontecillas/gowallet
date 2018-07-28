# gowallet


## TODO

- Test that a user can not view another user wallet
- Add good error handling: Define Error types that can be understood at
    the REST interface level to produce correct error codes
- Add Mutex for in-memory storage: Currently, the storage that we
    are passing to each request is the same, and is not thread safe.
    A sim
- Add pagination for wallet listing endpoint
- Add transactions to in-memory storage: Currently, in memory storage
    has the `Transactional` interface, but functions are not
    implemented.
