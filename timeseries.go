package usts

import (
	"fmt"
	"time"
)

type InterPolateOps int

const (
	Previous InterPolateOps = 0
	Linear   InterPolateOps = 1
)

type USTimeSerie struct {
	defVal interface{}
	t      []time.Time
	m      map[time.Time]interface{}
}

func NewUSTimeSerie(size int) *USTimeSerie {
	uts := &USTimeSerie{}
	uts.m = make(map[time.Time]interface{}, size)
	uts.t = make([]time.Time, 0)
	return uts
}

func (uts *USTimeSerie) Len() int {
	return len(uts.t)
}

func (uts *USTimeSerie) Has(key time.Time) bool {
	_, isPresent := uts.m[key]
	return isPresent
}

func (uts *USTimeSerie) Insert(t time.Time, value interface{}) (int, bool) {
	_, isPresent := uts.m[t]
	if isPresent {
		//overwrite
		uts.m[t] = value
		i, _ := uts.getIndex(t)
		uts.t[i] = t
		return i, true
	}

	uts.m[t] = value
	i, _ := uts.getIndex(t)
	//ilog.Debugf("**********DEBUG TIME %s INDEX %d (found %t)\n",t,i,ok)
	uts.t = append(uts.t[:i], append([]time.Time{t}, uts.t[i:]...)...)
	return i, false

}

//t:          v
//Points: --------[0]-------[1]--------

//t:                    v
//Points: --------[0]-------[1]--------

//t:                             v
//Points: --------[0]-------[1]--------

// getIndex Set Index for Placing new elements
func (uts *USTimeSerie) getIndex(t time.Time) (int, bool) {
	if len(uts.t) == 0 {
		//  ilog.Debugf("------------>INIT\n")
		return 0, false
	}
	for k, v := range uts.t {
		switch {
		case t.After(v):
			ilog.Tracef("[%d]-----[%s]------->AFTER[%s]\n", k, t, v)
		case t.Equal(v):
			ilog.Tracef("[%d]-----[%s]------->EQUAL[%s]\n", k, t, v)
			return k, true
		case t.Before(v):
			ilog.Tracef("[%d]-----[%s]------->BEFORE[%s]\n", k, t, v)
			return k, false
		}
	}
	ilog.Tracef("------------>FINAL\n")
	return len(uts.t), false
}

// Fist return the first time.Time , value, and true/false if found/not found
func (uts *USTimeSerie) First() (time.Time, interface{}, bool) {
	if len(uts.t) > 0 {
		t0 := uts.t[0]
		return t0, uts.m[t0], true
	}
	return time.Time{}, nil, false
}

// Last return the last time.Time , value, and true/false if found/not found
func (uts *USTimeSerie) Last() (time.Time, interface{}, bool) {
	if len(uts.t) > 0 {
		tl := uts.t[len(uts.t)-1]
		return tl, uts.m[tl], true
	}
	return time.Time{}, nil, false
}

// getLeftRightIndexInside get first and end index inside period
func (uts *USTimeSerie) getLeftRightIndexInsidePeriod(start, end time.Time) (int, int, error) {
	//check input values
	if end.Before(start) {
		return 0, 0, fmt.Errorf("End Time (%s) before start (%s)", end, start)
	}
	//check if there is no values in the TimeSerie
	if uts.Len() == 0 {
		return -1, -1, fmt.Errorf("There is no data in this TimeSerie")
	}
	//check if end < first point
	first, _, _ := uts.First()
	if end.Before(first) {
		return -1, 0, fmt.Errorf("End Time (%s) before first value at (%s)", end, first)
	}
	//check if start > last point
	last, _, _ := uts.Last()
	if start.After(last) {
		return 0, -1, fmt.Errorf("Start Time (%s) after last value at (%s)", start, last)
	}
	//Start
	var s, e int
	var t time.Time
	t = start
	for k, v := range uts.t {
		ilog.Tracef(" START[%d]-----[%s]------->[%s]\n", k, t, v)
		if t.Equal(v) {
			s = k
			break
		}
		if t.Before(v) {
			s = k
			break
		}
	}
	//check if next point time is also greater than end time
	// could be if no points between start/end

	ilog.Debugf(">>>>START ---> %d\n", s)
	//End
	t = end
	for k, v := range uts.t {
		ilog.Tracef(" END[%d]-----[%s]------->[%s]\n", k, t, v)
		if t.Equal(v) {
			e = k
			break
		}
		if t.Before(v) {
			e = k - 1
			break
		}
	}
	ilog.Tracef(" S[%d] E[%d]\n", s, e)
	//TODO: this condition is not always true... should be reviewed
	if e < s && s-e == 1 {
		//complete period inside 2 consecutive points
		//return swapped values with error
		ilog.Tracef(" SWAP VALUES now START %d and END %d\n", e, s)
		return e, s, fmt.Errorf("Period inside two consecutive points")
	}
	if e == 0 {
		e = len(uts.t) - 1
	}
	ilog.Debugf(">>>>>END ---> %d\n", e)
	ilog.Debugf("--------->DEBUG START %d / END %d\n", s, e)
	return s, e, nil
}

