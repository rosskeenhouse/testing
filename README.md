# Test Fixtures in Go

The purpose of this package is to provide a simple tool for creating test fixtures for go tests.

## What is a fixture?

A test fixture is generally used as a mechanism to provide consistent, repeatable state when testing software.

## Adding Fixtures

A fixture can be as simple as a function that returns a static value or it can leverage parameters to run the same test with different values.  You can also provide a matching result set to make it easier to verify the expected results.

```
// A fixture returns some data, often based on some input parameter
func FixtureExample(p fixture.Param) interface{} {
  return fmt.Sprintf("data %s", p)
}

func TestExample(t *testing.T) {
  // you may define a result set for verification
  r := fixture.R([]fixture.Result{"I do not like your data here","I do not like your data there","I do not like your data anywhere"})

  // creating a parameterized fixture will cause the test to be executed for parameter
  p := fixture.P([]fixture.Param{"here","there","anywhere"})

  f := fixture.New(t, FixtureExample, p, r)

  // The test body will be run  3 times (once for each parameter)
  f.RunWith(
    func (t *testing.T) {
      fv := f.Fixture()
      test_result := fmt.Sprintf("I do not like your %s", fv)
      // Verify the test a result that matches the result set
      f.AssertEqual(test_result)
  })

}

```
