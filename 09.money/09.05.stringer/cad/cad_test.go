package cad

import (
	"fmt"
	"testing"
)

var testCasesForParse = []struct {
	s             string
	expectError   bool
	expectedCents int64
}{
	{
		s:             "something",
		expectError:   true,
		expectedCents: 0,
	},
	{
		s:             "12312321",
		expectError:   true,
		expectedCents: 0,
	},
	{
		s:             "-$1234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	// from here
	{
		s:             "$-1234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "-$1,234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "$-1,234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "CAD -$1234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "CAD $-1234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "CAD-$1,234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "CAD$-1,234.56",
		expectError:   false,
		expectedCents: -123456,
	},
	{
		s:             "$1234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "$1,234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "CAD $1234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "CAD $1,234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "CAD$1234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "CAD$1,234.56",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "$0.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "$.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "-$0.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "-$.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "$-0.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "$-.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "CAD $0.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "CAD $.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "CAD -$0.09",
		expectError:   false,
		expectedCents: -9,
	},

	{
		s:             "CAD $-0.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "CAD $-.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "CAD$0.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "CAD$.09",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "CAD-$0.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "CAD-$.09",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "CAD$-0.09",
		expectError:   false,
		expectedCents: -9,
	},

	{
		s:             "9¢",
		expectError:   false,
		expectedCents: 9,
	},
	{
		s:             "-9¢",
		expectError:   false,
		expectedCents: -9,
	},
	{
		s:             "123456¢",
		expectError:   false,
		expectedCents: 123456,
	},
	{
		s:             "-123456¢",
		expectError:   false,
		expectedCents: -123456,
	},
}

func TestParseCAD(t *testing.T) {
	for _, testCase := range testCasesForParse {
		cad, err := ParseCAD(testCase.s)

		if err != nil {
			if testCase.expectError {
				continue
			}
			t.Errorf("unexpected error for parsing %s", testCase.s)
		}

		if cad.AsCents() != testCase.expectedCents {
			t.Errorf("expected to get %d for parsing %s but received %d", testCase.expectedCents, testCase.s, cad.AsCents())
		}
	}
}

func TestAbs(t *testing.T) {
	cad := Cents(-200)

	if cad.Abs().AsCents() != 200 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, 200, cad.Abs().AsCents())
	}
}

func TestAdd(t *testing.T) {
	cad1 := Cents(100)
	cad2 := Cents(-200)

	if cad1.Add(cad2).AsCents() != -100 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, -100, cad1.Add(cad2).AsCents())
	}
}

func TestAsCents(t *testing.T) {
	cad1 := Cents(-100)

	if cad1.AsCents() != -100 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, -100, cad1.AsCents())
	}
}

func TestCanonicalForm(t *testing.T) {
	cad := Cents(-1242)

	dolors, cents := cad.CanonicalForm()

	if dolors != -12 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, -1200, dolors)
	}

	if cents != -42 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, -42, cents)
	}
}

func TestMul(t *testing.T) {
	cad := Cents(-100)

	if cad.Mul(2).AsCents() != -200 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, -200, cad.Mul(2).AsCents())
	}
}

func TestSub(t *testing.T) {
	cad1 := Cents(100)
	cad2 := Cents(-200)

	if cad1.Sub(cad2).AsCents() != 300 {
		errorFormat := "expected %d but received %d"
		t.Errorf(errorFormat, 100, cad1.Sub(cad2).AsCents())
	}
}

func TestGoString(t *testing.T) {
	wallet := Cents(240)

	printed := fmt.Sprintf("%#v", wallet)

	if printed != "Cents(240)" {
		t.Errorf("expected Cents(240) but received %s", printed)
	}
}

func TestString(t *testing.T) {
	wallet := Cents(240)

	printed := fmt.Sprintln(wallet)

	if printed != "CAD$2.40\n" {
		t.Errorf("expected CAD$2.40 but received %s", printed)
	}
}
