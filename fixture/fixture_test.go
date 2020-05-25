// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package fixture_test

import (
  "testing"
  "math"
  "github.com/rosskeenhouse/testing/fixture"
)

func Fixture2Pow(p fixture.Param) interface{} {
  return math.Exp2(p.(float64))
}

func TestNewFixture(t *testing.T) {
  f := fixture.New(t, Fixture2Pow, fixture.P([]fixture.Param{1.0,2.0,3.0}), fixture.R([]fixture.Result{2.0,4.0,8.0}))
  
  f.RunWith(
    func (t *testing.T) {
      f.Fixture()
      f.Assert()
  })

}

func TargetMul2(input float64) float64 {
  return input * 2
}

func TestAssert(t *testing.T) {
  f := fixture.New(t, Fixture2Pow, fixture.P([]fixture.Param{1.0,2.0,3.0}), fixture.R([]fixture.Result{4.0,8.0,16.0}))
  
  f.RunWith(
    func (t *testing.T) {
      pow := f.Value()
      result := TargetMul2(pow.(float64))
      f.AssertGe(float32(result))
  })

}
