# v0.1.8 (2021/03/31)

* added CompactBoundaries function to retrieve min/max boundaries of repeated events on USTs

# v0.1.7 (2021/03/29)

* fixed usts index getter when the events are outside the using timeref period (start/end)

# v0.1.6 (2021/03/26)

* added new MarkPeriods function to merge periods into new UTS
* fixed distribution when the mask has more defined events than base UTS

# v0.1.5 (2020/10/14)

* added DumpBufferWithPrefix function

# v0.1.4 (2020/07/25)

* fixed timeevents when start or end time matches a timeslot End Event.
* Added timeslot and timewindow tests

# v0.1.3 (2020/07/24)

* Removed duplicated tests
* Published timeslot Scr[Start Cron Format]/Ecr[End Cron Format] parameters.
* added Github actions
* added USTimeSerie.Clone() method
* fix distribution func when timeseries  doesn't match elements with  start/end distribution interval

# v0.1.0 (2020/06/24)

* released first UStimeSeries Handler version 0.1.0 
* Added complete doc for USTimeSeries object and main Operators
* Added custom logger
* Improved log messages

# v0.0.0 (2020/06/19)

first alpha released version