func (uts *USTimeSerie) Delete(key time.Time) bool {
	_, isPresent := uts.m[key]
	if isPresent {
		delete(uts.m, key)

		i, _ := uts.getIndex(key)
		//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		uts.t = append(uts.t[:i], uts.t[i+1:]...)
		return true
	}
	return false
}

func (uts *USTimeSerie) BatchDelete(start time.Time, end time.Time) (int, error) {

	s, e, err := uts.getLeftRightIndexInsidePeriod(start, end)
	if err != nil {
		if s == 0 && e == 0 {
			return 0, err
		}
		//in any other case there is no points to delete in this interval
		return 0, nil
	}
	ilog.Debugf("DEBUG DELETE index START %d/END %d\n", s, e)
	var retval int
	if e == s {
		return 0, nil
	}
	for i := s; i <= e; i++ {
		retval++
		t := uts.t[i]
		ilog.Debugf(">>>>>>>>>>>>>>>> DEBUG BATCH DELETE index[%d] time[%s] value %v\n", i, t, uts.m[t])
		delete(uts.m, t)
		//uts.t = append(uts.t[:i], uts.t[i+1:]...)
	}
	uts.t = append(uts.t[:s], uts.t[e+1:]...)
	return retval, nil
}

func (uts *USTimeSerie) GetExact(t time.Time) (interface{}, bool) {
	if uts.Has(t) {
		return uts.m[t], true
	}
	return nil, false
}

func (uts *USTimeSerie) GetPrevious(t time.Time) (interface{}, bool) {
	i, _ := uts.getIndex(t)
	if i > 0 {
		return uts.m[uts.t[i-1]], true
	} else {
		//if == 0
		return uts.defVal, false
	}

}

func (uts *USTimeSerie) Keys() []time.Time {
	return uts.t
}

type TimeSortedMapIterFunc func(time.Time, interface{}, int) bool

func (uts *USTimeSerie) Iterate(reversed bool, start, end time.Time, f TimeSortedMapIterFunc) error {

	s, e, err := uts.getLeftRightIndexInsidePeriod(start, end)
	if err != nil {
		return fmt.Errorf("Index [%d/%d]: Error: %s", s, e, err)
	}
	var i int
	//iterate over all indexes start/end included
	if reversed {
		for i = e; i >= s; i-- {
			k := uts.t[i]
			v := uts.m[k]
			result := f(k, v, i)
			if !result {
				break
			}
		}
		return nil
	}

	for i = s; i <= e; i++ {
		k := uts.t[i]
		v := uts.m[k]
		result := f(k, v, i)
		if !result {
			break
		}
	}
	return nil
}

func (uts *USTimeSerie) Dump() {
	fmt.Printf("[INIT] Default VALUE: %v\n", uts.defVal)
	for i := 0; i < len(uts.t); i++ {
		k := uts.t[i]
		v := uts.m[k]
		fmt.Printf("[%d] TIME: %s | VALUE: %v\n", i, k, v)
	}
}

func (uts *USTimeSerie) DumpInTimezone(tz string) {

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("ERROR: bad timezone %s\n", tz)
		return
	}
	fmt.Printf("[INIT] Default VALUE: %v\n", uts.defVal)
	for i := 0; i < len(uts.t); i++ {
		k := uts.t[i]
		v := uts.m[k]
		t := k.In(loc)
		fmt.Printf("[%d] TIME: %s | VALUE: %v\n", i, t, v)
	}
}

func (uts *USTimeSerie) SetDefault(def interface{}) {
	uts.defVal = def
}

func (uts *USTimeSerie) SetInitialVal(val interface{}) {
	uts.defVal = val
}

func (uts *USTimeSerie) NumItems() int {
	return uts.Len()
}

func (uts *USTimeSerie) Remove(t time.Time) bool {
	return uts.Delete(t)
}

func (uts *USTimeSerie) RemoveFromInterval(start time.Time, end time.Time) (int, error) {
	return uts.BatchDelete(start, end)
}

func (uts *USTimeSerie) Add(t time.Time, value interface{}) (int, bool) {
	return uts.Insert(t, value)
}

