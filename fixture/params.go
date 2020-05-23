// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package fixture

import (
)

// Fixture params are a list of values used to initialize a fixture
type Param interface{}

type Parameters struct {
  values []Param
}

func P(v []Param) *Parameters { return &Parameters{ values: v } } 
 
func (p *Parameters) Values() []Param {
  return p.values
}

func (p *Parameters) Read() <-chan Param {
  rc := make(chan Param, len(p.values))
  for i := range(p.Values()) {
    rc <- i
  }
  return rc
}
