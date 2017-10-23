package main

import ("fmt"
        "math/rand")

func foo(){
    fmt.Println("WHAAAT?")
}

func foo2(){
    fmt.Println("A number from 1-100: ",rand.Intn(100))
}

//const x int = 5

func multiple(a,b string) (string,string) {
    return a,b
}
func add(x,y float64) float64{
    return x+y
}
func foo3(){
    var a int = 62
    var b float64 = float64(a)
    x := a //x int
    fmt.Println(b,x)
}

func foo4(){
    x := 15
    a := &x //mem. address
    fmt.Println(a)
    fmt.Println(*a)
    *a = 5
    fmt.Println(x)
    *a = *a**a
    fmt.Println(x)
    fmt.Println(*a)
}


func main(){
    foo()
    foo2()

    num1,num2 := 5.5,43.3
    fmt.Println(add(num1,num2))
    w1,w2 := "Hey","There"
    fmt.Println(multiple(w1,w2))
    foo3()
    foo4()

}
