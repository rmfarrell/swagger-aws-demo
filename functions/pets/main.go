package main

import (
  "encoding/json"
  "github.com/apex/go-apex"
)

// Example of a Lambda function handling arbitrary JSON input.
func main() {
  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

    var m interface{}

    if err := json.Unmarshal(event, &m); err != nil {
      return nil, err
    }
    return &m, nil
  })
}