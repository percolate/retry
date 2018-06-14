# ReTry

Percolate's Go retry package

## Description

ReTry is a simple Go package for implementing retry logic. It's partially based on the Python package, [retry](https://github.com/invl/retry).

## Installation

`go get github.com/percolate/retry`

## Usage

It's easy! Configure an instance of `Re` to your liking, and pass its `Try` method a `Func` of your choosing.

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
    var body []byte
    
    err := retry.Re{Max: 3, Delay: time.Duration(10*time.Millisecond)}.Try(func() error {
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
