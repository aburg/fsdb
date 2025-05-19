# fsdb

A file system database.

This is wholly unsafe for anything. Do not use it. :-D

## Installation

```[bash]
go install
```

## Usage

TODO

## What? How?

### Pools

- any directory can be made to be a pool that holds nodes of a specified (pool-)kind
- this is notified by a file ".<kind>.pool"
  - this will hold a uuid

Commands:

```[bash]
fsdb pool init <kind...> # make current directory a pool of given kind
```

### Nodes

- any directory can be made to be a node of a specified (node-)kind
- this is notified by a file ".<kind>.node"
  - this will hold a uuid

Commands:

```[bash]
fsdb node init <kind...> # make current directory a node of given kind
```

### Data

- any directory can hold data (each datum consisting of a kind and a value)
- file ".<kind>.datum"
  - this will contain the value for the datum

Commands:

```[bash]
fsdb data set <kind> <value>
fsdb data get <kind>
fsdb data remove <kind>
```

Upcoming commands:

```[bash]
fsdb data list
```

### Spawning

Spawning is the process of creating a node in a pool.

### Writing

Writing is the process of putting data onto a node.

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
fsdb pool project
fsdb spawn "Project 1" "Project 2"
fsdb depool project

# if this is a multipool you have to specify the kind
fsdb spawn create --kind project "Project 1" "Project 2"
```

## TODO

- implement subcommands
- clean up/refactor (pool, node, data)
- create node commands (node id, node deid, node list)
- rework this with subcommands (pool list, pool create)
