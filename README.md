# cgss
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
Cross-Group Secret Sharing.

This software implements the Cross-Group Secret Sharing scheme proposed in [Cross-group Secret Sharing for Secure Cloud Storage Service](http://hdl.handle.net/2324/1563374).

It is still under construction, and it currently supports the simple Shamir's
Secret Sharing scheme.

## Installation
```
$ git clone https://github.com/itslab-kyushu/cgss
$ cd cgss
$ go build
```

## Usage
### Distribute
```
$ cgss distribute <file> <number of shares> <threshold>
```

It produces share files and the file name of i-th share has `.i.json` as the
suffix.

### Reconstruct
```
$ cgss reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.


## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
