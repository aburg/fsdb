# fsdb

A file system database.

This is wholly unsafe for anything. Do not use it. :-D

## Installation

```[bash]
go install
```

## Usage

TODO

## Commands

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
fsdb id customer.node
fsdb deid customer.node
```

### level 2

This is what you want to use most of the time.

```[bash]
fsdb pool init project
fsdb spawn "Project 1" "Project 2"

# if this is a multipool you have to specify the kind
fsdb spawn create --kind project "Project 1" "Project 2"
```
