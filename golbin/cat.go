package golbin


import (
  "io/ioutil"
)


func Cat(filepath string) (string, error){
  bytes, err := ioutil.ReadFile(filepath)
  if err != nil { return "", err }
  return string(bytes), nil
}
