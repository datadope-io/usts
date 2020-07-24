package usts

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

var crparser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

// TimeSlot this object give us ability to create periodic time intervals from some logic expresion
type TimeSlot struct {
	ID string

	//-------------------------------------
	Scf string        //Start Cron Format
	Ecf string        //End Cron Format
	scs cron.Schedule //Start Cron Schedule
	ecs cron.Schedule //End Cron Schedule
}

// NewTimeSlot create a periodic time slot from cron based expression for start and end time slot
func NewTimeSlot(id, startexpr, endexpr string) (*TimeSlot, error) {

	var err error

	ret := &TimeSlot{}
	ret.ID = id
	ret.Scf = startexpr
	ret.Ecf = endexpr

	//START
	ret.scs, err = crparser.Parse(ret.Scf)
	if err != nil {
		return nil, fmt.Errorf("ERROR on parse Start cron expression : %s", err)
	}
	//END
	ret.ecs, err = crparser.Parse(ret.Ecf)
	if err != nil {
		return nil, fmt.Errorf("ERROR on parse End cron expression : %s", err)
	}

	return ret, nil
}

// RefreshCronTZ reload cron based expressions adding timezone info
func (ts *TimeSlot) RefreshCronTZ(tz string) error {
	var err error

	cronTplStart := ts.Scf
	cronTplEnd := ts.Ecf

	if len(tz) > 0 {
		cronTplStart = "CRON_TZ=" + tz + " " + ts.Scf
		cronTplEnd = "CRON_TZ=" + tz + " " + ts.Ecf
	}

	//START
	ts.scs, err = crparser.Parse(cronTplStart)
	if err != nil {
		return fmt.Errorf("ERROR on parse Start cron expression : %s", err)
	}
	//END
	ts.ecs, err = crparser.Parse(cronTplEnd)
	if err != nil {
		return fmt.Errorf("ERROR on parse End cron expression : %s", err)
	}
	return nil
}

// GetPreviousCronTime get Previous scheduled time from cron Scheduler
// this method will check if exist any previous sched value beggining in the past
// in order to avoid too expensive cost we will test iteratively with 1h,6h,24h,7d,30d,365d before
func GetPreviousCronTime(sch cron.Schedule, start time.Time) time.Time {

	intervals := []time.Duration{
		time.Hour,            //1h
		6 * time.Hour,        //6h
		24 * time.Hour,       //24h
		7 * 24 * time.Hour,   //7d
		30 * 24 * time.Hour,  //1moth
		365 * 24 * time.Hour, //1year
	}

	for _, i := range intervals {
		ilog.Debugf(">>>>USTS DEBUG [GetPreviousCronTime]: GetPrevious from interval %s", i)
		before := start.Add(-i)
		count := 0
		for {
			ilog.Debugf(">>>>USTS DEBUG [GetPreviousCronTime]:Test before %s/%d", before, count)
			iter := sch.Next(before)
			if iter.After(start) && count > 0 {
				return before
			}
			before = iter
			count++
		}

	}
	return time.Time{}
}

// GetClosestPreviousEvent helps to determine the default/first value
func (ts *TimeSlot) GetClosestPreviousEvent(start time.Time) (time.Time, bool) {

	tStart := GetPreviousCronTime(ts.scs, start)
	ilog.Debugf(">>>>USTS DEBUG [TimeSlot:GetClosestPreviousEvent]: from Start Expression [%s] got on time [%s]", ts.Scf, tStart)
	tEnd := GetPreviousCronTime(ts.ecs, start)
	ilog.Debugf(">>>>USTS DEBUG [TimeSlot:GetClosestPreviousEvent]: from End Expression [%s] got on time [%s]", ts.Ecf, tEnd)
	switch {
	case tStart.After(tEnd):
		ilog.Debugf(">>>>USTS DEBUG [TimeSlot:GetClosestPreviousEvent]: Start after end => (start expr wins)TRUE")
		return tStart, true
	case tStart.Before(tEnd):
		ilog.Debugf(">>>>USTS DEBUG [TimeSlot:GetClosestPreviousEvent]: Start before end => (end expr wins) FALSE")
		return tEnd, false
	case tStart.Equal(tEnd):
		ilog.Debugf(">>>>USTS DEBUG [TimeSlot:GetClosestPreviousEvent]: Start == end => Set to TRUE")
		return tStart, true

	}
	ilog.Warnf(">>>>USTS WARN [TimeSlot:GetClosestPreviousEvent]: can not selected previous event")
	return start, true
}

// GetTimeEvents get all initial(as true)/final(as false) slot events
func (ts *TimeSlot) GetTimeEvents(start, end time.Time, tz string) (*USTimeSerie, error) {

	err := ts.RefreshCronTZ(tz)
	if err != nil {
		return nil, err
	}

	tok := NewUSTimeSerie(0)

	tnok := NewUSTimeSerie(0)

	//1 second below Needed to get start time if match
	// with scheduled events with Next() function
	tnext := start.Add(-1 * time.Second)
	for {
		t := ts.scs.Next(tnext)
		if t.After(end) {
			break
		}
		tok.Add(t, true)
		tnext = t
	}
	tnext = start
	for {
		t := ts.ecs.Next(tnext)
		if t.After(end) {
			break
		}
		tnok.Add(t, false)
		tnext = t
	}
	tret, err := tok.Combine(tnok)
	tinit, state := ts.GetClosestPreviousEvent(start)
	ilog.Debugf(">>>>USTS DEBUG[TimeSlot:GetTimeEvents] Set default value for t = %s in %t", tinit, state)
	tret.SetDefault(state)
	return tret, err
}
