# fsdb

A file system database.

## commands

### level 0

This is totally unsafe.

```[bash]
fsdb set token myToken
fsdb get token
fsdb unset token
```

### level 1

This is the stuff working with uuids.

```[bash]
fsdb make customer.node
fsdb unmake customer.node
fsdb is
fsdb is customer.node
```

### level 2

This is what you want to use most of the time.

```[bash]
fsdb spawn
```
