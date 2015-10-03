package goltime

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

// the currency for goltime
type Timestamp struct {
	Year, Month, Day, Hour, Min, Sec, MilliSec, MicroSec, NanoSec, PicoSec int
}

// getting value for second fragments from second value
func secondFragments(second string) (sec, milliSec, microSec, nanoSec, picoSec int) {
	fragments := strings.Split(second, ".")

	if len(fragments) > 0 {
		sec, _ = strconv.Atoi(fragments[0])
	}
	if len(fragments) > 1 {
		milliSec, _ = strconv.Atoi(fragments[1])
	}
	if len(fragments) > 2 {
		microSec, _ = strconv.Atoi(fragments[2])
	}
	if len(fragments) > 3 {
		nanoSec, _ = strconv.Atoi(fragments[3])
	}
	if len(fragments) > 4 {
		picoSec, _ = strconv.Atoi(fragments[4])
	}
	return
}

// creating Timestamp struct from string array of counts [Y M D h m s]
func CreateTimestamp(time_point []string) Timestamp {
	year, _ := strconv.Atoi(time_point[0])
	month, _ := strconv.Atoi(time_point[1])
	day, _ := strconv.Atoi(time_point[2])
	hour, _ := strconv.Atoi(time_point[3])
	min, _ := strconv.Atoi(time_point[4])
	sec, milliSec, microSec, nanoSec, picoSec := secondFragments(time_point[5])

	return Timestamp{
		Year:     year,
		Month:    month,
		Day:      day,
		Hour:     hour,
		Min:      min,
		Sec:      sec,
		MilliSec: milliSec,
		MicroSec: microSec,
		NanoSec:  nanoSec,
		PicoSec:  picoSec,
	}
}

/*
creating Timestamp from Request handler of HTTP having formvalues for it
under "year", "month", "day", "hour", "min", "sec"
*/
func TimestampFromHTTPRequest(req *http.Request) Timestamp {
	return CreateTimestamp([]string{
		req.FormValue("year"), req.FormValue("month"), req.FormValue("day"),
		req.FormValue("hour"), req.FormValue("min"), req.FormValue("sec"),
	})
}

// create time.Time from a Timestamp struct
func (timestamp *Timestamp) Time() time.Time {
	return time.Date(timestamp.Year, time.Month(timestamp.Month), timestamp.Day,
		timestamp.Hour, timestamp.Min, timestamp.Sec, timestamp.MilliSec, time.UTC)
}

// create Timestamp struct from time.Now
func TimestampNow() Timestamp {
	var year, day, hour, min, sec int
	var month time.Month
	year, month, day = time.Now().Date()
	hour, min, sec = time.Now().Clock()
	return Timestamp{
		Year:     year,
		Month:    int(month),
		Day:      day,
		Hour:     hour,
		Min:      min,
		Sec:      sec,
		MilliSec: 0,
		MicroSec: 0,
		NanoSec:  0,
		PicoSec:  0,
	}
}
