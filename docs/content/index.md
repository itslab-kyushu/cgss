---
title: Cross-group Secret Sharing Scheme
type: homepage
menu:
  main:
    Name: Top
date: 2017-03-15
lastmod: 2017-03-15
weight: 1
description: >-
  This software provides both a Go library and command line tools implementing
  a threshold Secret Sharing scheme.
---
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/cgss/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/cgss/tree/master)
[![wercker status](https://app.wercker.com/status/062e76e7dff821dec72d1751c55b3402/s/master "wercker status")](https://app.wercker.com/project/byKey/062e76e7dff821dec72d1751c55b3402)
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/cgss/releases/tag/v0.2.0)

## Summary
Cloud datastore services have been essential to access your data from anywhere.
However, they still have risks of
data leakage since each cloud service occasionally becomes a target from
malicious attackers
and data loss since hard-wares of each cloud service occasionally are broken.
To solve this problem, we've introduced *cross-group secret sharing scheme*,
which distributes a secret into servers from several could providers so that
the secret can be reconstructed even if some of servers are attacked or become
out-of-service.

This software provides both a [Go](https://golang.org/)
[library](https://godoc.org/github.com/itslab-kyushu/cgss/cgss) and
command line tools implementing the cross-group secret sharing scheme.


## Contents
* To use the cross-group secret sharing from another go application,
  see the [API Reference](api) page.  
* To compute shares and reconstruct a secret in a computer,
  see the [local mode usage](local) page.
* To use a cross-group secret sharing based data storage service,
  see the [client usage](remote) and [server usage](server) pages.

## Publications

* [Hiroaki Anada](http://sun.ac.jp/prof/anada/),
  [Junpei Kawamoto](https://www.jkawamoto.info),
  Chenyutao Ke,
  [Kirill Morozov](http://www.is.c.titech.ac.jp/~morozov/), and
  [Kouichi Sakurai](http://itslab.inf.kyushu-u.ac.jp/~sakurai/),
  "[Cross-Group Secret Sharing Scheme for Secure Usage of Cloud Storage over Different Providers and Regions](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/article/10.1007%2Fs11227-017-2009-7),"
  [The Journal of Supercomputing](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/journal/11227),
  73(10), pp.4275-4301, 2017
* Chenyutao Ke,
  [Hiroaki Anada](http://sun.ac.jp/prof/anada/),
  [Junpei Kawamoto](https://www.jkawamoto.info),
  [Kirill Morozov](http://www.is.c.titech.ac.jp/~morozov/), and
  [Kouichi Sakurai](http://itslab.inf.kyushu-u.ac.jp/~sakurai/),
  "[Cross-group Secret Sharing for Secure Cloud Storage Service](http://hdl.handle.net/2324/1563374),"
  Proc. of the Annual International Conference on Ubiquitous Information Management and Communication (IMCOM 2016),
  pp.63:1-63:8, Vietnam, Jan.4-6, 2016.

Please consider to site those papers if you will publish articles using this application.

## License
This software is released under The GNU General Public License Version 3,
see [license](./licenses/) for more detail.
