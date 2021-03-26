package usts

import (
	"time"
)

// Distribution computes distribution in time of values from start to end
// and only on these periods where the USTimeSerie mask will be true. False periods won't compute on time distribution
// mask = nil will be equivalent to a true always USTimeSerie mask
// ir returns a interfaced keyed map of durations and the complete computed time ( will be end - start if no mask)
func (uts *USTimeSerie) Distribution(start time.Time, end time.Time, mask *USTimeSerie) (map[interface{}]time.Duration, time.Duration, error) {

	var tdur time.Duration
	ret := make(map[interface{}]time.Duration)
	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:Distribution]: start/end [%s/%s] \n", start, end)

	// if mask is not provided, we setup as default over all available period as true
	if mask == nil {
		mask = NewUSTimeSerie(0)
		mask.SetDefault(true)
		mask.SetIntervalValue(start, end, true)
	}

	//Distribution needs to have defined elements on exact start/end
	//in order to do it without data modification we need to clone this object
	cloned := uts.Clone()
	sVal, oks := uts.Get(start)
	if !oks {
		cloned.Insert(start, sVal)
	}
	sEnd, oke := uts.Get(end)
	if !oke {
		cloned.Insert(end, sEnd)
	}

	// Distribution needs to create a new UST with all available events during time
	// logically: cloned && mask will create C cloned periods + M mask periods
	clonedMarked, err := cloned.MarkPeriods(mask)
	if err != nil {
		return nil, 0, err
	}

	// The logical operation mask && cloned can return nil if mask is not defined
	// in this case, we will use directly the cloned one

	mask.IterateOnPeriods(start, end, true,
		func(tm0, tm1 time.Time, valmask interface{}) bool {
			clonedMarked.IterateOnPeriods(tm0, tm1, nil,
				func(t0, t1 time.Time, val interface{}) bool {
					dur := t1.Sub(t0)
					tdur += dur
					if _, ok := ret[val]; ok {
						ret[val] += dur
					} else {
						ret[val] = dur
					}
					ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:Distribution]: Interval from [%s/%s] Counting value [%v] with Duration %s | Total: %s\n", t0, t1, val, dur, ret[val])
					return true
					//End TimeSerie Period
				})
			//End Mask Period
			return true
		})

	return ret, tdur, nil
}
