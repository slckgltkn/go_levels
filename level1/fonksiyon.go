package main

import "fmt"

func main(){
	var sayi1 int = 999999
	var sayi2 = 999999999
	var sayi3 = 10000000
	yazdir(sayi1,sayi2,sayi3)
	fmt.Println(islem (5,4,4))
	fmt.Println(ikiliislem(11,21,34))
	toplam, carpim := degiskenlidondur(1,2)
	fmt.Println(toplam, carpim)
}

func yazdir(a int, b int, c int){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func islem(a int, b int ,c int) int{
	return a+b+c 
}

func ikiliislem(a int, b int, c int) (int, int){
	return a*b, a+c
}

func degiskenlidondur(a int, b int) (toplam int, carpim int){
	toplam = a + b
	carpim = a * b
	return
}

