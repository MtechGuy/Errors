//Demonstrate errors
package main

import (
	"fmt"
	
)

func area(lenght float64 , width float64) float64 {
	result := lenght * width
	return result

}

func main(){
	length := 5.5
	width := -10
	//call the area function
	result := area(length, float64(width))
	fmt.Println(result)
}
