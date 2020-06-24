package usts

import (
	"fmt"
	"time"
)

func ExampleTimeSlot_1() {
	slot, err := NewTimeSlot("work_hours_1", "00 08 * * *", "00 15 * * * ")
	if err != nil {
		ilog.Errorf("Error: %s", err)
		return
	}
	t0 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC)
	ilog.Debugf("%+v", slot)
	events, err := slot.GetTimeEvents(t0, t1, "Europe/Madrid")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	events.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2000-01-01 08:00:00 +0100 CET | VALUE: true
	// [1] TIME: 2000-01-01 15:00:00 +0100 CET | VALUE: false
	// [2] TIME: 2000-01-02 08:00:00 +0100 CET | VALUE: true
	// [3] TIME: 2000-01-02 15:00:00 +0100 CET | VALUE: false
	// [4] TIME: 2000-01-03 08:00:00 +0100 CET | VALUE: true
	// [5] TIME: 2000-01-03 15:00:00 +0100 CET | VALUE: false
	// [6] TIME: 2000-01-04 08:00:00 +0100 CET | VALUE: true
	// [7] TIME: 2000-01-04 15:00:00 +0100 CET | VALUE: false
	// [8] TIME: 2000-01-05 08:00:00 +0100 CET | VALUE: true
	// [9] TIME: 2000-01-05 15:00:00 +0100 CET | VALUE: false
	// [10] TIME: 2000-01-06 08:00:00 +0100 CET | VALUE: true
	// [11] TIME: 2000-01-06 15:00:00 +0100 CET | VALUE: false

}

func ExampleTimeSlot_2() {
	slot, err := NewTimeSlot("24x7", "00 00 * * *", "00 00 * * *")
	if err != nil {
		ilog.Errorf("Error: %s", err)
		return
	}
	t0 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC)
	ilog.Debugf("%+v", slot)
	events, err := slot.GetTimeEvents(t0, t1, "Europe/Madrid")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	events.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2000-01-02 00:00:00 +0100 CET | VALUE: true
	// [1] TIME: 2000-01-03 00:00:00 +0100 CET | VALUE: true
	// [2] TIME: 2000-01-04 00:00:00 +0100 CET | VALUE: true
	// [3] TIME: 2000-01-05 00:00:00 +0100 CET | VALUE: true
	// [4] TIME: 2000-01-06 00:00:00 +0100 CET | VALUE: true
	// [5] TIME: 2000-01-07 00:00:00 +0100 CET | VALUE: true

}

func ExampleTimeSlot_3() {
	slot, err := NewTimeSlot("24x7_b", "00 00 * * MON", "00 00 * * MON")
	if err != nil {
		ilog.Errorf("Error: %s", err)
		return
	}
	t0 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2000, 1, 7, 0, 0, 0, 0, time.UTC)
	ilog.Debugf("%+v", slot)
	events, err := slot.GetTimeEvents(t0, t1, "Europe/Madrid")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	events.DumpInTimezone("Europe/Madrid")
	// Output:
	// [INIT] Default VALUE: false
	// [0] TIME: 2000-01-01 23:00:00 +0000 UTC | VALUE: true
	// [1] TIME: 2000-01-02 23:00:00 +0000 UTC | VALUE: true
	// [2] TIME: 2000-01-03 23:00:00 +0000 UTC | VALUE: true
	// [3] TIME: 2000-01-04 23:00:00 +0000 UTC | VALUE: true
	// [4] TIME: 2000-01-05 23:00:00 +0000 UTC | VALUE: true
	// [5] TIME: 2000-01-06 23:00:00 +0000 UTC | VALUE: true

}
