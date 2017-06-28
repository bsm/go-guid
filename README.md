# GUID

[![Build Status](https://travis-ci.org/bsm/go-guid.png?branch=master)](https://travis-ci.org/bsm/go-guid)
[![GoDoc](https://godoc.org/github.com/bsm/go-guid?status.png)](http://godoc.org/github.com/bsm/go-guid)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/go-guid)](https://goreportcard.com/report/github.com/bsm/go-guid)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Simple, thread-safe MongoDB style GUID generator.

### Examples

```go
func main {
	// Create a new 12-byte globally-unique identifier
	id := guid.New96()
	fmt.Println(hex.EncodeToString(id.Bytes()))
}
```

```go
func main {
	// Create a new 16-byte globally-unique identifier
	id := guid.New128()
	fmt.Println(hex.EncodeToString(id.Bytes()))
}
```
