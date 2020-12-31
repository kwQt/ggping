package ggping_test

import (
	"github.com/kwQt/ggping"
	"testing"
)

func TestStatus_GetAll(t *testing.T) {
	t.Run("Return all values (all fit in capacity)", func(t *testing.T) {
		// setup
		status := ggping.NewStatus(3)

		status.Update(1)
		status.Update(2)

		expect := []float64{1, 2}

		// test & verify
		for i, v := range status.GetAll() {
			if v != expect[i] {
				t.Errorf("expect: %g, actual: %g", expect[i], v)
			}
		}
	})

	t.Run("Return all values without ones out of capacity", func(t *testing.T) {
		// setup
		status := ggping.NewStatus(3)

		status.Update(1)
		status.Update(2)
		status.Update(3)
		status.Update(4)

		expect := []float64{2, 3, 4}

		// test & verify
		for i, v := range status.GetAll() {
			if v != expect[i] {
				t.Errorf("expect: %g, actual: %g", expect[i], v)
			}
		}
	})
}

func TestStatus_GetMax(t *testing.T) {
	t.Run("Return max value (all fit in capacity)", func(t *testing.T) {
		// setup
		status := ggping.NewStatus(3)

		status.Update(1)
		status.Update(2)

		var expect float64 = 2

		// test & verify
		if status.GetMax() != expect {
			t.Errorf("expect: %g, actual: %g", expect, status.GetMax())
		}
	})

	t.Run("Return max value ignoring ones out of capacity", func(t *testing.T) {
		// setup
		status := ggping.NewStatus(3)

		status.Update(5)
		status.Update(4)
		status.Update(3)
		status.Update(2)
		status.Update(1)

		var expect float64 = 3

		// test & verify
		if status.GetMax() != expect {
			t.Errorf("expect: %g, actual: %g", expect, status.GetMax())
		}
	})
}
