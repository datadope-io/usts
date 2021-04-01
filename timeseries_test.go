package usts

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGetIndex(t *testing.T) {

	var i int
	var ok bool

	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1995")
	ts.Add(time.Date(1997, 1, 1, 0, 0, 0, 0, time.UTC), "1997")
	ts.Add(time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC), "1997")

	i, ok = ts.getIndex(time.Date(1993, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 0 && ok == false {
		t.Errorf("Get Index  got: %d/%t want: 0/false", i, ok)
	}

	i, ok = ts.getIndex(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 0 && ok == true {
		t.Errorf("Get Index  got: %d/%t want: 0/true", i, ok)
	}

	i, ok = ts.getIndex(time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 1 && ok == false {
		t.Errorf("Get Index  got: %d/%t want: 1/false", i, ok)
	}

	i, ok = ts.getIndex(time.Date(1997, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 1 && ok == true {
		t.Errorf("Get Index  got: %d/%t want: 1/true", i, ok)
	}

	i, ok = ts.getIndex(time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 2 && ok == false {
		t.Errorf("Get Index  got: %d/%t want: 2/false", i, ok)
	}

	i, ok = ts.getIndex(time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 2 && ok == true {
		t.Errorf("Get Index  got: %d/%t want: 2/true", i, ok)
	}

	i, ok = ts.getIndex(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	if i != 3 && ok == false {
		t.Errorf("Get Index  got: %d/%t want: 3/false", i, ok)
	}

}

func TestAdd(t *testing.T) {
	var tm time.Time
	var i int
	var ok bool

	ts := NewUSTimeSerie(0)

	// Insert record 0

	tm = time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC)
	i, ok = ts.Add(tm, "1995")
	if i != 0 || ok != false {
		t.Errorf("Insert Record in Time %s Got Index %d/%t want 0/false", tm, i, ok)
	}

	// Test record last
	tm = time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	i, ok = ts.Add(tm, "2002")
	if i != 1 || ok != false {
		t.Errorf("Insert Record in Time %s Got Index %d/%t want 1/false", tm, i, ok)
	}

	tm = time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
	i, ok = ts.Add(tm, "1992")
	if i != 0 || ok != false {
		t.Errorf("Insert Record in Time %s Got Index %d/%t want 0/false", tm, i, ok)
	}

	tm = time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC)
	i, ok = ts.Add(tm, "1996")
	if i != 2 || ok != false {
		t.Errorf("Insert Record in Time %s Got Index %d/%t want 0/false", tm, i, ok)
	}

	tm = time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC)
	i, ok = ts.Add(tm, "1996-2")
	if i != 2 || ok != true {
		t.Errorf("Insert Record in Time %s Got Index %d/%t want 0/true", tm, i, ok)
	}

}

func TestGet(t *testing.T) {
	var v interface{}
	var ok bool

	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), "1973")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC), "1991")
	ts.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC), "2002")
	ts.Add(time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), "2008")
	ts.Add(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC), "2012")
	ts.Add(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC), "2014")
	ts.Add(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), "2018")

	// ------GET STATE ON 1970/2005/2011/2013----------------
	t.Log("Get year 1970 data --> nil/false")
	v, ok = ts.Get(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	if v != nil || ok != false {
		t.Errorf("Get time data Got %v/%t want nil/false", v, ok)
	}

	t.Log("Get year 2000 data --> 1995/false")
	v, ok = ts.Get(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	if v != "1995" || ok != false {
		t.Errorf("Get time data Got %v/%t want 1995/false", v, ok)
	}

	t.Log("Get year 2012 data --> 2012/true")
	v, ok = ts.Get(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC))
	if v != "2012" || ok != true {
		t.Errorf("Get time data Got %v/%t want 2012/true", v, ok)
	}

	t.Log("Get year 2005 data --> 2002/false")
	v, ok = ts.Get(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC))
	if v != "2002" || ok != false {
		t.Errorf("Get time data Got %v/%t want 2002/false", v, ok)
	}

	t.Log("Get year 2013 data --> 2012/false")
	v, ok = ts.Get(time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC))
	if v != "2012" || ok != false {
		t.Errorf("Get time data Got %v/%t want 2012/false", v, ok)
	}

	t.Log("Get Current data --> 2018/false")
	v, ok = ts.Get(time.Now())
	if v != "2018" || ok != false {
		t.Errorf("Get time data Got %v/%t want 2018/false", v, ok)
	}

}

