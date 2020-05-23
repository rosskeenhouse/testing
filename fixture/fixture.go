// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package fixture

import (
  "testing"
  "strconv"
)

type Test func(t *testing.T)

type F func(p Param) interface{}

type Ft struct {
  *testing.T
  params *Parameters
  fixture F
  data []interface{}
  results *Results
}

func New(t *testing.T, f F, p *Parameters, r *Results) *Ft {
  return &Ft{ T: t, params: p, fixture: f, results: r }
}

func (f *Ft) Fixture() []interface{} {
  return f.data
}

func (f *Ft) RunWith(t Test) {
  f.data = nil
  for i, r := range(f.params.Values()) {
    f.data = append(f.data, f.fixture(r))
    f.Run(strconv.Itoa(i), t)
  }
}

func (f *Ft) Assert() {
  r := f.results.Values()
  i := len(f.data) - 1
  if f.data[i] != r[i] {
    f.Errorf("Failed value does not match expected result: %s != %s", f.data[i], r[i])
  }
}

