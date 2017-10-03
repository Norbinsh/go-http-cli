### go-http-cli [![Build Status](https://travis-ci.org/visola/go-http-cli.svg?branch=master)](https://travis-ci.org/visola/go-http-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/visola/go-http-cli)](https://goreportcard.com/report/github.com/visola/go-http-cli)

An HTTP client inspired by `curl` made in Go.

### Example

Example command pointing to a test server:

```bash
go-http-cli \
  -H Content-Type=application/json \
  -X POST \
  -d "Some test data" \
  http://localhost:3000/test
```

Output:

```
POST http://localhost:3000/test
>> 'Content-Type' = 'application/json'
>>
>> Some test data
--
<< 'Connection' = '[keep-alive]'
<< 'X-Powered-By' = '[Express]'
<< 'Content-Type' = '[text/html; charset=utf-8]'
<< 'Content-Length' = '[12]'
<< 'Etag' = '[W/"c-Lve95gjOVATpfV8EL5X4nxwjKHE"]'
<< 'Date' = '[Fri, 29 Sep 2017 13:52:50 GMT]'
<<
<< Hello World!
<<
```
