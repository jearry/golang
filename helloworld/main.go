package main

import (
	"fmt"
	"reflect"
)

func main(){


	count := 1

	t1 := reflect.TypeOf(count)
	v1 := reflect.ValueOf(count)

	t2 := reflect.TypeOf(&count)
	v2 := reflect.ValueOf(&count)

	fmt.Println(t1, t2, v1, v2)

	v1.SetInt(2)

	//v2.SetInt( 3)
}
