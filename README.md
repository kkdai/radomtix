RandomTix: The Radom Ticket invitation code (Random string) generator in Go.
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/radomtix/master/LICENSE)![Go](https://github.com/kkdai/radomtix/workflows/Go/badge.svg)


This is a console tool will generate invitation code by requeust length and number of code you need. 

This tool for community organizers who need to generate meetup invitation code by themselves.


Install
--------------

    go get -u -x github.com/kkdai/radomtix

Usage
---------------------
    //To display 10 invitation code with length 5
    radomtix -n 10  -l 5

    //Response.
    EHXPM
    SYTKS
    UFIRP
    JWSTS
    BRGCE
    SOYGS
    PPAIQ
    UPKRJ
    UCTHD
    TTJNX



Options
---------------

- `-n` number of tickets. (default is 10)
- `-l` length of string. (default is 6)


Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`


License
---------------

This package is licensed under MIT license. See LICENSE for details.
