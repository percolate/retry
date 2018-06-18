# ReTry

[![circleci](https://circleci.com/gh/percolate/retry.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/percolate/retry)
[![codecov](https://codecov.io/gh/percolate/retry/branch/master/graph/badge.svg?token=p3WLUQEQqf)](https://codecov.io/gh/percolate/retry)

Percolate's Go retry package

## Description

ReTry is a simple Go package for implementing retry logic. It's partially based
on the Python package, [retry](https://github.com/invl/retry).

## Installation

`go get github.com/percolate/retry`

## Usage

It's easy! Configure an instance of `Re` to your liking, and pass its `Try`
method a `Func` of your choosing.

## Example

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/percolate/retry"
)

func main() {

    url := "http://example.com"
    delay := time.Duration(10*time.Millisecond)

    var body []byte
    err := retry.Re{Max: 3, Delay: delay}.Try(func() error {
        resp, err := http.Get(url)
        if err != nil {
            return err
        }

        defer resp.Body.Close()
        body, err = ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(body)
}
```