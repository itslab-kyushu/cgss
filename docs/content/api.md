---
title: API Reference
menu: main
date: 2017-03-15
lastmod: 2017-03-15
weight: 30
description: >-
  Package github.com/itslab-kyushu/cgss/cgss provides Distribute and Reconstruct
  functions, which creates a set of shares from a given secret and reconstruct
  the secret from a subset of the shares, respectively.
---
[![GoDoc](https://godoc.org/github.com/itslab-kyushu/cgss/cgss?status.svg)](https://godoc.org/github.com/itslab-kyushu/cgss/cgss)

## Summary
Package `github.com/itslab-kyushu/cgss/cgss` provides Distribute and Reconstruct
functions, which creates a set of shares from a given secret and reconstruct
the secret from a subset of the shares, respectively.

## Installation
```sh
$ go get -d github.com/itslab-kyushu/cgss
```

## Distribute a secret
In cross-group secret sharing scheme, shares are divided into groups,
and you can choose how many shares are assigned in each group.
We call the assignment as share *allocation*.
Type [`Allocation`](https://godoc.org/github.com/itslab-kyushu/cgss/cgss#Allocation),
which is a list of integers, is used to define a share allocation.
For example, `Allocation{3, 3, 2}` means three shares are assigned to the first
and second groups and two shares are assigned to the third group.

[`Distribute`](https://godoc.org/github.com/itslab-kyushu/cgss/cgss#Distribute)
function, which computes shares from a secret, takes an allocation,
a group threshold, and a data threshold in addition to the secret.
The group threshold *l* constrains shares must be gathered from at least *l*
groups
and the data threshold *k* constrains at least *k* share must be gathered
in order to reconstruct the secret.

The Distribute function also takes chunk size parameter.
Bigger secrets will be divided based on the size before they are distributed.
Smaller chunk size may cause of overhead to handle fragments but
bigger chunk size may cause of huge computation to find prime numbers used in
the secret sharing process.
We use 8 for all evaluations.

### Example
The following example assumes three groups and assigns two shares to
each of them, and the group and data thresholds are set to 2 and 3,
respectively. It means, to reconstruct the secret, at least three shares must
be collected from at least 2 groups.

Since the chunk size is set to 8bytes, the secret will be divided every
8bytes and each chunk will be converted to a set of shares.

From the share assignment, the distribute function makes totally 6 shares,
and it thus returns a slice of shares of which the length is 6.
Note that the returned shares do not have any information about groups but
the order of them are associated with the given allocation.
More precisely, shares[0] and shares[1] are for the group 1, shares[2] and
shares[3] are for the group 2, and shares[4] and shares[5] are for the
group 3 in this example.

```go
secret := []byte(`abcdefgaerogih:weori:ih:opih:oeijhg@roeinv;dlkjh:
	roihg:3pw9bdlnbmxznd:lah:orsihg:operinbk:sldfj:aporinb`)

ctx := context.Background()
shares, err := Distribute(ctx, secret, &DistributeOpt{
	ChunkSize:      8,
	Allocation:     Allocation{2, 2, 2},
	GroupThreshold: 2,
	DataThreshold:  3,
}, nil)
if err != nil {
	log.Fatal(err.Error())
}
```

## Reconstruct a secret
[`Reconstruct`](https://godoc.org/github.com/itslab-kyushu/cgss/cgss#Reconstruct)
function takes a set of shares and returns the secret in byte slice.
The given shares must satisfy the group and data constraints set to
distribute the secret.
If the number of groups and/or the number of shares are not enough, the result
will be a list of meaningless bytes.

### Example
The following example reconstructs the secret from a subset of distributed
shares made in the above example.
Since the group threshold and data threshold are 2 and 3, respectively,
we use two shares from group 1 and one share from group 2, i.e. three shares
from two groups, to reconstruct the secret.

```go
// Pick up two shares from group 1 and one share from group 2.
subset := []Share{shares[0], shares[1], shares[2]}
res, err := Reconstruct(ctx, subset, nil)
if err != nil {
  log.Fatal(err.Error())
}
fmt.Println(string(res))
```

## Marshal/Unmarshal shares
All components of a share implement
[Marshaler](https://golang.org/pkg/encoding/json/#Marshaler)
interface so that it can be marshaled and unmarshaled to/from a JSON document.

```go
// share is a share object.
bytes, err := json.Marshal(share)
if err != nil {
  log.Fatal(err.Error())
}

var res Share
if err = json.Unmarshal(bytes, &res); err != nil {
  log.Error(err.Error())
}
```
