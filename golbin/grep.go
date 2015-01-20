package golbin


import (
  "regexp"
  "strings"
)


func Grep(re, lines string) (string, error) {
  match := ""
  regex, regex_err := regexp.Compile(re)
  if regex_err != nil { return "", regex_err }

  for _, line := range strings.Split(lines, "\n") {
    if regex.MatchString(line) {
      match = match + "\n" + line
    }
  }
  return match, nil
}
