# USTS ( Unevently Spaced Time Series )

Unevenly spaced time series library for Go and other time/slots/window helpers to deal with Unevenly spaced time events
This library ia a partial Golang port for the python traces library , with a few time event handling additions.

* https://github.com/datascopeanalytics/traces

* https://en.wikipedia.org/wiki/Unevenly_spaced_time_series
* https://traces.readthedocs.io/en/latest/


# Example

```go

package main

import (
	"fmt"
    "time"
    "github.com/datadope-io/usts"
)


func main() {

	ts := usts.NewUSTimeSerie(0)

	ts.Add(time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC), "NOK")
	ts.Add(time.Date(2042, 2, 1, 7, 45, 56, 0, time.UTC), "OK")
	ts.Add(time.Date(2042, 2, 1, 8, 51, 42, 0, time.UTC), "NOK")
	ts.Add(time.Date(2042, 2, 1, 12, 3, 56, 0, time.UTC), "OK")
	ts.Add(time.Date(2042, 2, 1, 12, 7, 13, 0, time.UTC), "NOK")

	t0 := time.Date(2042, 2, 1, 6, 0, 0, 0, time.UTC)
	t1 := time.Date(2042, 2, 1, 13, 0, 0, 0, time.UTC)

	m, total, err := ts.Distribution(t0, t1, nil)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return
    }
    
	totalsec := int64(total / time.Second)

	for k, v := range m {
		percent := float64(v/time.Second) * 100.0 / float64(totalsec)
		fmt.Printf("VALUE %v present for %s :  %.2f %%\n", k, v, percent)
	}

	// Unordered output:
	// VALUE NOK present for 5h50m57s :  83.56 %
	// VALUE OK present for 1h9m3s :  16.44 %
}

```
