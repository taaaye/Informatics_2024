package lab6

import (
	"fmt"
)

type Airplane struct {
	name           string
	flightNumber   int
	power          int
	purchasedSeats int
}

func (a Airplane) speedCalculation(airRes int) int {
	speed := a.power / airRes
	return speed
}

func (a Airplane) revenueCalculation(priceForPlace int) int {
	money := priceForPlace * a.purchasedSeats
	return money
}

func (a Airplane) getRoutes() string {
	routes := []string{"Ivanovo-Moscow", "Moscow-Ivanovo", "Moscow-Seoul", "Seoul-Moscow", "Moscow-Tokyo", "Tokyo-Moscow"}
    if a.flightNumber >= 0 && a.flightNumber < len(routes) {
        return routes[a.flightNumber]
    }   
    return "Неизвестный маршрут"
}    
func Laba6run() {
	Plane := Airplane{"Boing34", 5, 490000, 80}
	fmt.Printf("%v направляется по маршруту %v\n", Plane.name, Plane.getRoutes())
	fmt.Printf("Скорость %v равна %d м/с\n", Plane.name, Plane.speedCalculation(2000))
	fmt.Printf("Прибыль с продажи билетов составит: %d рублей\n", Plane.revenueCalculation(40000))
}
