// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package fixture

import (
  "testing"
  "strconv"
  "reflect"
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

func (f *Ft) Value() interface{} {
  return f.data[len(f.data) - 1]
}

func (f *Ft) Result() interface{} {
  return f.results.Values()[len(f.data) - 1]
}

func (f *Ft) RunWith(t Test) {
  f.data = nil
  for i, r := range(f.params.Values()) {
    f.data = append(f.data, f.fixture(r))
    f.Run(strconv.Itoa(i), t)
  }
}

func (f *Ft) assertOp(value interface{}, o func (a interface{}, b interface{}) bool) {
  r := f.Result()
  if ! o(r, value) {
    f.Errorf("Failure: value does not match expected result: [%s] != [%s]", r, value)
  }

}

func (f *Ft) AssertStrEq(value string) {
  f.assertOp(value, func (a interface{}, b interface{}) bool { return a.(string) == b.(string) })
}

func (f *Ft) AssertGt(value interface{}) {
  cmpOp := func (a, b interface{}) bool {
    rType := reflect.TypeOf(a)
    switch rType.Name() { 
    case "float64", "float32":
      return reflect.ValueOf(a).Convert(rType).Float() > reflect.ValueOf(b).Convert(rType).Float()
    default:
      return reflect.ValueOf(a).Convert(rType).Int() > reflect.ValueOf(b).Convert(rType).Int()
    }
  }
  f.assertOp(value, cmpOp)
}

func (f *Ft) AssertGe(value interface{}) {
  cmpOp := func (a, b interface{}) bool {
    rType := reflect.TypeOf(a)
    switch rType.Name() { 
    case "float64", "float32":
      return reflect.ValueOf(a).Convert(rType).Float() >= reflect.ValueOf(b).Convert(rType).Float()
    default:
      return reflect.ValueOf(a).Convert(rType).Int() >= reflect.ValueOf(b).Convert(rType).Int()
    }
  }
  f.assertOp(value, cmpOp)
}

func (f *Ft) AssertLt(value interface{}) {
  cmpOp := func (a, b interface{}) bool {
    rType := reflect.TypeOf(a)
    switch rType.Name() { 
    case "float64", "float32":
      return reflect.ValueOf(a).Convert(rType).Float() < reflect.ValueOf(b).Convert(rType).Float()
    default:
      return reflect.ValueOf(a).Convert(rType).Int() < reflect.ValueOf(b).Convert(rType).Int()
    }
  }
  f.assertOp(value, cmpOp)
}

func (f *Ft) AssertLe(value interface{}) {
  cmpOp := func (a, b interface{}) bool {
    rType := reflect.TypeOf(a)
    switch rType.Name() { 
    case "float64", "float32":
      return reflect.ValueOf(a).Convert(rType).Float() <= reflect.ValueOf(b).Convert(rType).Float()
    default:
      return reflect.ValueOf(a).Convert(rType).Int() <= reflect.ValueOf(b).Convert(rType).Int()
    }
  }
  f.assertOp(value, cmpOp)
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