func TestCompact(t *testing.T) {
	var ok bool
	var num int

	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC), "A") //<-
	ts.Add(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC), "B") //<-
	ts.Add(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC), "B")
	ts.Add(time.Date(2004, 1, 1, 0, 0, 0, 0, time.UTC), "C") //<-
	ts.Add(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC), "C")
	ts.Add(time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC), "C")
	ts.Add(time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC), "C")
	ts.Add(time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), "B") //<-
	ts.Add(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC), "C") //<-
	ts.Add(time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), "B") //<-
	ts.Add(time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC), "C") //<-
	ts.Add(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC), "C")
	ts.Add(time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC), "A") //<-
	ts.Add(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC), "C") //<-
	ts.Add(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), "C")

	ok, num = ts.Compact()
	if ok != true || num != 6 {
		t.Errorf("Compact  Got %t/%d  want true/6", ok, num)
	}
	ts.Dump()
	ok, num = ts.Compact()
	if ok != false || num != 0 {
		t.Errorf("Compact  Got %t/%d  want false/0", ok, num)
	}
	len := ts.Len()
	if len != 9 {
		t.Errorf("Compact  Len Got %d  want 9", len)
	}
}

func ExampleCompactBoundariesBoolUSTS() {
	var ok bool
	var num int

	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), false) // 0
	ts.Add(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC), false) // 0
	ts.Add(time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 1
	ts.Add(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 2
	ts.Add(time.Date(2004, 1, 1, 0, 0, 0, 0, time.UTC), false) // 3
	ts.Add(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC), false) // 4
	ts.Add(time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC), false) // 5
	ts.Add(time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC), false) // 6
	ts.Add(time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 7
	ts.Add(time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 8
	ts.Add(time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 9
	ts.Add(time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC), false) // 10
	ts.Add(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC), false) // 11
	ts.Add(time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC), false) // 12
	ts.Add(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC), true)  // 13
	ts.Add(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), false) // 14

	ok, num = ts.CompactBoundaries(false)
	fmt.Printf("Encountered: %t , num: %d\n", ok, num)
	ts.Dump()

	// Output:
	// Encountered: true , num: 9
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 2001-01-01 00:00:00 +0000 UTC | VALUE: false
	// [1] TIME: 2002-01-01 00:00:00 +0000 UTC | VALUE: true
	// [2] TIME: 2007-01-01 00:00:00 +0000 UTC | VALUE: false
	// [3] TIME: 2008-01-01 00:00:00 +0000 UTC | VALUE: true
	// [4] TIME: 2013-01-01 00:00:00 +0000 UTC | VALUE: false
	// [5] TIME: 2014-01-01 00:00:00 +0000 UTC | VALUE: true
	// [6] TIME: 2015-01-01 00:00:00 +0000 UTC | VALUE: false

}

func ExampleCompactBoundariesBoolUSTS_MaskExternal() {
	var ok bool
	var num int

	ts := NewUSTimeSerie(0)

	// The fist mask is from 20 to 40
	ts.Add(time.Date(2000, 1, 1, 10, 20, 0, 0, time.UTC), true)
	ts.Add(time.Date(2000, 1, 1, 10, 40, 0, 0, time.UTC), false)

	// The second mask is from 30 to 50
	ts.Add(time.Date(2000, 1, 1, 10, 30, 0, 0, time.UTC), true)
	ts.Add(time.Date(2000, 1, 1, 10, 50, 0, 0, time.UTC), false)

	// As masks, we need to retrieve the maximum extension of it
	//
	// -------|________|------------------
	// -----------|__________|------------

	// Result with only 2 periods marked
	// -------|______________|------------

	ok, num = ts.CompactBoundaries(false)
	fmt.Printf("Encountered: %t , num: %d\n", ok, num)
	ts.Dump()

	// Output:
	// Encountered: true , num: 2
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 2000-01-01 10:20:00 +0000 UTC | VALUE: true
	// [1] TIME: 2000-01-01 10:50:00 +0000 UTC | VALUE: false

}

