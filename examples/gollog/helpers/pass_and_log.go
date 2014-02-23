package gollog_example


type ReceiveStringReturnNil func(msg string)


func PassAndLog(foo ReceiveStringReturnNil){
  foo("passed and logged")
}
