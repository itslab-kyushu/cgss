---
title: Remote mode
menu: main
date: 2017-03-15
lastmod: 2017-03-15
weight: 15
description: >-
  Remote mode of the client application provides an interface to a client/server
  style data storage service. This mode has four commands:
  get is downloading a set of shares and reconstruct a secret,
  put is creating a set of shares from a secret and upload each share to each
  server,
  delete is deleting shares associated with a given secret file,
  list is showing names of secret files stored in the data storage service.
---
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/cgss/releases/tag/v0.2.0)

## Summary
Remote mode of the client application provides an interface to a client/server
style data storage service.
This mode has four commands:

* get: download a set of shares and reconstruct a secret,
* put: creates a set of shares from a secret and upload each share to each
  server,
* delete: delete shares associated with a given secret file,
* list: show names of secret files stored in the data storage service,

and all commands connect a set of [data storage servers](../server).

To specify address information of those servers, all commands receives a
configuration file in YAML.
The configuration file has a root element `groups`, which is a list of group
information; a group information has `name` and `servers` elements.
`name` represents a name of the group and `servers` takes a list of server
information, which is an object consisting of two element `address` and `port`.
The following example defines two groups; group1 has two servers and group2
has one server:

```yaml
group:
  - name: Group1
    servers:
    - address: 192.168.0.1
      port: 13009
    - address: 192.168.0.2
      port: 13009
  - name: Group2
    servers:
    - address: 192.168.1.1
      port: 13009
```

The default name of the configuration file is `cgss.yml` but you can set another
name via `--config` flag.

The get command gathers shares from the servers defined the configuration file,
and put command distributes shares to the servers.

## Get command
```shell
$ cgss remote get --config cgss.yml --output result.dat <file name>
```

Get command gathers shares associated with the given file name from the servers
defined in the configuration file, and then reconstructs and stores them as
the given file name via `--output` flag.

The number of servers defined in the configuration file must be greater then or
equal to the threshold which is used to put the secret file.

This command downloads shares from the servers defined the configuration file,
even if the number of necessary shares is smaller than the number of the servers
defined in the configuration file.
In other words, it is not good to use the same configuration file when you used
in `put` command.

If `--config` flag is omitted, `cgss.yml` is used, and if `--output` flag is
omitted, `<file name>` is used.

To find available file names, use list command.


## Put command
```shell
$ cgss remote put --config cgss.yml <file> <group threshold> <data threshold>
```

Put command reads the given file and runs distribute procedure to create shares.
The number of total shares are as same as defined in the server configuration
file.

If `--config` flag is omitted, `cgss.yml` is used.

Put command also takes `--chunk` flag to set the byte size of each chunk.
The default value is 256.
The distribute procedure creates a finite filed Z/pZ, where p is a prime number
which has chunk size + 1 bit length.

## Delete command
```shell
$ cgss remote delete --config cgss.yml <file name>
```

Delete command deletes all shares associated with the given file name from all
servers defined in the configuration file.

If `--config` flag is omitted, `cgss.yml` is used.

## List command
```shell
$ cgss remote list --config cgss.yml
```

List command shows all file names stored in the servers.
If `--config` flag is omitted, `cgss.yml` is used.

## Installation
If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```shell
$ brew tap itslab-kyushu/cgss
$ brew install cgss
```

Compiled binaries for some platforms are available on
[Github](https://github.com/itslab-kyushu/cgss/releases).
To use these binaries, after downloading a binary to your environment, decompress and put it in a directory included in your $PATH.

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