func ExampleCompactBoundariesBoolUSTS_MaskInternal() {
	var ok bool
	var num int

	ts := NewUSTimeSerie(0)

	// The fist mask is from 20 to 40
	ts.Add(time.Date(2000, 1, 1, 10, 20, 0, 0, time.UTC), true)
	ts.Add(time.Date(2000, 1, 1, 10, 40, 0, 0, time.UTC), false)

	// The second mask is from 30 to 50
	ts.Add(time.Date(2000, 1, 1, 10, 30, 0, 0, time.UTC), true)
	ts.Add(time.Date(2000, 1, 1, 10, 50, 0, 0, time.UTC), false)

	// As masks, we need to retrieve the minimum extension of it
	//
	// -------|________|------------------
	// -----------|__________|------------

	// Result with only 2 periods marked
	// -----------|____|------------

	ok, num = ts.CompactBoundaries(true)
	fmt.Printf("Encountered: %t , num: %d\n", ok, num)
	ts.Dump()

	// Output:
	// Encountered: true , num: 2
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 2000-01-01 10:30:00 +0000 UTC | VALUE: true
	// [1] TIME: 2000-01-01 10:40:00 +0000 UTC | VALUE: false
}

func ExampleDumpBuffer() {
	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1995")
	ts.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), "1980")
	ts.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), "1985")
	ts.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "1990")
	buf := ts.DumpBufferWithPrefix("prefix")

	reader := bufio.NewReader(&buf)
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Process the line here.
		fmt.Printf("%s\n", strings.TrimSpace(line))
		if err != nil {
			break
		}
	}

	// Output:
	//prefix [INIT] Default VALUE: <nil>
	//prefix [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: 1980
	//prefix [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: 1985
	//prefix [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: 1990
	//prefix [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: 1995

}

func ExampleAdd() {
	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1995")
	ts.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), "1980")
	ts.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), "1985")
	ts.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "1990")
	ts.Dump()
	// Output:
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: 1980
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: 1985
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: 1990
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: 1995

}

func ExampleDelete() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(2004, 10, 18, 8, 37, 1, 0, time.UTC), "2004")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), "1973")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2014, 10, 18, 8, 37, 1, 0, time.UTC), "2014")
	ts.Add(time.Date(2010, 10, 18, 8, 37, 1, 0, time.UTC), "2010")
	ts.Add(time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC), "2008")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012-2")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")
	ts.Delete(time.Date(2014, 10, 18, 8, 37, 1, 0, time.UTC))
	ts.Add(time.Date(2016, 10, 18, 8, 37, 1, 0, time.UTC), "2016")
	ts.Dump()
	// Output:
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 1971-01-01 00:00:00 +0000 UTC | VALUE: 1971
	// [1] TIME: 1973-01-01 00:00:00 +0000 UTC | VALUE: 1973
	// [2] TIME: 1976-01-01 00:00:00 +0000 UTC | VALUE: 1976
	// [3] TIME: 1991-08-25 20:57:08 +0000 UTC | VALUE: 1991
	// [4] TIME: 1992-10-18 08:37:01 +0000 UTC | VALUE: 1992
	// [5] TIME: 1995-10-18 08:37:01 +0000 UTC | VALUE: 1995
	// [6] TIME: 2002-10-18 08:37:01 +0000 UTC | VALUE: 2002
	// [7] TIME: 2004-10-18 08:37:01 +0000 UTC | VALUE: 2004
	// [8] TIME: 2008-04-10 00:00:00 +0000 UTC | VALUE: 2008
	// [9] TIME: 2010-10-18 08:37:01 +0000 UTC | VALUE: 2010
	// [10] TIME: 2012-10-18 08:37:01 +0000 UTC | VALUE: 2012-2
	// [11] TIME: 2016-10-18 08:37:01 +0000 UTC | VALUE: 2016
	// [12] TIME: 2018-10-18 08:37:01 +0000 UTC | VALUE: 2018

}

