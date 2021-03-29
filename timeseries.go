package usts

import (
	"bytes"
	"fmt"
	"time"
)

// type InterPolateOps int

// const (
// 	Previous InterPolateOps = 0
// 	Linear   InterPolateOps = 1
// )

// USTimeSerie the Unevently Spaced Time Serie
type USTimeSerie struct {
	defVal interface{}
	t      []time.Time
	m      map[time.Time]interface{}
}

// NewUSTimeSerie creates a new USTimeSerie with and alocates memory for size elements
func NewUSTimeSerie(size int) *USTimeSerie {
	uts := &USTimeSerie{}
	uts.m = make(map[time.Time]interface{}, size)
	uts.t = make([]time.Time, 0)
	return uts
}

// Clone Return a new allocated TimeSeries with the exact data from the origin
func (uts *USTimeSerie) Clone() *USTimeSerie {
	cloned := NewUSTimeSerie(uts.Len())
	for _, v := range uts.t {
		cloned.t = append(cloned.t, v)
	}
	for k, v := range uts.m {
		cloned.m[k] = v
	}
	cloned.defVal = uts.defVal
	return cloned
}

// Len Return the number of elements in the Serie
func (uts *USTimeSerie) Len() int {
	return len(uts.t)
}

// Has return true if one element exist at the exact time
func (uts *USTimeSerie) Has(key time.Time) bool {
	_, isPresent := uts.m[key]
	return isPresent
}

// Insert inserts a new  element on the time serie or updates and existing one in the exact time
// it returns the index at which has been inserted and true if the value has been updated, false if new
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
	//ilog.Debugf(">>>>USTS DEBUG:**********DEBUG TIME %s INDEX %d (found %t)\n",t,i,ok)
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

		return 0, false
	}
	ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getIndex]:->INIT\n")
	for k, v := range uts.t {
		switch {
		case t.After(v):
			ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getIndex]:[%d]--[%s]-->AFTER[%s]\n", k, t, v)
		case t.Equal(v):
			ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getIndex]:[%d]--[%s]-->EQUAL[%s]\n", k, t, v)
			return k, true
		case t.Before(v):
			ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getIndex]:[%d]--[%s]-->BEFORE[%s]\n", k, t, v)
			return k, false
		}
	}
	ilog.Tracef(">>>>USTS TRACE:->FINAL\n")
	return len(uts.t), false
}

// First return the first time.Time , value, and true/false if found/not found
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
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getLeftRightIndexInsidePeriod]: START[%d]-----[%s]------->[%s]\n", k, t, v)
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

	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:getLeftRightIndexInsidePeriod]:>>>>START ---> %d\n", s)
	//End
	t = end
	bf := false
	for k, v := range uts.t {
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getLeftRightIndexInsidePeriod]: END[%d]-----[%s]------->[%s]\n", k, t, v)
		if t.Equal(v) {
			e = k
			break
		}
		if t.Before(v) {
			e = k - 1
			bf = true
			break
		}
	}
	ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getLeftRightIndexInsidePeriod]: S[%d] E[%d]\n", s, e)

	// if the end is the first index, means that no period is equal or before than the provided t
	// in this case, a single event is xpected on s/e
	// TODO: review if it fits all the cases
	if e == 0 && !bf {
		e = len(uts.t) - 1
	}

	//TODO: this condition is not always true... should be reviewed
	if e < s && s-e == 1 {
		//complete period inside 2 consecutive points
		//return swapped values with error
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:getLeftRightIndexInsidePeriod]: Swapping Start/End values now START %d and END %d\n", e, s)
		return e, s, fmt.Errorf("Period inside two consecutive points")
	}

	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:getLeftRightIndexInsidePeriod]: START %d / END %d\n", s, e)
	return s, e, nil
}

// Delete Remove value in time key and return true if value found
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

