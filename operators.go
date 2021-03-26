package usts

import (
	"fmt"
	"time"
)

//------------------------------------
//Tools to combine and sort time arrays
// https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html
// https://www.golangprograms.com/remove-duplicate-values-from-slice.html
//
//-------------------------------------

func keyCombineOrdered(ar1 []time.Time, ar2 []time.Time) []time.Time {
	return mergeTimeSort(append(ar1, ar2...))

}

func mergeTimeSort(items []time.Time) []time.Time {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]time.Time, middle)
		right = make([]time.Time, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return mergeTime(mergeTimeSort(left), mergeTimeSort(right))
}

func mergeTime(left, right []time.Time) (result []time.Time) {
	result = make([]time.Time, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].Before(right[0]) {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func unique(in []time.Time) []time.Time {
	keys := make(map[time.Time]bool)
	out := []time.Time{}
	for _, entry := range in {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			out = append(out, entry)
		}
	}
	return out
}

// And Logical AND from this TimeSerie with the other TimeSerie the resulting
// TimeSerie will have the same time points in both previous series (max will be N+M if any repeated time)
// computing logical AND : this is `output[t] = x[t1] AND y[t1]` if no value in y[t1] it gets the previous value
func (uts *USTimeSerie) And(other *USTimeSerie) (*USTimeSerie, error) {

	var res *USTimeSerie

	if other.defVal == nil {
		return nil, fmt.Errorf("new USTimeSerie has no default value")
	}
	if uts.defVal == nil {
		return nil, fmt.Errorf("this USTimeSerie has no default value")
	}

	if uts.Len() == 0 && other.Len() == 0 {
		res = NewUSTimeSerie(0)
		v1 := uts.defVal
		v2 := other.defVal
		result := (v1.(bool) && v2.(bool))
		res.SetDefault(result)
		return res, nil
	}

	combined := unique(keyCombineOrdered(uts.Keys(), other.Keys()))

	res = NewUSTimeSerie(len(combined))
	v1 := uts.defVal
	v2 := other.defVal
	result := (v1.(bool) && v2.(bool))
	res.SetDefault(result)

	for _, t := range combined {

		v1, _ := uts.Get(t)
		v2, _ := other.Get(t)
		ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:And]:computing value for time %s [ %+v AND %+v ]", t, v1, v2)
		result := (v1.(bool) && v2.(bool))
		res.Add(t, result)

	}
	return res, nil
}

// Mark periods on this TimeSerie with the other TimeSerie the resulting
// TimeSerie will have the same time points of this TimeSeries (max will be N+M if any repeated time)
// with all available periods of this TimeSerie and the other TimeSerie (max will be N+M periods)
// marking periods: this is `output[t] = x[t1..tM]` with M all available time periods combinations,
// if x[tx] has no value, it will retrieve the previous
func (uts *USTimeSerie) MarkPeriods(other *USTimeSerie) (*USTimeSerie, error) {

	var res *USTimeSerie

	if uts.Len() == 0 {
		res = NewUSTimeSerie(0)
		res.SetDefault(uts.defVal)
		return res, nil
	}

	combined := unique(keyCombineOrdered(uts.Keys(), other.Keys()))

	res = NewUSTimeSerie(len(combined))
	res.SetDefault(uts.defVal)

	for _, t := range combined {
		v1, _ := uts.Get(t)
		ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:MarkPeriods]:marking periods with value for time %s [ %+v ]", t, v1)
		result := v1
		res.Add(t, result)
	}
	return res, nil
}

// Or Logical OR from this TimeSerie with the other TimeSerie the resulting
// TimeSerie will have the same time points in both previous series (max will be N+M if any repeated time)
// computing logical OR : this is `output[t] = x[t1] OR y[t1]` if no value in y[t1] it gets the previous value
func (uts *USTimeSerie) Or(other *USTimeSerie) (*USTimeSerie, error) {
	var res *USTimeSerie

	if other.defVal == nil {
		return nil, fmt.Errorf("new USTimeSerie has no default value")
	}
	if uts.defVal == nil {
		return nil, fmt.Errorf("this USTimeSerie has no default value")
	}

	if uts.Len() == 0 && other.Len() == 0 {
		res = NewUSTimeSerie(0)
		v1 := uts.defVal
		v2 := other.defVal
		result := (v1.(bool) || v2.(bool))
		res.SetDefault(result)
		return res, nil
	}

	combined := unique(keyCombineOrdered(uts.Keys(), other.Keys()))

	res = NewUSTimeSerie(len(combined))
	v1 := uts.defVal
	v2 := other.defVal
	result := (v1.(bool) || v2.(bool))
	res.SetDefault(result)

	for _, t := range combined {

		v1, _ := uts.Get(t)
		v2, _ := other.Get(t)
		ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:Or]:computing value for time %s [ %+v AND %+v ]", t, v1, v2)
		result := (v1.(bool) || v2.(bool))
		res.Add(t, result)

	}
	return res, nil
}

