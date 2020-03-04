acmutil
=======

[![Build Status](http://img.shields.io/travis/Actinium-project/acmutil.svg)](https://travis-ci.org/Actinium-project/acmutil) 
[![Coverage Status](http://img.shields.io/coveralls/Actinium-project/acmutil.svg)](https://coveralls.io/r/Actinium-project/acmutil?branch=master) 
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Actinium-project/acmutil)

Package acmutil provides litecoin-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for acmd, an alternative full-node implementation of
litecoin based on acmd, which is under active development by Conformal.
Although it was primarily written for acmd, this package has intentionally been
designed so it can be used as a standalone package for any projects needing the
functionality provided.

## Installation and Updating

```bash
$ go get -u github.com/Actinium-project/acmutil
```

## License

Package acmutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
