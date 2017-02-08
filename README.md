# cgss
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)

Cross-Group Secret Sharing.

This software implements the Cross-Group Secret Sharing scheme proposed in
[Cross-group Secret Sharing for Secure Cloud Storage Service](http://hdl.handle.net/2324/1563374).

## Installation
Download the code base:

```
$ git clone https://github.com/itslab-kyushu/cgss $GOPATH/src/itslab-kyushu/cgss
```

Build client command `cgss`:

```
$ cd $GOPATH/src/itslab-kyushu/cgss/client
$ go get -d -t -v .
$ go build
```

Build server command `cgss-server`:

```
$ cd $GOPATH/src/itslab-kyushu/cgss/server
$ go get -d -t -v .
$ go build
```

To build both commands, [Go](https://golang.org/) > 1.7.4 is required.

## Client Usage
The client application provides two way to run the Cross-Group Secret Sharing
(CGSS) scheme.
One of them is local mode, which stores shares into a local file system.
It is suitable to test our CGSS scheme easily.
The other one is remote mode, which stores shares into servers provided the
server command.

### Local mode
The local mode provides two sub commands, distribute and reconstruct.
Distribute command reads a file and creates a set of shares,
on the other hand, reconstruct command reads a set of shares and reconstructs
the original file.

#### Distribute
```
$ cgss local distribute <file> <group threshold> <data threshold> <allocation>
```

It produces share files.
Allocation takes a comma separated allocations.
If you want to allocate two shares to the first group, three shares to the
second one, and one share to the last group, the allocation value is `2,3,1`.

The produced share files has the original filename as the prefix,
and the j-th share for the i-th group has suffix `.i.j.json`.

#### Reconstruct
```
$ cgss local reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.

### Remote mode
Remote mode provides four sub command: get, put, delete, and list.
All commands take a YAML based server configuration file.
The format is as follows:

```yaml
groups:
  - name: Group-1
    servers:
      - address: 192.168.0.1
        port: 13009
      - address: 192.168.0.2
        port: 13009
  - name: Group-2
    servers:
      - address: 192.168.1.1
        port: 13009
```

The above example defines two groups, Group-1 and Group-2,
and two servers in the Group-1 and one server in the Group-2.

The get command gathers shares from the servers defined the configuration file,
and put command distributes shares to the servers.

The default name of the configuration file is `cgss.yml` but you can set another
name via `--config` flag.

#### Get
```
cgss remote get --config cgss.yml --output result.dat <file name>
```

Get command gathers shares associated with the given file name from the servers
defined in the configuration file, and then reconstructs and stores them as
the given file name via `--output` flag.

If `--config` flag is omitted, `cgss.yml` is used, and if `--output` flag is
omitted, `<file name>` is used.

To find available file names, use list command.

The number of groups and the number of total servers must be greater then or
equal to the group threshold and the data threshold, which are given when those
shares were created.

#### Put
```
cgss remote put --config cgss.yml <file> <group threshold> <data threshold>
```

Put command reads the given file and runs distribute procedure to create shares.
The group threshold and the data threshold are parameters of CGSS scheme.
The number of groups and the number of total shares are as same as defined in
the server configuration file.

If `--config` flag is omitted, `cgss.yml` is used.

For example, if you use the above example configuration, put command creates
two shares to the Group-1 and one share to the Group-2.

Put command also takes `--chunk` flag to set the byte size of each chunk.
The default value is 256.
The distribute procedure creates a finite filed Z/pZ, where p is a prime number
which has chunk size + 1 bit length.

### Delete
```
cgss remote delete --config cgss.yml <file name>
```

Delete command deletes all shares associated with the given file name from all
servers defined in the configuration file.

If `--config` flag is omitted, `cgss.yml` is used.

### List
```
cgss remote list --config cgss.yml
```

List command shows all file names stored in the servers.
If `--config` flag is omitted, `cgss.yml` is used.


## Server Usage
The server application runs a simple data store service using CGSS scheme.

It takes three flags,
* `--port`: the port number the server will listen,
* `--root`: the document root path to store uploaded shares,
* `--no-compress`: if set, all shares will be stored without compression.

If those flags are omitted, default values are used.
Thus, you can start a server by just run `cgss-server`.


## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
