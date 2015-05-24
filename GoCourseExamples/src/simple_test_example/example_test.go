package simple_test_example

import (
	"average_pkg"
	"testing"
	)

func TestAverage(t *testing.T) {
  var v float64
  v = average_pkg.Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}
