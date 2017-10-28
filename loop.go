package main

import "fmt"

func loop1(){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
}

func loop2(){
	x := 5
	for {
		fmt.Println("Do stuff",x)
		x+=3
		if x>25 {
			break
		}
	}
}

func loop3(){
	s:=1;
	for i :=30;i>0;i-=3{
		fmt.Println(i,s)
		s++
	}
}

func main(){
	//loop1()
	//loop2()
	loop3()
}