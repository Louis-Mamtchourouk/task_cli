package tests

import "fmt"


func tests2() {
  var intArr [3]int32
  intArr[1] = 123
  fmt.Println(intArr[0])
  fmt.Println(intArr[1:3])
}
