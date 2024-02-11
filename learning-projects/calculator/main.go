package main

import (
	"errors"
	"fmt"
	"reflect"
)

func userInterface() int{
  var menu int
  fmt.Println(`
     ______                  __
 /  _____)                |
|  |  __ _   ____   ___  _|  ____   ___
|  | |_ \| |    \ /  _)| | / _  ) /  _)
|  |__| | | | | |(  (_ | |( (/ / (  (__
 \______| |_|_|_| \____)_| \____) \____)
  `)
  fmt.Println("1. Addition")
  fmt.Println("2. Substraction")
  fmt.Println("3. Multiplication")
  fmt.Println("4. Division")
  fmt.Println("5. Exit")
  fmt.Print("Choose Menu's Number: ")
  fmt.Scanln(&menu)
  return menu
}

func Addition(num1 float64, num2 float64)float64{
 var total float64 = num1 + num2
 return total
}

func Substraction(num1 float64, num2 float64)float64{
  result := num1 - num2
  return result
}

func Multiplication(num1,num2 float64)float64{
  result := num1 * num2
  return result
}

func Division(num1,num2 float64)float64{
  result := num1 / num2
  return result
}

func checkDataType(input interface{})error{
  expectedType := reflect.TypeOf(float64(0))
  var check interface{} = reflect.TypeOf(input)
  if check != expectedType{
    return errors.New("Wrong Data Type! Must be Float!")
  }
  return nil
}

func Run()bool{
  var err error = nil
  for {
   var number1,number2 float64

   err = checkDataType(number1)
   if err != nil {
     fmt.Println("Error: ", err)
   }
   err = checkDataType(number2)
   if err != nil{
     fmt.Println("Error: ", err)
   }

   menuInput := userInterface()
   if menuInput == 5{
     fmt.Println("Exiting . . . ")
     break
   }

   fmt.Print("First number: ")
   _, err = fmt.Scanf("%f", &number1)
   if err != nil {
     fmt.Println(err)
     continue
   }
     fmt.Print("Second number: ")
   _,err = fmt.Scanf("%f", &number2)
   if err != nil {
     fmt.Println(err)
     continue
   }
   switch menuInput{
     case 1:
       result:= Addition(number1, number2)
       fmt.Println("Result: ", result)
     case 2:
       result := Substraction(number1, number2)
       fmt.Println("Result: ", result)
     case 3:
       result := Multiplication(number1, number2)
       fmt.Println("Result: ", result)
     case 4:
       result := Division(number1, number2)
       fmt.Println("Result: ", result)
     default:
       fmt.Println("Wrong Menu Choice!")
   }
}
  return true
}

func main(){
  Run()
}
