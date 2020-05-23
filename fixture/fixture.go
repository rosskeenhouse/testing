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
  params ParamReader
  fixture F
  data []interface{}
  results *Results
}

func New(t *testing.T, f F, p ParamReader, r *Results) *Ft {
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

func (f *Ft) AssertEqual(value interface{}) {
  r := f.results.Values()
  i := len(f.data) - 1
  switch r[i].(type) {
  case string:
    res := r[i].(string)
    v := value.(string)
  default:
    res := r[i]
    v := value
  }
  if value != res {
    f.Errorf("Failed value does not match expected result: [%s] != [%s]", value, res)
  }
}

func (f *Ft) Assert() {
  r := f.results.Values()
  i := len(f.data) - 1
  data := f.data[i]
  res := r[i]
  switch r[i].(type) {
  case string:
    data = string(f.data[i].([]byte))
    res = string(r[i].(string))
  }
  if data != res {
    f.Errorf("Failed value does not match expected result: [%s] != [%s]", f.data[i], r[i])
  }
}

