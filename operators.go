package usts

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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
		log.Debugf("computing value for time %s [ %+v AND %+v ]", t, v1, v2)
		result := (v1.(bool) && v2.(bool))
		res.Add(t, result)

	}
	return res, nil
}

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
		log.Debugf("computing value for time %s [ %+v AND %+v ]", t, v1, v2)
		result := (v1.(bool) || v2.(bool))
		res.Add(t, result)

	}
	return res, nil
}

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
		log.Debugf("computing value for time %s [ %+v XOR %+v ]", t, v1, v2)
		result := false
		if v1.(bool) == v2.(bool) {
			result = true
		}
		res.Add(t, result)

	}
	return res, nil
}

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
		log.Debugf("computing value for time %s [ ! %+v ]", t, val)
	}
	return res, nil
}

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
		log.Debugf("computing value for time %s [ %+v COMBINED %+v ]", t, v1, v2)
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

func (uts *USTimeSerie) Multiply(other *USTimeSerie) {

}

func (uts *USTimeSerie) Difference(other *USTimeSerie) {

}

func (uts *USTimeSerie) Mean(start time.Time, end time.Time, mask *USTimeSerie, interpolate InterPolateOps) {

}

func (uts *USTimeSerie) MonvingAverage(samplPeriod time.Duration, winSize time.Duration, start time.Time, end time.Time) {

}

func (uts *USTimeSerie) Resample(period time.Duration) {

}
