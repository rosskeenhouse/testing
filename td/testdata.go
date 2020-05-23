// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package testdata

import (
//  "testing"
//  "io/ioutil"
//  "path/filepath"
//  "log"
//  "fmt"
//  "strings"
//  "strconv"
  "github.com/rosskeenhouse/testing/fixture"
)

// Raw file fixture data

func FixtureRawFile(p fixture.Param) interface{} {
  return Read(p.(string))
}
