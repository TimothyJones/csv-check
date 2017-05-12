csv-check
=========

Is an ultra light command line tool to check that a csv file matches [RFC 4180](https://www.ietf.org/rfc/rfc4180.txt) format.

It reads from standard in and writes to standard error. It will exit on the first error encountered.

Usage: 

     cat <malformed.csv> | csv-check

Installation:

     go get github.com/TimothyJones/csv-check

##See also

For a simple way to fix common malformed CSV files, see [csv-clean](http://github.com/TimothyJones/csv-clean)

