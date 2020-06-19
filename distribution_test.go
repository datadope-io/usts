package usts

import (
	"fmt"
	"testing"
	"time"
)

func TestDistribution(t *testing.T) {
	var err error
	var t0, t1, tm0, tm1 time.Time
	var total time.Duration
	var tot1, tot0, totalsec int64
	var m map[interface{}]time.Duration

	ts := NewUSTimeSerie(100)

	ts.Add(time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC), 0)
	ts.Add(time.Date(2042, 2, 1, 7, 45, 56, 0, time.UTC), 1)
	ts.Add(time.Date(2042, 2, 1, 8, 51, 42, 0, time.UTC), 0)
	ts.Add(time.Date(2042, 2, 1, 12, 3, 56, 0, time.UTC), 1)
	ts.Add(time.Date(2042, 2, 1, 12, 7, 13, 0, time.UTC), 0)

	//A (1)
	t.Log("* Distribition basic(no mask)full range-------------------")

	t0 = time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 13, 0, 0, 0, time.UTC)

	//VALUE [0] took [5h50m57s] (=21057 seconds) (0.835595)
	//VALUE [1] took [1h9m3s] (=4143 seconds) (0.164405)
	//Total Seconds: 25200

	m, total, err = ts.Distribution(t0, t1, nil)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 25200 {
		t.Errorf("Distribution Error Got TotalTime %d want 25200 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = int64(m[1] / time.Second)
	if tot0 != 21057 || tot1 != 4143 {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 21057/4143", tot0, tot1)
	}

	//B (2)
	t.Log("* Distribition basic(no mask) from 6:16/6:30 only 0------------------")

	t0 = time.Date(2042, 2, 1, 6, 15, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 6, 30, 0, 0, time.UTC)

	//VALUE [0] took [15m] (=900 seconds) (1)
	//VALUE [1] ------------
	//Total Seconds: 900

	m, total, err = ts.Distribution(t0, t1, nil)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 900 {
		t.Errorf("Distribution Error Got TotalTime %d want 900 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = 0
	if tot0 != 900 || tot1 != 0 {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 900/0", tot0, tot1)
	}

	//C (3)
	t.Log("* Distribition basic(no mask) from 6:45/8:00 both 0/1------------------")

	t0 = time.Date(2042, 2, 1, 6, 45, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 8, 00, 0, 0, time.UTC)

	//VALUE [0] took [1h0m56s] (=3656 seconds)
	//VALUE [1] took [14m4s] (=844 seconds)
	//Total Seconds: (1h15m) (=4500)

	m, total, err = ts.Distribution(t0, t1, nil)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 4500 {
		t.Errorf("Distribution Error Got TotalTime %d want 4500 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = int64(m[1] / time.Second)
	if tot0 != 3656 || tot1 != 844 {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 3656/844", tot0, tot1)
	}

	//D (4)
	t.Log("* Distribition basic(no mask) from 8:45/10:00 both 0/1------------------")

	t0 = time.Date(2042, 2, 1, 8, 45, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 10, 00, 0, 0, time.UTC)

	//VALUE [1] took 8:45:00-> 8:51:42  [6m42s](=402 seconds) (https://www.timeanddate.com/date/timeduration.html)
	//VALUE [0] took 8:51:42-> 10:00:00 [1h8m18s](=4098 seconds)
	//Total Seconds: (1h15m) (=4500)

	m, total, err = ts.Distribution(t0, t1, nil)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 4500 {
		t.Errorf("Distribution Error Got TotalTime %d want 4500 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = int64(m[1] / time.Second)
	if tot0 != 4098 || tot1 != 402 {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 4098/402", tot0, tot1)
	}

	// MASK1 (5)
	t.Log("* Distribition basic( mask from 8:45/10:00) full range---------------")

	tm0 = time.Date(2042, 2, 1, 8, 45, 0, 0, time.UTC)
	tm1 = time.Date(2042, 2, 1, 10, 00, 0, 0, time.UTC)

	mask := NewUSTimeSerie(0)
	mask.SetDefault(false)
	mask.SetIntervalValue(tm0, tm1, true)

	t0 = time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 13, 0, 0, 0, time.UTC)

	m, total, err = ts.Distribution(t0, t1, mask)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 4500 {
		t.Errorf("Distribution Error Got TotalTime %d want 4500 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = int64(m[1] / time.Second)
	if tot0 != 4098 || tot1 != 402 {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 4098/402", tot0, tot1)
	}

	//MASK2 (6)
	t.Log("* Distribition basic( with mask from 6:45/8:00 and 8:45/10:00) full range---------------")

	tm0 = time.Date(2042, 2, 1, 6, 45, 0, 0, time.UTC)
	tm1 = time.Date(2042, 2, 1, 8, 00, 0, 0, time.UTC)
	mask.SetIntervalValue(tm0, tm1, true) //now 2 windows

	m, total, err = ts.Distribution(t0, t1, mask)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	totalsec = int64(total / time.Second)
	if totalsec != 9000 {
		t.Errorf("Distribution Error Got TotalTime %d want 4500 ", totalsec)
	}
	tot0 = int64(m[0] / time.Second)
	tot1 = int64(m[1] / time.Second)
	if tot0 != (4098+3656) || tot1 != (402+844) {
		t.Errorf("Distribution Error Got 0/1 Seconds %d/%d want 4098+3656/402+844 (7754/1246)", tot0, tot1)
	}
}

func ExampleDistribution() {
	var err error
	var t0, t1 time.Time
	var total time.Duration
	var totalsec int64
	var m map[interface{}]time.Duration

	ts := NewUSTimeSerie(100)

	ts.Add(time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC), "NOK")
	ts.Add(time.Date(2042, 2, 1, 7, 45, 56, 0, time.UTC), "OK")
	ts.Add(time.Date(2042, 2, 1, 8, 51, 42, 0, time.UTC), "NOK")
	ts.Add(time.Date(2042, 2, 1, 12, 3, 56, 0, time.UTC), "OK")
	ts.Add(time.Date(2042, 2, 1, 12, 7, 13, 0, time.UTC), "NOK")

	//A (1)

	t0 = time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC)
	t1 = time.Date(2042, 2, 1, 13, 0, 0, 0, time.UTC)

	//VALUE [0] took [5h50m57s] (=21057 seconds) (0.835595)
	//VALUE [1] took [1h9m3s] (=4143 seconds) (0.164405)
	//Total Seconds: 25200

	m, total, err = ts.Distribution(t0, t1, nil)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
	}
	totalsec = int64(total / time.Second)

	for k, v := range m {
		percent := float64(v/time.Second) * 100.0 / float64(totalsec)
		fmt.Printf("VALUE %v present for %s :  %.2f %%\n", k, v, percent)
	}

	// Unordered output:
	// VALUE NOK present for 5h50m57s :  83.56 %
	// VALUE OK present for 1h9m3s :  16.44 %
}
