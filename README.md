# cgss
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)

Cross-Group Secret Sharing.

This software implements the Cross-Group Secret Sharing scheme proposed in [Cross-group Secret Sharing for Secure Cloud Storage Service](http://hdl.handle.net/2324/1563374).

## Installation
```
$ git clone https://github.com/itslab-kyushu/cgss
$ cd cgss
$ go build
```

## Usage
### Distribute
```
$ cgss local distribute <file> <group threshold> <data threshold> <allocation>
```

It produces share files.
Allocation takes a comma separated allocations.
If you want to allocate two shares to the first group, three shares to the
second one, and one share to the last group, the allocation value is `2,3,1`.

The produced share files has the original filename as the prefix,
and the j-th share for the i-th group has suffix `.i.j.json`.

### Reconstruct
```
$ cgss local reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.

### simple
It provides the simple Shamir's Secret Sharing scheme.

#### Distribute
```
$ cgss simple distribute <file> <threshold> <number of shares>
```

It produces share files and the file name of i-th share has `.i.json` as the
suffix.

#### Reconstruct
```
$ cgss simple reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.


## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
