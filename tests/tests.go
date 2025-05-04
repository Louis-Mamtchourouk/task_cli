package tests

import (
	"fmt"
	//"reflect"
	"errors"
)

func tests() {
  fmt.Println("Feur")

  //var test uint = 23
  //caca := 42.6
  //fmt.Println(caca)
  //fmt.Println(reflect.TypeOf(caca))

  var testi int16 = 44
  var denom0 int16 = 0
  var denom int16 = 4

  var result, remainder, err = intDivision(testi,denom)
  if err != nil {
    fmt.Println(err.Error())
  }else{
    fmt.Printf("Le résultat de la division est %v avec pour reste %v\n", result, remainder)
  }

  result, remainder, err = intDivision(testi,denom0)
  if err != nil {
    fmt.Println(err.Error())
  }else{
    fmt.Printf("Le résultat de la division est %v avec pour reste %v\n", result, remainder)
  }

  switch  {
  case err != nil:
    break
  case remainder == 0:
    break
  }

  switch remainder {
    case 1:
      break
    case 2:
      break
    default:
      break
  }

}

func intDivision(numerator int16, denominator int16) (int16,int16,error){
  var err error
  if denominator == 0{
    err = errors.New("Cannot divide by 0")
    return 0, 0, err
  }
  result := numerator/denominator
  remainder := numerator%denominator
  return result, remainder, err
}
