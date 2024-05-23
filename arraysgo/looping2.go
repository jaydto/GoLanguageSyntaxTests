package arraysgo

import "fmt"

func Looping2(){

	names:=[]string{"name", "name2", "name4"}

	for i, v:=range names{
		fmt.Printf("\n position: %v value: %v", i, v)
	}

}