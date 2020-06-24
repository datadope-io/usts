package usts

import (
	"time"
)

func (uts *USTimeSerie) Distribution(start time.Time, end time.Time, mask *USTimeSerie) (map[interface{}]time.Duration, time.Duration, error) {

	var tdur time.Duration
	ret := make(map[interface{}]time.Duration)
	ilog.Debugf("####################### DISTRIBUTION ##################\n")

	ilog.Debugf("START %s\n", start)
	ilog.Debugf("END %s\n", end)

	if mask == nil {

		mask = NewUSTimeSerie(0)
		mask.SetDefault(true)

	}

	mask.IterateOnPeriods(start, end, true,
		func(tm0, tm1 time.Time, valmask interface{}) bool {
			uts.IterateOnPeriods(tm0, tm1, nil,
				func(t0, t1 time.Time, val interface{}) bool {
					dur := t1.Sub(t0)
					tdur += dur
					if _, ok := ret[val]; ok {
						ret[val] += dur
					} else {
						ret[val] = dur
					}
					ilog.Debugf("Counting value [%v] Interval  from [%s] to [%s] - Duration %s\n", val, t0, t1, dur)
					return true
					//End TimeSerie Period
				})
			//End Mask Period
			return true
		})

	return ret, tdur, nil
}
