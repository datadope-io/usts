package usts

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
)

type AddSlotMode int

const (
	Add    AddSlotMode = 0
	Remove AddSlotMode = 1
)

type TimeWindow struct {
	ID       string
	Slots    []*TimeSlot
	SlotMode []AddSlotMode
	Calendar *cal.Calendar
	timeZone string
}

func NewTimeWindow(id string) *TimeWindow {
	return &TimeWindow{ID: id}

}

func (tw *TimeWindow) SetCalendar(c *cal.Calendar) {
	tw.Calendar = c
}

func (tw *TimeWindow) SetTimeZone(tz string) (*time.Location, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	tw.timeZone = tz
	return loc, nil
}

func (tw *TimeWindow) AddSlot(sl *TimeSlot, mode AddSlotMode) {
	tw.Slots = append(tw.Slots, sl)
	tw.SlotMode = append(tw.SlotMode, mode)

}

func CalendarWindowEvents(c *cal.Calendar, tz string, start, end time.Time) (*USTimeSerie, error) {

	if c == nil {
		return nil, fmt.Errorf("Error Calendar not set ")
	}

	cron_tpl := "00 00 * * *"

	if len(tz) > 0 {
		cron_tpl = "CRON_TZ=" + tz + " 00 00 * * * "
	}
	ilog.Debugf("CRON START: %s ", cron_tpl)
	sched, err := crparser.Parse(cron_tpl)
	if err != nil {
		return nil, fmt.Errorf("ERROR on parse Start cron expression : %s", err)
	}

	ret := NewUSTimeSerie(0)

	yesterday := sched.Next(start.AddDate(0, 0, -1))
	a, b, _ := c.IsHoliday(yesterday)

	ret.SetDefault(!(a || b))

	tnext := start.Add(-1 * time.Second)
	for {
		t := sched.Next(tnext)
		if t.After(end) {
			break
		}
		a, b, _ := c.IsHoliday(t)
		ret.Add(t, !(a || b))
		tnext = t
	}
	return ret, nil
}

func (tw *TimeWindow) GetTimeEvents(start, end time.Time) (*USTimeSerie, error) {

	var ret *USTimeSerie
	var err error
	//first events from calendar
	if tw.Calendar != nil {
		ret, err = CalendarWindowEvents(tw.Calendar, tw.timeZone, start, end)
		if err != nil {
			return nil, err
		}
	} else {
		ret = NewUSTimeSerie(0)
		ret.SetDefault(false)
	}
	// for each slot add/remove slots from

	for k, v := range tw.Slots {
		slotEvents, err := v.GetTimeEvents(start, end, tw.timeZone)
		if err != nil {
			return ret, fmt.Errorf("Error on get time events for slot %d/%s", k, v.ID)
		}
		//fmt.Printf("-------%s/%d/%s---", tw.ID, k, v.ID)
		//slotEvents.DumpInTimezone(tw.timeZone)
		switch tw.SlotMode[k] {
		case Add:
			ret, err = ret.Or(slotEvents)
			if err != nil {
				return ret, fmt.Errorf("Error on Window [%s] can not Add slot  %d/[%s] : Err: %s", tw.ID, k, v.ID, err)
			}
		case Remove:

			slotEventsNeg, _ := slotEvents.Not()
			//fmt.Printf("-------%s/%d/NOT[%s]---", tw.ID, k, v.ID)
			//slotEventsNeg.DumpInTimezone(tw.timeZone)
			ret, err = ret.And(slotEventsNeg)
			if err != nil {
				return ret, fmt.Errorf("Error on Window [%s] can not Remove slot  %d/[%s] : Err: %s", tw.ID, k, v.ID, err)
			}
		}
		//fmt.Printf("-------%s/%d/ CURRENT RET---", tw.ID, k)
		//ret.DumpInTimezone(tw.timeZone)
	}

	return ret, nil
}