// Xor Logical XOR from this TimeSerie with the other TimeSerie the resulting
// TimeSerie will have the same time points in both previous series (max will be N+M if any repeated time)
// computing logical XOR : this is `output[t] = x[t1] XOR y[t1]` if no value in y[t1] it gets the previous value
func (uts *USTimeSerie) Xor(other *USTimeSerie) (*USTimeSerie, error) {
	var res *USTimeSerie

	if other.defVal == nil {
		return nil, fmt.Errorf("new USTimeSerie has no default value")
	}
	if uts.defVal == nil {
		return nil, fmt.Errorf("this USTimeSerie has no default value")
	}

	if uts.Len() == 0 && other.Len() == 0 {
		res = NewUSTimeSerie(0)
		v1 := uts.defVal
		v2 := other.defVal
		result := false
		if v1.(bool) == v2.(bool) {
			result = true
		}
		res.SetDefault(result)
		return res, nil
	}

	combined := unique(keyCombineOrdered(uts.Keys(), other.Keys()))

	res = NewUSTimeSerie(len(combined))
	v1 := uts.defVal
	v2 := other.defVal
	result := false
	if v1.(bool) == v2.(bool) {
		result = true
	}
	res.SetDefault(result)

	for _, t := range combined {

		v1, _ := uts.Get(t)
		v2, _ := other.Get(t)
		ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:Xor]:computing value for time %s [ %+v XOR %+v ]", t, v1, v2)
		result := false
		if v1.(bool) == v2.(bool) {
			result = true
		}
		res.Add(t, result)

	}
	return res, nil
}

// Not Logical negation from this TimeSerie. The resulting TimeSerie will have the same time points
// computing logical  negation : this is `output[t] = NOT x[t1]`
func (uts *USTimeSerie) Not() (*USTimeSerie, error) {

	var res *USTimeSerie

	if uts.defVal == nil {
		return nil, fmt.Errorf("this USTimeSerie has no default value")
	}

	res = NewUSTimeSerie(uts.Len())
	v1 := uts.defVal
	result := !v1.(bool)
	res.SetDefault(result)

	for _, t := range uts.t {
		val := uts.m[t]
		result := !val.(bool)
		res.Add(t, result)
		ilog.Debugf(">>>>USTS DEBUG [USTimeSerie:Not]:computing value for time %s [ ! %+v ]", t, val)
	}
	return res, nil
}

// Combine merges this TimeSerie with other , The resulting TimeSerie will have the
// same time points in both previous series (max will be N+M if any repeated time)
// no logical operations done, it only merges data except on time concidence where logical OR will be placed.
// `output[t1] = x[t1] (if not exist y[t1])`
// `output[t2] = y[t2] (if not exist x[t2])`
// `output[t3] = x[t3] OR y[t3] ( if both series has values in t3)`
func (uts *USTimeSerie) Combine(other *USTimeSerie) (*USTimeSerie, error) {
	var res *USTimeSerie

	if uts.Len() == 0 && other.Len() == 0 {
		res = NewUSTimeSerie(0)
		return res, nil
	}

	combined := unique(keyCombineOrdered(uts.Keys(), other.Keys()))

	res = NewUSTimeSerie(len(combined))

	for _, t := range combined {

		v1, ok1 := uts.GetExact(t)
		v2, ok2 := other.GetExact(t)
		ilog.Debugf(">>>>USTS DEBUG:computing value for time %s [ %+v COMBINED %+v ]", t, v1, v2)
		var result interface{}
		switch {
		case ok1 && ok2:
			result = v1.(bool) || v2.(bool)
		case ok1:
			result = v1
		case ok2:
			result = v2
		}
		res.Add(t, result)

	}
	return res, nil
}

// func (uts *USTimeSerie) Multiply(other *USTimeSerie) {

// }

// func (uts *USTimeSerie) Difference(other *USTimeSerie) {

// }

// func (uts *USTimeSerie) Mean(start time.Time, end time.Time, mask *USTimeSerie, interpolate InterPolateOps) {

// }

// func (uts *USTimeSerie) MonvingAverage(samplPeriod time.Duration, winSize time.Duration, start time.Time, end time.Time) {

// }

// func (uts *USTimeSerie) Resample(period time.Duration) {

// }
