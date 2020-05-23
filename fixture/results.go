// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package fixture

import (
)

type Result interface{}

type Results struct {
  values []Result
}

func R(v []Result) *Results { return &Results{ values: v } } 
 
func (r *Results) Values() []Result {
  return r.values
}

func (r *Results) Read() <-chan Result {
  rc := make(chan Result, len(r.values))
  for i := range(r.Values()) {
    rc <- i
  }
  return rc
}
