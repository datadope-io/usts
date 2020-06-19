package usts

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/es"
)

// No calendar - Monday => Friday

func ExampleTimeWindow_1() {
	var slot *TimeSlot
	var ts *USTimeSerie
	var err error
	var loc *time.Location

	win := NewTimeWindow("mon-to-fri")
	loc, err = win.SetTimeZone("Europe/Madrid")
	if err != nil {
		fmt.Printf("ERROR on set timeZone %s\n", err)
		return
	}
	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)
	//24x7
	slot, err = NewTimeSlot("24x5", "00 00 * * MON-FRI", "00 00 * * SAT")

	//Laborables
	//	slot1, err = NewTimeSlot("00 09 * * MON-FRI,"00 14 * * MON-FRI")
	//	slot2, err = NewTimeSlot("00 15 * * MON-FRI,"00 18 * * MON-FRI")
	win.AddSlot(slot, Add)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-27 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true

}

//Calender ES (only holidays) with no slots
func ExampleTimeWindow_2() {

	var ts *USTimeSerie
	var err error
	var loc *time.Location

	var calEspain = &cal.Calendar{}
	for _, h := range es.Holidays {
		calEspain.AddHoliday(h)
	}

	win := NewTimeWindow("calendar_spain")
	loc, err = win.SetTimeZone("Europe/Madrid")
	if err != nil {
		fmt.Printf("ERROR on set timeZone %s\n", err)
		return
	}
	win.SetCalendar(calEspain)
	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 2020-04-27 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [5] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [8] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true

}

//Calender ES (only holidays) with no slots
func ExampleTimeWindow_3() {

	var ts *USTimeSerie
	var err error
	var loc *time.Location

	var calEspain = &cal.Calendar{}
	for _, h := range es.Holidays {
		calEspain.AddHoliday(h)
	}

	win := NewTimeWindow("calendar_spain")
	loc, err = win.SetTimeZone("Europe/Madrid")
	if err != nil {
		fmt.Printf("ERROR on set timeZone %s\n", err)
		return
	}
	win.SetCalendar(calEspain)
	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true

}

//Calender ES (only holidays) with no slots
func ExampleTimeWindow_4() {

	var ts *USTimeSerie
	var err error
	var loc *time.Location

	var calEspain = &cal.Calendar{}
	for _, h := range es.Holidays {
		calEspain.AddHoliday(h)
	}

	win := NewTimeWindow("calendar_spain")
	loc, err = win.SetTimeZone("Europe/Madrid")
	if err != nil {
		fmt.Printf("ERROR on set timeZone %s\n", err)
		return
	}
	win.SetCalendar(calEspain)
	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: true
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true

}

func ExampleTimeWindow_5() {
	var slot *TimeSlot
	var ts, ts1 *USTimeSerie
	var err error
	var loc *time.Location

	win := NewTimeWindow("mon-to-fri-2_5_am")
	loc, err = win.SetTimeZone("Europe/Madrid")
	if err != nil {
		fmt.Printf("ERROR on set timeZone %s\n", err)
		return
	}
	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	slot, err = NewTimeSlot("mon-fri", "00 00 * * MON", "00 00 * * SAT")
	win.AddSlot(slot, Add)

	ts1, _ = slot.GetTimeEvents(t0, t1, "Europe/Madrid")
	ts1.DumpInTimezone("Europe/Madrid")
	//weekends
	slot, err = NewTimeSlot("2to5am", "00 02 * * *", "00 05 * * * ")
	win.AddSlot(slot, Remove)
	ts1, _ = slot.GetTimeEvents(t0, t1, "Europe/Madrid")
	ts1.DumpInTimezone("Europe/Madrid")
	//mondays from 02:00 to 03:00

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: false
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [2] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [5] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [6] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [7] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true

}
