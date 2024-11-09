package lab6
import (
    "fmt"
	"os"
)

var Routes = map[int]string{
    1: "Tokyo-Moscow",
    2: "Moscow-Tokyo",
    3: "Moscow-Dubai",
    4: "Seoul-Moscow",
    5: "Moscow-Seoul",
    6: "Dubai-Moscow",
}

type Airplane struct {
    Name            string
    FlightNumber   int
    Power           int
    PurchasedSeats int
}

func GetAirplaneName() string {
    var Name string
    fmt.Print("Введите имя самолета: ")
    fmt.Fscan(os.Stdin, &Name)
    return Name
}

func (a Airplane) SpeedCalculation(AirRes int) int {
    Speed := a.Power / AirRes
    return Speed
}

func (a Airplane) RevenueCalculation(PriceForPlace int) int {
    Money := PriceForPlace * a.PurchasedSeats
    return Money
}

func (a Airplane) GetRoute() string {
    return Routes[a.FlightNumber]
}

func RunLab6() {
    fmt.Println("Выберите номер маршрута из списка:")
    for num, Route := range Routes {
        fmt.Printf("%d: %s\n", num, Route)
    }

    var Route int
    fmt.Print("Введите номер маршрута: ")
    fmt.Scan(&Route)

    var OccupiedSeats int
    fmt.Print("Введите количество занятых мест: ")
    fmt.Scan(&OccupiedSeats)

    AirplaneName := GetAirplaneName()

    Plane := Airplane{
        Name:           AirplaneName,
        FlightNumber:   Route,
        Power:          490000,
        PurchasedSeats: OccupiedSeats,
    }

    fmt.Printf("%v направляется по маршруту %v\n", Plane.Name, Plane.GetRoute())
    fmt.Printf("Скорость %v равна %d м/с\n", Plane.Name, Plane.SpeedCalculation(2000))
    fmt.Printf("Прибыль с продажи билетов составит: %d рублей \n", Plane.RevenueCalculation(40000))
}