func ExampleRemoveFromInterval() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(2004, 10, 18, 8, 37, 1, 0, time.UTC), "2004")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), "1973")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2014, 10, 18, 8, 37, 1, 0, time.UTC), "2014")
	ts.Add(time.Date(2010, 10, 18, 8, 37, 1, 0, time.UTC), "2010")
	ts.Add(time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC), "2008")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")

	//remove decade 2000
	ts.RemoveFromInterval(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2009, 12, 31, 23, 59, 59, 0, time.UTC))

	//remove decade 2000
	ts.RemoveFromInterval(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC))
	ts.Dump()

	// Output:
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 1971-01-01 00:00:00 +0000 UTC | VALUE: 1971
	// [1] TIME: 1973-01-01 00:00:00 +0000 UTC | VALUE: 1973
	// [2] TIME: 1976-01-01 00:00:00 +0000 UTC | VALUE: 1976
	// [3] TIME: 2010-10-18 08:37:01 +0000 UTC | VALUE: 2010
	// [4] TIME: 2012-10-18 08:37:01 +0000 UTC | VALUE: 2012
	// [5] TIME: 2014-10-18 08:37:01 +0000 UTC | VALUE: 2014
	// [6] TIME: 2018-10-18 08:37:01 +0000 UTC | VALUE: 2018

}

func ExampleIterateNormal() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")

	start := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()

	err := ts.Iterate(false, start, end, func(t time.Time, val interface{}, i int) bool {
		fmt.Printf("[%d] TIME: %s : VALUE %v\n", i, t, val)
		return true
	})
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	// Output:
	// [2] TIME: 1991-08-25 20:57:08 +0000 UTC : VALUE 1991
	// [3] TIME: 1992-10-18 08:37:01 +0000 UTC : VALUE 1992
	// [4] TIME: 1995-10-18 08:37:01 +0000 UTC : VALUE 1995
	// [5] TIME: 2002-10-18 08:37:01 +0000 UTC : VALUE 2002
	// [6] TIME: 2012-10-18 08:37:01 +0000 UTC : VALUE 2012
	// [7] TIME: 2018-10-18 08:37:01 +0000 UTC : VALUE 2018

}

func ExampleIterateNorma_SubEvents() {
	ts := NewUSTimeSerie(0)
	ts.SetInitialVal(false)

	// Insert example records:
	ts.Add(time.Date(2021, 03, 24, 16, 30, 0, 0, time.UTC), true)
	ts.Add(time.Date(2021, 03, 24, 17, 30, 0, 0, time.UTC), false)

	start := time.Date(2021, 03, 24, 17, 0, 0, 0, time.UTC)
	end := time.Date(2021, 03, 24, 18, 0, 0, 0, time.UTC)

	//end := time.Now()

	err := ts.Iterate(false, start, end, func(t time.Time, val interface{}, i int) bool {
		fmt.Printf("[%d] TIME: %s : VALUE %v\n", i, t, val)
		return true
	})
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	// Output:
	// [1] TIME: 2021-03-24 17:30:00 +0000 UTC : VALUE false

}

func ExampleIterateNorma_SupEvents() {
	ts := NewUSTimeSerie(0)
	ts.SetInitialVal(false)

	// Insert example records:
	ts.Add(time.Date(2021, 03, 24, 17, 30, 0, 0, time.UTC), false)
	ts.Add(time.Date(2021, 03, 24, 18, 30, 0, 0, time.UTC), true)

	start := time.Date(2021, 03, 24, 17, 0, 0, 0, time.UTC)
	end := time.Date(2021, 03, 24, 18, 0, 0, 0, time.UTC)

	//end := time.Now()

	err := ts.Iterate(false, start, end, func(t time.Time, val interface{}, i int) bool {
		fmt.Printf("[%d] TIME: %s : VALUE %v\n", i, t, val)
		return true
	})
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	// Output:
	// [0] TIME: 2021-03-24 17:30:00 +0000 UTC : VALUE false

}

func ExampleIterateReverse() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")

	start := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()

	err := ts.Iterate(true, start, end, func(t time.Time, val interface{}, i int) bool {
		fmt.Printf("[%d] TIME: %s : VALUE %v\n", i, t, val)
		return true
	})
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	// Output:
	// [7] TIME: 2018-10-18 08:37:01 +0000 UTC : VALUE 2018
	// [6] TIME: 2012-10-18 08:37:01 +0000 UTC : VALUE 2012
	// [5] TIME: 2002-10-18 08:37:01 +0000 UTC : VALUE 2002
	// [4] TIME: 1995-10-18 08:37:01 +0000 UTC : VALUE 1995
	// [3] TIME: 1992-10-18 08:37:01 +0000 UTC : VALUE 1992
	// [2] TIME: 1991-08-25 20:57:08 +0000 UTC : VALUE 1991
}