// BatchDelete removes items inside time series from start to end time both included if they exist
func (uts *USTimeSerie) BatchDelete(start time.Time, end time.Time) (int, error) {

	s, e, err := uts.getLeftRightIndexInsidePeriod(start, end)
	if err != nil {
		if s == 0 && e == 0 {
			return 0, err
		}
		//in any other case there is no points to delete in this interval
		return 0, nil
	}
	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:BatchDelete]: index START %d/END %d\n", s, e)
	var retval int
	if e == s {
		return 0, nil
	}
	for i := s; i <= e; i++ {
		retval++
		t := uts.t[i]
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:BatchDelete]: deleting index[%d] time[%s] value %v\n", i, t, uts.m[t])
		delete(uts.m, t)
		//uts.t = append(uts.t[:i], uts.t[i+1:]...)
	}
	uts.t = append(uts.t[:s], uts.t[e+1:]...)
	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:BatchDelete]: deleted %d values\n", retval)
	return retval, nil
}

// GetExact returns the value in the time t and true if found a value in this exact time, nil/false if not found
func (uts *USTimeSerie) GetExact(t time.Time) (interface{}, bool) {
	if uts.Has(t) {
		return uts.m[t], true
	}
	return nil, false
}

// GetPrevious returns the value in the time t or closest previous value in tim and true if found
// the default value and false will return if no value exist previous time t
func (uts *USTimeSerie) GetPrevious(t time.Time) (interface{}, bool) {
	i, _ := uts.getIndex(t)
	if i > 0 {
		return uts.m[uts.t[i-1]], true
	} else {
		//if == 0
		return uts.defVal, false
	}

}

// Keys return and ordered array of times as TimeSerie Keys
func (uts *USTimeSerie) Keys() []time.Time {
	return uts.t
}

// IterateElementFunc helps users to iterate over the time serie entries
// return true if would like continue the iteration process
type IterateElementFunc func(time.Time, interface{}, int) bool

