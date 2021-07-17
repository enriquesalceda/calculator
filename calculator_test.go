package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	fn   func(float64, float64) float64
	a, b float64
	test string
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tcs := []testCase{
		{fn: calculator.Add, a: 2, b: 2, want: 4, test: "2 + 2 = 4"},
		{fn: calculator.Add, a: 1, b: 1, want: 2, test: "1 + 1 = 2"},
		{fn: calculator.Add, a: 5, b: 0, want: 5, test: "5 + 0 = 5"},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf(tc.test, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tcs := []testCase{
		{fn: calculator.Substract, a: 4, b: 2, want: 2, test: "4 - 2 = 2"},
		{fn: calculator.Substract, a: 6, b: 1, want: 5, test: "6 - 1 = 5"},
		{fn: calculator.Substract, a: 2, b: 0, want: 2, test: "2 - 0 = 2"},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf(tc.test, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	tcs := []testCase{
		{fn: calculator.Multiply, a: 4, b: 5, want: 20, test: "4 * 5 = 20"},
		{fn: calculator.Multiply, a: 1, b: 9, want: 9, test: "1 * 9 = 9"},
		{fn: calculator.Multiply, a: 6, b: 5, want: 30, test: "6 * 5 = 30"},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf(tc.test, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	tcs := []struct {
		a, b        float64
		want        float64
		errExpected bool
	}{
		{a: 9, b: 3, want: 3, errExpected: false},
		{a: 9, b: 0, want: 1, errExpected: true},
		{a: 1, b: 2, want: 0.5, errExpected: false},
	}

	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%.1f, %.1f): unexpeted error status %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%.1f, %.1f): want %.1f, got %.1f", tc.a, tc.b, tc.want, got)
		}
	}
}