func randates(n int) []time.Time {
	ret := []time.Time{}
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < n; i++ {
		sec := rand.Int63n(delta) + min
		ret = append(ret, time.Unix(sec, 0))
	}
	return ret
}

func BenchmarkAdd(b *testing.B) {

	dates := randates(b.N)
	ts := NewUSTimeSerie(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts.Add(dates[i], i)
	}

}

func TestGetLeftRightIndexInsidePeriod(t *testing.T) {

	var t0, t1 time.Time
	var s, e int
	var err error

	ts := NewUSTimeSerie(0)

	ts.Add(time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), "1995")
	ts.Add(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), "1980")
	ts.Add(time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC), "1985")
	ts.Add(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "1990")

	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: 1980
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: 1985
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: 1990
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: 1995

	//A ----x--[--x----x--]--x----
	t0 = time.Date(1982, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 2 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//A1 ----x--[--x----x]----x----
	t0 = time.Date(1982, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 2 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//A1 ----x----[x----x--]--x------
	t0 = time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 2 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//A1 ----x----[x----x]----x----
	t0 = time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 2 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: 1980
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: 1985
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: 1990
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: 1995

	//B ----x--[--x--]--x----x----

	t0 = time.Date(1982, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1987, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 1 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//B1 ----x----[x--]--x----x----

	t0 = time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1987, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 1 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//B2 ----x----[x]----x----x----

	t0 = time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 1 || err != nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	// [0] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: 1980
	// [1] TIME: 1985-01-01 00:00:00 +0000 UTC | VALUE: 1985
	// [2] TIME: 1990-01-01 00:00:00 +0000 UTC | VALUE: 1990
	// [3] TIME: 1995-01-01 00:00:00 +0000 UTC | VALUE: 1995

	//C0 -[-]--x----x----x----x----

	t0 = time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1977, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != -1 || e != 0 || err == nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want -1/0 ", s, e)
	}

	//C1 ----x----x----x----x--[-]-

	t0 = time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 0 || e != -1 || err == nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 0/-1 ", s, e)
	}

	//            v      v (with error)
	//C2 ----x----x-[--]-x----x----  instead of inside got outside

	t0 = time.Date(1986, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1987, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 1 || e != 2 || err == nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 1/2 ", s, e)
	}

	//                 v      v (with error)
	//C2 ----x----x----x-[--]-x----  instead of inside got outside

	t0 = time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != 2 || e != 3 || err == nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want 2/3 ", s, e)
	}

	//D --[-]---

	ts = NewUSTimeSerie(0)

	t0 = time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 = time.Date(1997, 1, 1, 0, 0, 0, 0, time.UTC)
	s, e, err = ts.getLeftRightIndexInsidePeriod(t0, t1)
	if s != -1 || e != -1 || err == nil {
		t.Errorf("Get Indexes inside period Error Got Index %d/%d want -1/-1 ", s, e)
	}

}

func ExampleBatchDelete() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(2004, 10, 18, 8, 37, 1, 0, time.UTC), "2004")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), "1973")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2014, 10, 18, 8, 37, 1, 0, time.UTC), "2014")
	ts.Add(time.Date(2010, 10, 18, 8, 37, 1, 0, time.UTC), "2010")
	ts.Add(time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC), "2008")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")

	t0 := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	n, err := ts.BatchDelete(t0, t1)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err)
	}
	fmt.Printf("DELETED: %d\n", n)
	ts.Dump()
	// Output:
	// DELETED: 9
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 1971-01-01 00:00:00 +0000 UTC | VALUE: 1971
	// [1] TIME: 1973-01-01 00:00:00 +0000 UTC | VALUE: 1973
	// [2] TIME: 1976-01-01 00:00:00 +0000 UTC | VALUE: 1976
	// [3] TIME: 2018-10-18 08:37:01 +0000 UTC | VALUE: 2018
}

