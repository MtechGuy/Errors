//Demonstrate errors
package main

import (
	"errors"
	"fmt"
	"log"
)

func area(lenght float64 , width float64) (float64, error) {
	//check if length is greater that 0
	if lenght < 0 {
		err:=errors.New("you can not have a negative length")
		return -1, err
	}
	if width < 0 {
		err:=errors.New("you can not have a negative width")
		return -1, err
	}
	result := lenght * width
	return result, nil

}

func main(){
	length := 5.5
	width := 10
	//call the area function
	result, err:= area(length, float64(width))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
