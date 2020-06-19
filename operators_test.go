package usts

import (
	"fmt"
	"time"
)

func ExampleAnd1() {

	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	ts2.Add(time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.And(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: false
	// [1] TIME: 1981-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1984-01-01 00:00:00 +0000 UTC | VALUE: true
	// [3] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [4] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [5] TIME: 1992-01-01 00:00:00 +0000 UTC | VALUE: true
	// [6] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false

}

func ExampleAnd2() {

	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	ts2.Add(time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.And(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: false
	// [1] TIME: 1981-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [3] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [4] TIME: 1992-01-01 00:00:00 +0000 UTC | VALUE: true
	// [5] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false
}

// ts1(4) And ts2(0)
func ExampleAnd3() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	result, err := ts2.And(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: false
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: false
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false
}

// ts1(0) And ts2(4)
func ExampleAnd4() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(false)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(true)

	ts2.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.And(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: false
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: false
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false
}

func ExampleAnd5() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	result, err := ts2.And(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: false
}

func ExampleOr1() {

	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	ts2.Add(time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1984, 10, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.Or(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: true
	// [1] TIME: 1981-01-01 00:00:00 +0000 UTC | VALUE: true
	// [2] TIME: 1984-01-01 00:00:00 +0000 UTC | VALUE: true
	// [3] TIME: 1984-10-01 00:00:00 +0000 UTC | VALUE: true
	// [4] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [5] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [6] TIME: 1992-01-01 00:00:00 +0000 UTC | VALUE: true
	// [7] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: true

}

func ExampleOr2() {

	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	ts2.Add(time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.Or(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: true
	// [1] TIME: 1981-01-01 00:00:00 +0000 UTC | VALUE: true
	// [2] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [3] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [4] TIME: 1992-01-01 00:00:00 +0000 UTC | VALUE: true
	// [5] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: true
}

// ts1(4) Or ts2(0)
func ExampleOr3() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts1.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts1.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts1.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	result, err := ts2.Or(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: true
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false
}

// ts1(0) And ts2(4)
func ExampleOr4() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(false)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(true)

	ts2.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), true)
	ts2.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), false)
	ts2.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), true)

	result, err := ts2.Or(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: true
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: true
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: false
}

func ExampleOr5() {
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	result, err := ts2.Or(ts1)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		result.Dump()
	}

	// Output:
	// [INIT] Default VALUE: true
}

func ExampleAnd_Not() {

	tz := "Europe/Madrid"
	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Errorf("Error on load lcoation %s", err)
		return
	}
	ts1 := NewUSTimeSerie(0)
	ts1.SetDefault(false)

	ts1.Add(time.Date(2020, 4, 28, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 4, 29, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 4, 30, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 5, 1, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 5, 2, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 5, 3, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 5, 4, 0, 0, 0, 0, loc), true)
	ts1.Add(time.Date(2020, 5, 5, 0, 0, 0, 0, loc), true)

	ts2 := NewUSTimeSerie(0)
	ts2.SetDefault(false)

	ts2.Add(time.Date(2020, 5, 2, 0, 0, 0, 0, loc), true)
	ts2.Add(time.Date(2020, 5, 3, 0, 0, 0, 0, loc), true)
	ts2.Add(time.Date(2020, 5, 4, 0, 0, 0, 0, loc), false)

	ts3, _ := ts2.Not()
	fmt.Println("--TS1--")
	ts1.DumpInTimezone(tz)
	fmt.Println("--TS2--")
	ts2.DumpInTimezone(tz)
	fmt.Println("-- NOT TS2--")
	ts3.DumpInTimezone(tz)
	fmt.Println("-- TS1 AND (NOT TS2)--")
	ts4, _ := ts1.And(ts3)
	ts4.DumpInTimezone(tz)
	ts4.Compact()
	fmt.Println("-- TS1 AND (NOT TS2) COMPACTED --")
	ts4.DumpInTimezone(tz)
	// Output:
	// --TS1--
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
	// --TS2--
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: false
	// -- NOT TS2--
	// [INIT] Default VALUE: true
	// [0] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [1] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [2] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// -- TS1 AND (NOT TS2)--
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
	// -- TS1 AND (NOT TS2) COMPACTED --
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [2] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true

}
