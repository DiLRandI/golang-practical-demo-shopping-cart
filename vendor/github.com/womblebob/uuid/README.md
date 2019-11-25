
# uuid ![build status](https://travis-ci.org/womblebob/uuid.svg?branch=master)
The uuid package generates and inspects UUIDs based on [RFC 4122](http://tools.ietf.org/html/rfc4122) and DCE 1.1: Authentication and Security Services. 

###### Install
`go get github.com/womblebob/uuid`

###### Documentation 
[![GoDoc](https://godoc.org/github.com/womblebob/uuid?status.svg)](http://godoc.org/github.com/womblebob/uuid)

Full `go doc` style documentation for the package can be viewed online without installing this package by using the GoDoc site here: 
http://godoc.org/github.com/womblebob/uuid

This is a fork of code.google.com/p/go-uuid with only one modification, the default implementation of the Scanner interface for database marshalling does not work for SQL Server due to the fact that it uses a different endian format.

uuid.DefaultScanner = uuid.SqlServerScanner{}

will change the behaviour of the library to use the logic taken from https://github.com/denisenkom/go-mssqldb/uniqueidentifier.go and ensure that this uuid can be used across all structs.