func (uts *USTimeSerie) Compact() (bool, int) {

	var lastval interface{}

	validkeys := []int{}
	lastval = uts.defVal
	for k, v := range uts.t {
		if lastval != uts.m[v] {
			validkeys = append(validkeys, k)
			lastval = uts.m[v]
		}
	}
	newlen := len(validkeys)
	reduced := uts.Len() - newlen
	if reduced > 0 {
		newmap := make(map[time.Time]interface{}, newlen)
		newarray := make([]time.Time, newlen)
		index := 0
		for _, validkey := range validkeys {
			t := uts.t[validkey]
			val := uts.m[t]
			newarray[index] = t
			newmap[t] = val
			index++
		}
		for k := range uts.m {
			delete(uts.m, k)
		}
		uts.m = newmap
		uts.t = newarray
		return true, reduced
	}
	return false, 0
}

type IteratePeriodFunc func(t0, t1 time.Time, value interface{}) bool

func (uts *USTimeSerie) IterateOnPeriods(start, end time.Time, filter interface{}, f IteratePeriodFunc) error {
	ilog.Debugf("START TIME %s", start)
	ilog.Debugf("END TIME %s", end)
	s, e, err := uts.getLeftRightIndexInsidePeriod(start, end)
	ilog.Debugf("START/END INDEX [%d/%d]", s, e)
	if err != nil {
		switch {
		case s == 0 && e == 0:
			ilog.Debugf(" Distribution Error Case 1 [%d/%d]\n", s, e)
			return err
		case s == -1 && e == -1:
			ilog.Debugf(" Distribution Error Case 2 [%d/%d]\n", s, e)
			fallthrough
		case s == -1 && e == 0: // start/end prior to any data (using defVal)
			ilog.Debugf(" Distribution Error Case 3 [%d/%d]\n", s, e)
			val := uts.defVal
			if filter != nil {
				ilog.Debugf(" Distribution Error Case 3 filter %v / value %v\n", filter, val)
				if val == filter {
					f(start, end, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		case s == 0 && e == -1: // start/end after last data
			ilog.Debugf(" Distribution Error Case 4 [%d/%d]\n", s, e)
			_, val, _ := uts.Last()
			if filter != nil {
				if val == filter {
					f(start, end, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		default:
			ilog.Debugf(" Distribution Error Case 5 [%d/%d]\n", s, e)
			//period inside 2 consecutive values s/e ouside period in this case
			//value from start
			t := uts.t[s]
			val := uts.m[t]
			if filter != nil {
				if val == filter {
					f(start, end, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		}
	}

	if start.Before(uts.t[s]) {
		ilog.Debugf(" Distribution Case 0 [%d/%d]\n", s, e)
		var val interface{}
		if s == 0 {
			val = uts.defVal
		} else {
			tprior := uts.t[s-1]
			val = uts.m[tprior]
		}
		t0 := start
		t1 := uts.t[s]
		if filter != nil {
			if val == filter {
				ret := f(t0, t1, val)
				if !ret {
					return nil
				}
			}
		} else { // filter == nil
			ret := f(t0, t1, val)
			if !ret {
				return nil
			}
		}

	}

	for i := s; i < e; i++ {
		ilog.Debugf(" Distribution Case 1 [%d/%d] interval %d\n", s, e, i)
		t0 := uts.t[i]
		val := uts.m[t0]

		if filter != nil {
			if val != filter {
				continue
			}
		}
		t1 := uts.t[i+1]
		ret := f(t0, t1, val)
		if !ret {
			return nil
		}

	}

	if end.After(uts.t[e]) {
		ilog.Debugf(" Distribution Case 2 [%d/%d]\n", s, e)
		t0 := uts.t[e]
		val := uts.m[t0]
		t1 := end
		if filter != nil {
			if val == filter {
				ret := f(t0, t1, val)
				if !ret {
					return nil
				}
			}
		} else {
			ret := f(t0, t1, val)
			if !ret {
				return nil
			}
		}
	}
	return nil
}

func (uts *USTimeSerie) SetIntervalValue(start, end time.Time, value interface{}) error {
	if uts.defVal == nil {
		return fmt.Errorf("Default value should be configured ")
	}
	//remove all values in this window
	n, err := uts.BatchDelete(start, end)
	if err != nil {
		ilog.Errorf("Error on batch delete from %s/%s: Error: %s", start, end, err)
		return err
	}
	ilog.Debugf("Deleted %d values", n)
	//create 2 entries
	// start with value
	// end with default value
	uts.Insert(start, value)
	uts.Insert(end, uts.defVal)
	return nil
}

func (uts *USTimeSerie) Exist(t time.Time) bool {
	return uts.Has(t)
}

// Get return the value and boolean if value exist or not
func (uts *USTimeSerie) Get(t time.Time) (interface{}, bool) {
	v, ok := uts.GetExact(t)
	if ok {
		return v, ok
	}
	v, _ = uts.GetPrevious(t)
	return v, false
}
