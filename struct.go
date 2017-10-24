package main

import "fmt"


const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
    gas_pedal uint16 //0-64k
    brake_pedal uint16
    steering_wheel int16 // -32k - +32k
    top_speed_kmh float64
}

func (c car) kmh() float64 {
    return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
    return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}
func (c *car) new_top_speed(newspped float64){
    c.top_speed_kmh = newspped
}

func newer_top_speed(c car,speed float64) car{
    c.top_speed_kmh = speed
    return c
}



func main(){
    a_car := car{gas_pedal:22341,
                brake_pedal:0,
                steering_wheel:12345,
                top_speed_kmh:222.3}
    //a_car := car{22341,0,12345,222.3}

    fmt.Println(a_car.gas_pedal)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
    //a_car.new_top_speed(640)
    a_car = newer_top_speed(a_car,343)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())

}
