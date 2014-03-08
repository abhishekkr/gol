package goltime

import (
  "strconv"
  "time"
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


func (timestamp *Timestamp) Time() time.Time{
  return time.Date(timestamp.Year, time.Month(timestamp.Month), timestamp.Day,
                   timestamp.Hour, timestamp.Min, timestamp.Sec, 0, time.UTC)
}
