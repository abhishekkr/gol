package goltime

import (
  "strconv"
  "time"
  "net/http"
)


type Timestamp struct{
  Year, Month, Day, Hour, Min, Sec int
}


func CreateTimestamp(time_point []string) Timestamp{
  year, _   := strconv.Atoi(time_point[0])
  month, _  := strconv.Atoi(time_point[1])
  day, _    := strconv.Atoi(time_point[2])
  hour, _   := strconv.Atoi(time_point[3])
  min, _    := strconv.Atoi(time_point[4])
  sec, _    := strconv.Atoi(time_point[5])

  return Timestamp{
    Year: year,
    Month: month,
    Day: day,
    Hour: hour,
    Min: min,
    Sec: sec,
  }
}


func TimestampFromHTTPRequest(req *http.Request) Timestamp{
  return CreateTimestamp([]string {
    req.Form["year"][0], req.Form["month"][0], req.Form["day"][0],
    req.Form["hour"][0], req.Form["min"][0], req.Form["sec"][0],
  })
}


func (timestamp *Timestamp) Time() time.Time{
  return time.Date(timestamp.Year, time.Month(timestamp.Month), timestamp.Day,
                   timestamp.Hour, timestamp.Min, timestamp.Sec, 0, time.UTC)
}
