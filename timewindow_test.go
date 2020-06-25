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
	//24x5
	slot, err = NewTimeSlot("24x5", "00 00 * * MON-FRI", "00 00 * * SAT")

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

//Calender ES (only holidays) with  24x5 slot
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

	slot, err := NewTimeSlot("24x5", "00 00 * * MON-FRI", "00 00 * * SAT")
	if err != nil {
		fmt.Printf("ERROR on get slot 24x5 %s\n", err)
		return
	}
	win.AddSlot(slot, And)

	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

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
	// [4] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [5] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [6] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [7] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [8] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
}

//Calender ES (only holidays) with  24x5 slot less monday from 3 to 5
func ExampleTimeWindow_4() {

	var ts *USTimeSerie
	var err error
	var loc *time.Location
	var slot *TimeSlot

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

	slot, err = NewTimeSlot("24x5", "00 00 * * MON-FRI", "00 00 * * SAT")
	if err != nil {
		fmt.Printf("ERROR on get slot 24x5 %s\n", err)
		return
	}
	win.AddSlot(slot, And)
	slot, err = NewTimeSlot("mon_3_to_5", "00 03 * * MON", "00 05 * * MON")
	if err != nil {
		fmt.Printf("ERROR on get slot  mon_3_to_5 %s\n", err)
		return
	}
	win.AddSlot(slot, Remove)

	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-27 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-27 03:00:00 +0200 CEST | VALUE: false
	// [2] TIME: 2020-04-27 05:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [7] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [8] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [9] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [10] TIME: 2020-05-04 03:00:00 +0200 CEST | VALUE: false
	// [11] TIME: 2020-05-04 05:00:00 +0200 CEST | VALUE: true
	// [12] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
}

//Calender ES (only holidays) with  24x5 slot less monday from 3 to 5 plus saturday from 14:00 to 18:00
func ExampleTimeWindow_5() {

	var ts *USTimeSerie
	var err error
	var loc *time.Location
	var slot *TimeSlot

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

	slot, err = NewTimeSlot("24x5", "00 00 * * MON-FRI", "00 00 * * SAT")
	if err != nil {
		fmt.Printf("ERROR on get slot 24x5 %s\n", err)
		return
	}
	win.AddSlot(slot, And)
	slot, err = NewTimeSlot("mon_3_to_5", "00 03 * * MON", "00 05 * * MON")
	if err != nil {
		fmt.Printf("ERROR on get slot  mon_3_to_5 %s\n", err)
		return
	}
	win.AddSlot(slot, Remove)

	slot, err = NewTimeSlot("sat_15_to_18", "00 15 * * SAT", "00 18 * * SAT")
	if err != nil {
		fmt.Printf("ERROR on get slot  mon_3_to_5 %s\n", err)
		return
	}
	win.AddSlot(slot, Add)

	t0 := time.Date(2020, 4, 27, 0, 0, 0, 0, loc)
	t1 := time.Date(2020, 5, 5, 0, 0, 0, 0, loc)

	ts, err = win.GetTimeEvents(t0, t1)
	if err != nil {
		fmt.Printf("ERROR on set timewindow %s\n", err)
		return
	}
	ts.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2020-04-27 00:00:00 +0200 CEST | VALUE: true
	// [1] TIME: 2020-04-27 03:00:00 +0200 CEST | VALUE: false
	// [2] TIME: 2020-04-27 05:00:00 +0200 CEST | VALUE: true
	// [3] TIME: 2020-04-28 00:00:00 +0200 CEST | VALUE: true
	// [4] TIME: 2020-04-29 00:00:00 +0200 CEST | VALUE: true
	// [5] TIME: 2020-04-30 00:00:00 +0200 CEST | VALUE: true
	// [6] TIME: 2020-05-01 00:00:00 +0200 CEST | VALUE: false
	// [7] TIME: 2020-05-02 00:00:00 +0200 CEST | VALUE: false
	// [8] TIME: 2020-05-02 15:00:00 +0200 CEST | VALUE: true
	// [9] TIME: 2020-05-02 18:00:00 +0200 CEST | VALUE: false
	// [10] TIME: 2020-05-03 00:00:00 +0200 CEST | VALUE: false
	// [11] TIME: 2020-05-04 00:00:00 +0200 CEST | VALUE: true
	// [12] TIME: 2020-05-04 03:00:00 +0200 CEST | VALUE: false
	// [13] TIME: 2020-05-04 05:00:00 +0200 CEST | VALUE: true
	// [14] TIME: 2020-05-05 00:00:00 +0200 CEST | VALUE: true
}
