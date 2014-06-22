package goltime

import (
	"net/http"
	"strconv"
	"time"
)

// the currency for goltime
type Timestamp struct {
	Year, Month, Day, Hour, Min, Sec int
}

// creating Timestamp struct from string array of counts [Y M D h m s]
func CreateTimestamp(time_point []string) Timestamp {
	year, _ := strconv.Atoi(time_point[0])
	month, _ := strconv.Atoi(time_point[1])
	day, _ := strconv.Atoi(time_point[2])
	hour, _ := strconv.Atoi(time_point[3])
	min, _ := strconv.Atoi(time_point[4])
	sec, _ := strconv.Atoi(time_point[5])

	return Timestamp{
		Year:  year,
		Month: month,
		Day:   day,
		Hour:  hour,
		Min:   min,
		Sec:   sec,
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
		timestamp.Hour, timestamp.Min, timestamp.Sec, 0, time.UTC)
}

// create Timestamp struct from time.Now
func TimestampNow() Timestamp {
	var year, day, hour, min, sec int
	var month time.Month
	year, month, day = time.Now().Date()
	hour, min, sec = time.Now().Clock()
	return Timestamp{
		Year:  year,
		Month: int(month),
		Day:   day,
		Hour:  hour,
		Min:   min,
		Sec:   sec,
	}
}
