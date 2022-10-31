package main

import (
	"fmt"
	"strconv"
)

func main() {
	myFavFloat, _ := strconv.ParseInt("42", 10, 64)
	fmt.Println(myFavFloat)
}