// Iterate Initialize an Iterator process over timeseries in time reversed order if reversed = true
func (uts *USTimeSerie) Iterate(reversed bool, start, end time.Time, f IterateElementFunc) error {

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

// Dump Prints detailed info on the time series in stdout
func (uts *USTimeSerie) Dump() {
	fmt.Printf("[INIT] Default VALUE: %v\n", uts.defVal)
	for i := 0; i < len(uts.t); i++ {
		k := uts.t[i]
		v := uts.m[k]
		fmt.Printf("[%d] TIME: %s | VALUE: %v\n", i, k, v)
	}
}

// DumpBufferWithPrefix Prints detailed info on the time series in stdout
func (uts *USTimeSerie) DumpBufferWithPrefix(prefix string) bytes.Buffer {
	var buffer bytes.Buffer

	init := fmt.Sprintf(" %s [INIT] Default VALUE: %v\n", prefix, uts.defVal)
	buffer.WriteString(init)
	for i := 0; i < len(uts.t); i++ {
		k := uts.t[i]
		v := uts.m[k]
		l := fmt.Sprintf("%s [%d] TIME: %s | VALUE: %v\n", prefix, i, k, v)
		buffer.WriteString(l)
	}
	return buffer
}

// DumpInTimezone Prints detailed info on the time serie on certain Timezone
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

// SetDefault set a value to use when no more info in the timeseries
// used basically as value tu return when get value from time before
// the first entry or when no entries yet in the timeseries
func (uts *USTimeSerie) SetDefault(def interface{}) {
	uts.defVal = def
}

// SetInitialVal the same behaviour than SetDefault
func (uts *USTimeSerie) SetInitialVal(val interface{}) {
	uts.defVal = val
}

// NumItems return length of the time serie
func (uts *USTimeSerie) NumItems() int {
	return uts.Len()
}

// Remove one element in the time serie ( Synonym of Delete)
func (uts *USTimeSerie) Remove(t time.Time) bool {
	return uts.Delete(t)
}

// RemoveFromInterval removes all items in interval from start to end ( Synonym of BatchDelete )
func (uts *USTimeSerie) RemoveFromInterval(start time.Time, end time.Time) (int, error) {
	return uts.BatchDelete(start, end)
}

// Add insert a new element with time t and value v or updates it if existing yet in the same time t
func (uts *USTimeSerie) Add(t time.Time, v interface{}) (int, bool) {
	return uts.Insert(t, v)
}

// Compact removes all repeated values maintaining only those who changes,
// it returs true and the number of reduced elements if could be compacted
// false and 0 if not
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

// IteratePeriodFunc function helper to iterate over periods instead of iteration over elements
// the initial time t0 and end time t1 of the period will be sent to the funcion , also the value
// at this interval
type IteratePeriodFunc func(t0, t1 time.Time, value interface{}) bool

// IterateOnPeriods helps iterate over periods from start to end instead of iteration over elements.
// it will execute function f only on periods with value == filter and will stop iteration if
// iteration function returns false
func (uts *USTimeSerie) IterateOnPeriods(start, end time.Time, filter interface{}, f IteratePeriodFunc) error {
	s, e, err := uts.getLeftRightIndexInsidePeriod(start, end)
	ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: start/end Time[ %s/%s] indexes found [%d/%d]", start, end, s, e)
	if err != nil {
		switch {
		case s == 0 && e == 0:
			ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 1 [%d/%d]\n", s, e)
			return err
		case s == -1 && e == -1:
			ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 2 [%d/%d]\n", s, e)
			fallthrough
		case s == -1 && e == 0: // start/end prior to any data (using defVal)
			ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 3 [%d/%d]\n", s, e)
			val := uts.defVal
			if filter != nil {
				ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 3 filter %v / value %v\n", filter, val)
				if val == filter {
					f(start, end, val)
				} else {
					ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", start, end, filter, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		case s == 0 && e == -1: // start/end after last data
			ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 4 [%d/%d]\n", s, e)
			_, val, _ := uts.Last()
			if filter != nil {
				if val == filter {
					f(start, end, val)
				} else {
					ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", start, end, filter, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		default:
			ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]: Indexes Error Case 5 [%d/%d]\n", s, e)
			//period inside 2 consecutive values s/e ouside period in this case
			//value from start
			t := uts.t[s]
			val := uts.m[t]
			if filter != nil {
				if val == filter {
					f(start, end, val)
				} else {
					ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", start, end, filter, val)
				}
			} else {
				f(start, end, val)
			}
			return nil
		}
	}

	if start.Before(uts.t[s]) {
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:IterateOnPeriods]: Process Start Interval [%d/%d]\n", s, e)
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
			} else {
				ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", t0, t1, filter, val)
			}
		} else { // filter == nil
			ret := f(t0, t1, val)
			if !ret {
				return nil
			}
		}

	}

	for i := s; i < e; i++ {
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:IterateOnPeriods]: Process Mid Interval %d [%d/%d]\n", i, s, e)
		t0 := uts.t[i]
		t1 := uts.t[i+1]
		val := uts.m[t0]
		if filter != nil {
			if val != filter {
				ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", t0, t1, filter, val)
				continue
			}
		}
		ret := f(t0, t1, val)
		if !ret {
			return nil
		}

	}

	if end.After(uts.t[e]) {
		ilog.Tracef(">>>>USTS TRACE [USTimeSerie:IterateOnPeriods]: Process End Interval  [%d/%d]\n", s, e)
		t0 := uts.t[e]
		val := uts.m[t0]
		t1 := end
		if filter != nil {
			if val == filter {
				ret := f(t0, t1, val)
				if !ret {
					return nil
				}
			} else {
				ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:IterateOnPeriods]:Skipping period [%s/%s] filter: %v | value %v\n", t0, t1, filter, val)
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

// SetIntervalValue force a value in a interval at the end of the  interval
// the value is forced to be the default value. In this operation all points in this interval
// will be removed
func (uts *USTimeSerie) SetIntervalValue(start, end time.Time, value interface{}) error {
	if uts.defVal == nil {
		return fmt.Errorf("Default value should be configured ")
	}
	//remove all values in this window
	n, err := uts.BatchDelete(start, end)
	if err != nil {
		ilog.Errorf(">>>>USTS ERROR:Error on batch delete from %s/%s: Error: %s", start, end, err)
		return err
	}
	ilog.Debugf(">>>>USTS DEBUG:Deleted %d values", n)
	//create 2 entries
	// start with value
	// end with default value
	uts.Insert(start, value)
	uts.Insert(end, uts.defVal)
	return nil
}

// Exist return true if one element exist at the exact time ( Synonym of Has )
func (uts *USTimeSerie) Exist(t time.Time) bool {
	return uts.Has(t)
}

// Get return the value in the exact time or any previous existing and true if value in the exact time
func (uts *USTimeSerie) Get(t time.Time) (interface{}, bool) {
	v, ok := uts.GetExact(t)
	if ok {
		return v, ok
	}
	v, _ = uts.GetPrevious(t)
	return v, false
}
