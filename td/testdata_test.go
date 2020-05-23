// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package testdata_test

import (
  "testing"
  "github.com/rosskeenhouse/testing/fixture"
  "github.com/rosskeenhouse/testing/td"
)

func TestTDFixture(t *testing.T) {
  r := fixture.R([]fixture.Result{"empty"})
  //"./testdat/fixture_TestTDFixture_empty.yml"})
  f := fixture.New(t, testdata.FixtureRawFile, testdata.P(t.Name()), r)
  
  f.RunWith(
    func (t *testing.T) {
      f.Fixture()
      f.Assert()
  })

}
