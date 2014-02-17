package gol


type RecieveAndRun func()


func Gol(_compute RecieveAndRun) {
  for {
    _compute()
  }
}
