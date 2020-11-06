## Thunes Money Transfer Backend

[![Build Status](https://travis-ci.com/fghpdf/thunes_sample.svg?token=nEHz11y83KAa9EpLU95Y&branch=develop)](https://travis-ci.com/fghpdf/thunes_sample)

---

### Overview

This project is for [thunes API](https://developers.thunes.com/money-transfer/v2/?go#example-flow) to transfer money by golang.



### Directory

Root
```shell
.
├── api
├── config
├── internal
│   ├── app
│   └── pkg
│       ├── common
│       ├── config
│       ├── country
│       ├── payer
│       ├── ping
│       ├── quotation
│       ├── routers
│       ├── thunes    // thunes API client
│       └── transaction
└── test
    └── mocks
```
In `internal/pkg`, exclude the `thunes`, all is the web service for money transfering. 

It offered many useful interface by HTTP.

quotation/
```
.
├── handler.go
├── handler_test.go
└── quotation.go
```
In this directory, `quotation.go` is the model for interface. 

And all interface and implement is in `handler.go`. 

All unit test is in `handler_test.go`.

### Start

You can use command `go run main.go` to start a HTTP service.

It will listen `:8080`.

### Test

This project contains unit test and integration test.

You can check the `Makefile`.

Uint test: `make unittest`
Integration test: `make integration_test`

Also you want to see coverage for unit test:
```
make coverage
```

### CI/CD
This project is connected to Travis.

So there are tags in this document header.

A picture for automated test CI
[![Bfc4QU.png](https://s1.ax1x.com/2020/11/06/Bfc4QU.png)](https://imgchr.com/i/Bfc4QU)

### What's Next

Duo to time limited, some detail isnot very well.

This project can be better.

- [] static string
- [] more edge unit test
- [] more error handler
- [] more log
- [] opentracing
- [] Docker deploy

and so on.
