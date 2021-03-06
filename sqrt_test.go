package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal("error")
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}

func TestVariousCasesFromCSV(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("can't open cases file - %s", err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	// Infinite Loop, because we don't know the count of Records in the File.
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("error reading cases file - %s", err)
		}

		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[0])
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[1])
		}

		// In below method, 2 arguments are supplied :-
		// * Name of the testCase.
		// * Anonynous Test method.
		t.Run(fmt.Sprintf("Testcase of SquareRoot for Value : %f", val), func(t *testing.T) {
			out, err := Sqrt(val)
			fmt.Printf("The SQRT value computed by our fn for Value:%v is : %v. \n", val, out)

			if err != nil {
				t.Fatal(err)
			}

			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			} else {
				fmt.Println("TestCase Passed, Expected value :%v from TC Excel File is "+
					" almost Equal to the Value we computed :%v. \n", expected, out)
			}
		})
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
