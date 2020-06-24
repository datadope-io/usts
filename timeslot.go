package usts

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

var crparser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

type TimeSlot struct {
	ID string
	// Month    string
	// Dom      string //DayOfMonth
	// Dow      string //DayOfWeek
	// InitHour string
	// InitMin  string
	// EndHour  string
	// EndMin   string

	//tzName string
	//-------------------------------------
	scf string        //Start Cron Format
	ecf string        //End Cron Format
	scs cron.Schedule //Start Cron Schedule
	ecs cron.Schedule //End Cron Schedule
}

func NewTimeSlot(id, startexpr, endexpr string) (*TimeSlot, error) {

	var err error

	ret := &TimeSlot{}
	ret.ID = id
	ret.scf = startexpr
	ret.ecf = endexpr

	//START
	ret.scs, err = crparser.Parse(ret.scf)
	if err != nil {
		return nil, fmt.Errorf("ERROR on parse Start cron expression : %s", err)
	}
	//END
	ret.ecs, err = crparser.Parse(ret.ecf)
	if err != nil {
		return nil, fmt.Errorf("ERROR on parse End cron expression : %s", err)
	}

	return ret, nil
}

// func NewTimeSlot(m, md, wd, ih, eh string) (*TimeSlot, error) {

// 	ret := &TimeSlot{}

// 	ret.ID = fmt.Sprintf("%s - %s - %s - %s/%s", m, md, wd, ih, eh)
// 	ret.Month = m
// 	ret.Dom = md
// 	ret.Dow = wd

// 	re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)

// 	if !re.MatchString(ih) {
// 		return nil, fmt.Errorf("Initial Hour [%s] error format ", ih)
// 	}

// 	var data []string
// 	data = strings.Split(ih, ":")
// 	ret.InitHour = data[0]
// 	ret.InitMin = data[1]

// 	if !re.MatchString(ih) {
// 		return nil, fmt.Errorf("Initial Hour [%s] error format ", ih)
// 	}
// 	data = strings.Split(eh, ":")
// 	ret.EndHour = data[0]
// 	ret.EndMin = data[1]

// 	return ret, nil
// }

func (ts *TimeSlot) RefreshCronTZ(tz string) error {
	var err error

	cron_tpl_start := ts.scf
	cron_tpl_end := ts.ecf

	if len(tz) > 0 {
		cron_tpl_start = "CRON_TZ=" + tz + " " + ts.scf
		cron_tpl_end = "CRON_TZ=" + tz + " " + ts.ecf
	}

	//START
	ts.scs, err = crparser.Parse(cron_tpl_start)
	if err != nil {
		return fmt.Errorf("ERROR on parse Start cron expression : %s", err)
	}
	//END
	ts.ecs, err = crparser.Parse(cron_tpl_end)
	if err != nil {
		return fmt.Errorf("ERROR on parse End cron expression : %s", err)
	}
	return nil
}

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
		ilog.Debugf(">>>>USTS DEBUG:GetPrevious interval %s", i)
		before := start.Add(-i)
		count := 0
		for {
			ilog.Debugf(">>>>USTS DEBUG:Test bfore %s/%d", before, count)
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
	tEnd := GetPreviousCronTime(ts.ecs, start)
	switch {
	case tStart.After(tEnd):
		return tStart, true
	case tStart.Before(tEnd):
		return tEnd, false
	case tStart.Equal(tEnd):
		return tStart, true

	}
	ilog.Warnf(">>>>USTS WARN:Warn:not selected previous event")
	return start, true
}

func (ts *TimeSlot) GetTimeEvents(start, end time.Time, tz string) (*USTimeSerie, error) {

	err := ts.RefreshCronTZ(tz)
	if err != nil {
		return nil, err
	}

	tok := NewUSTimeSerie(0)
	//tok.SetDefault(false)

	tnok := NewUSTimeSerie(0)
	//tnok.SetDefault(true)

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
	ilog.Debugf(">>>>USTS DEBUG:Set default value for t = %s in %t", tinit, state)
	tret.SetDefault(state)
	return tret, err
}
