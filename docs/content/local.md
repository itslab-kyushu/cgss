---
title: Local mode
menu: main
date: 2017-03-15
lastmod: 2017-03-15
weight: 10
description: >-
  Local mode of the client application provides two commands, distribute and
  reconstruct.
  Distribute command reads a given file and computes n shares over m groups
  with a data threshold k and a group threshold l.
  It means totally *n* shares will be made from a secret file,
  and you must have at least k shares from at least l different groups
  to reconstruct the secret.
  Reconstruct command does the reconstruction process;
  i.e. it reads at least *k* share files and reconstruct the secret file.
---
[![Release](https://img.shields.io/badge/release-0.1.0-brightgreen.svg)](https://github.com/itslab-kyushu/cgss/releases/tag/v0.1.0)

## Summary
Local mode of the client application provides two commands, distribute and
reconstruct.
Distribute command reads a given file and computes *n* shares over *m* groups
with a data threshold *k* and a group threshold *l*.
It means totally *n* shares will be made from a secret file,
and you must have at least *k* shares from at least *l* different groups
to reconstruct the secret.

Reconstruct command does the reconstruction process;
i.e. it reads at least *k* share files and reconstruct the secret file.

## Distribute command
```sh
$ cgss local distribute <file> <group threshold> <data threshold> <allocation>
```

This command reads a secret file `<file>` and makes share files.
Allocation takes a comma separated share assignment information.
For example, you want to allocate two shares to the first group,
three shares to the second one, and one share to the last group,
the allocation value is `2,3,1`.

The produced share files has the original filename as the prefix,
and the j-th share for the i-th group has suffix `.i.j.xz`.

This command also takes an optional flag `--chunk` to specify the byte size of
each chunk.
The given secret file is divided to chunks based on this size and distributed
in shares.

## Reconstruct command
```sh
$ cgss local reconstruct <file>...
```

This command reconstructs a secret from a list of share files.
It produces a file based on the given share's file name by removing the above
suffix.
For example, if the names of share files are `sample.txt.1.1.xz`,
`sample.txt.2.1.xz`, ..., then the default file name of the reconstructed secret
will be `sample.txt`.

You can use `--output` flag to use another file name.

## Tutorial
Suppose `secret.dat` is a secret file and distributing it allocations share to
three groups so that each group has two shares with both group and data
threshold are 3, the following commands creates totally six shares:

```shell
$ cgss local distribute secret.dat 3 3 2,2,2
```

The above command creates a set of secrets, `secret.dat.1.1.xz`,
`secret.dat.1.2.xz`, `secret.dat.2.1.xz`, ..., `secret.dat.3.2.xz`.
We can store each share file into a different storage in order to prevent
information leakage, and now we can delete the secret file `secret.dat`.

To reconstruct the secret from shares, we must to collect at least 3 share
files from each group. Suppose we have `secret.dat.1.1.xz`, `secret.dat.2.1.xz`,
and `secret.dat.3.1.xz`.

```shell
$ cgss local reconstruct secret.dat.1.1.xz secret.dat.2.1.xz secret.dat.3.1.xz
```

The above command reconstructs the secret and stores it as `secret.dat`.


## Installation
If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```sh
$ brew tap itslab-kyushu/cgss
$ brew install cgss
```

Compiled binaries for some platforms are available on
[Github](https://github.com/itslab-kyushu/cgss/releases).
To use these binaries, after downloading a binary to your environment,
decompress and put it in a directory included in your $PATH.

You can also compile the client application by yourself.
To compile it, you first download the code base:

```shell
$ git clone https://github.com/itslab-kyushu/cgss $GOPATH/src/itslab-kyushu/cgss
```

Then, build the client application `cgss`:

```shell
$ cd $GOPATH/src/itslab-kyushu/cgss/client
$ go get -d -t -v .
$ go build -o cgss
```

To build the command, [Go](https://golang.org/) > 1.7.4 is required.