func ExampleBatchDelete2() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(1996, 10, 18, 8, 37, 1, 0, time.UTC), "1996")
	ts.Add(time.Date(1997, 10, 18, 8, 37, 1, 0, time.UTC), "1997")
	ts.Add(time.Date(1998, 10, 18, 8, 37, 1, 0, time.UTC), "1998")
	ts.Add(time.Date(1999, 10, 18, 8, 37, 1, 0, time.UTC), "1999")

	t0 := time.Date(1996, 10, 18, 8, 37, 1, 0, time.UTC)
	t1 := time.Date(1998, 10, 18, 8, 37, 1, 0, time.UTC)

	n, err := ts.BatchDelete(t0, t1)
	if err != nil {
		fmt.Errorf("ERROR: %s\n", err)
	}
	fmt.Printf("DELETED: %d\n", n)
	ts.Dump()
	// Output:
	// DELETED: 3
	// [INIT] Default VALUE: <nil>
	// [0] TIME: 1995-10-18 08:37:01 +0000 UTC | VALUE: 1995
	// [1] TIME: 1999-10-18 08:37:01 +0000 UTC | VALUE: 1999
}

func ExampleSetWindowValue() {
	var t0, t1, tm0, tm1 time.Time

	t0 = time.Date(2042, 2, 1, 6, 45, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 8, 00, 0, 0, time.UTC)
	tm0 = time.Date(2042, 2, 1, 8, 45, 0, 0, time.UTC)
	tm1 = time.Date(2042, 2, 1, 10, 00, 0, 0, time.UTC)

	mask := NewUSTimeSerie(0)
	mask.SetDefault(false)
	mask.SetIntervalValue(tm0, tm1, true)
	mask.SetIntervalValue(t0, t1, true)
	mask.Dump()
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2042-02-01 06:45:00 +0000 UTC | VALUE: true
	// [1] TIME: 2042-02-01 08:00:00 +0000 UTC | VALUE: false
	// [2] TIME: 2042-02-01 08:45:00 +0000 UTC | VALUE: true
	// [3] TIME: 2042-02-01 10:00:00 +0000 UTC | VALUE: false
}

func ExampleSetIntervalValue() {
	ts := NewUSTimeSerie(0)

	// Insert example records:
	ts.Add(time.Date(1995, 10, 18, 8, 37, 1, 0, time.UTC), "1995")
	ts.Add(time.Date(2002, 10, 18, 8, 37, 1, 0, time.UTC), "2002")
	ts.Add(time.Date(1992, 10, 18, 8, 37, 1, 0, time.UTC), "1992")
	ts.Add(time.Date(2004, 10, 18, 8, 37, 1, 0, time.UTC), "2004")
	ts.Add(time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC), "1976")
	ts.Add(time.Date(1973, 1, 1, 0, 0, 0, 0, time.UTC), "1973")
	ts.Add(time.Date(1991, 8, 25, 20, 57, 8, 0, time.UTC), "1991")
	ts.Add(time.Date(2014, 10, 18, 8, 37, 1, 0, time.UTC), "2014")
	ts.Add(time.Date(2010, 10, 18, 8, 37, 1, 0, time.UTC), "2010")
	ts.Add(time.Date(2008, 4, 10, 0, 0, 0, 0, time.UTC), "2008")
	ts.Add(time.Date(2012, 10, 18, 8, 37, 1, 0, time.UTC), "2012")
	ts.Add(time.Date(1971, 1, 1, 0, 0, 0, 0, time.UTC), "1971")
	ts.Add(time.Date(2018, 10, 18, 8, 37, 1, 0, time.UTC), "2018")

	t0 := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	ts.SetDefault("NOOK")
	err := ts.SetIntervalValue(t0, t1, "OK")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	ts.Dump()
	// Output:
	// [INIT] Default VALUE: NOOK
	// [0] TIME: 1971-01-01 00:00:00 +0000 UTC | VALUE: 1971
	// [1] TIME: 1973-01-01 00:00:00 +0000 UTC | VALUE: 1973
	// [2] TIME: 1976-01-01 00:00:00 +0000 UTC | VALUE: 1976
	// [3] TIME: 1980-01-01 00:00:00 +0000 UTC | VALUE: OK
	// [4] TIME: 2015-01-01 00:00:00 +0000 UTC | VALUE: NOOK
	// [5] TIME: 2018-10-18 08:37:01 +0000 UTC | VALUE: 2018
}
