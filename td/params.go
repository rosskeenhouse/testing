// Copyright 2020 RHK Development <dev@rosskeen.house>. All rights reserved.
package testdata

import (
  "io/ioutil"
  "path/filepath"
  "log"
  "fmt"
  "github.com/rosskeenhouse/testing/fixture"
)


type Parameters struct {
  *fixture.Parameters
  name string
  basepath string
}

func P(name string) *Parameters { return &Parameters{ name: name } }

// testdata.BasePath('./testdata/fixture_{testname}_')
func (p *Parameters) BasePath() string {
  if p.basepath == "" {
    g := fmt.Sprintf("fixture_%s_", p.name) 
    p.basepath = filepath.Join("./testdata", g)
  }
  return p.basepath
}

// testdata.Paths(base string)
func (p *Parameters) Paths() []string {
  df, e := filepath.Glob(p.BasePath() + "*")
  if e != nil {
    log.Fatal(e)
  }
  return df
}

func (p *Parameters) Values() []fixture.Param {
  // return TestData files
  paths := p.Paths()
  r := make([]fixture.Param, len(paths))
  for i := range(paths) {
    r[i] = paths[i]
  }
  return r
}

func Read(path string) []byte {
  d,e := ioutil.ReadFile(path)
  if e != nil {
    log.Fatalf("Error reading fixture data from %s: %s", path, e) 
  }
  return d
}


