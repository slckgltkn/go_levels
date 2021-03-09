package main

import "fmt"

func main(){
	//DİZİLER ARRAYS
	var a = []int{1,2,3,4,}
	fmt.Println(a)
	var b = a[0:3]
	fmt.Println(b)
	c := make([]int, 15)
	c[1] = 35
	c[14] = 155
	fmt.Println(c)
	}
