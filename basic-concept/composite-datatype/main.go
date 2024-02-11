package main

import (
	"fmt"
)

func main(){
  // Declaring an array
  var odds [5]int
  odds[0] = 1
  odds[1] = 3
  odds[2] = 5
  odds[3] = 7
  odds[4] = 9
  even := [5]int{2,4,6,8,10}
  var warna = make([]string, 5)
  warna[0] = "Merah"
  warna[1] = "Kuning"
  warna[2] = "Hijau"
  warna[3] = "Biru"
  warna[4] = "Ungu"
  
  for i:= 0 ; i < len(odds); i++ {
    fmt.Println("odds array: ", odds[i])
  }

  for _, j := range even{
    fmt.Println("even array: ", j)
  }

  for i:=0; i<len(warna); i++{
    fmt.Println("Warna: ", warna[i])
  }
  
  // Multidimensional array
  var goodies = [2][3]string{{"Apple", "Biscuit", "Cake"},{"Diet Coke", "Energen", "Fanta"}}
  for i:= 0; i<len(goodies); i++{
    for j:= 0; j<len(goodies[i]); j++{
      fmt.Println("goodies: ", goodies[i][j])
    }
  }
  var numbers = [2][5]int{odds, even}
  for _,row := range numbers {
    for _, value := range row {
      fmt.Println("Numbers: ", value)
    }
  }

  // Declaring slice
  var odds_slice []int = odds[1:3]
  for _, odds_slice_value := range odds_slice{
    fmt.Println("odds slice: ", odds_slice_value)
  }
  even_slice := even[1:3]
  for i:=0; i < len(even_slice); i++{
    fmt.Println("even slice: ", even_slice[i])
  }
}
