package lab6
import (
    "fmt"
	"os"
)

type Airplane struct {
    name            string
    flight_number   int
    power           int
    purchased_seats int
}

func GetAirplaneName() string {
    var name string
    fmt.Print("Введите имя самолета: ")
    fmt.Fscan(os.Stdin, &name)
    return name
}

func (a Airplane) speed_calculation(air_res int) int {
    speed := a.power / air_res
    return speed
}

func (a Airplane) revenue_calculation(price_for_place int) int {
    money := price_for_place * a.purchased_seats
    return money
}

func (a Airplane) get_route() string {
    routes := map[int]string{
        1: "Tokyo-Moscow",
        2: "Moscow-Tokyo",
        3: "Moscow-Dubai",
        4: "Seoul-Moscow",
        5: "Moscow-Seoul",
        6: "Dubai-Moscow",
    }

    return routes[a.flight_number]
}

func RunLab6() {
    routes := map[int]string{
        1: "Tokyo-Moscow",
        2: "Moscow-Tokyo",
        3: "Moscow-Dubai",
        4: "Seoul-Moscow",
        5: "Moscow-Seoul",
        6: "Dubai-Moscow",
    }

    fmt.Println("Выберите номер маршрута из списка:")
    for num, route := range routes {
        fmt.Printf("%d: %s\n", num, route)
    }

    var route int
    fmt.Print("Введите номер маршрута: ")
    fmt.Scan(&route)

    var occupiedSeats int
    fmt.Print("Введите количество занятых мест: ")
    fmt.Scan(&occupiedSeats)

    // Ввод имени самолета
    AirplaneName := GetAirplaneName()

    Plane := Airplane{
        name:            AirplaneName,
        flight_number:   route,
        power:           480000,
        purchased_seats: occupiedSeats,
    }

    fmt.Printf("%v направляется по маршруту %v\n", Plane.name, Plane.get_route())
    fmt.Printf("Скорость %v равна %d м/с\n", Plane.name, Plane.speed_calculation(2000))
    fmt.Printf("Прибыль с продажи билетов составит: %d рублей \n", Plane.revenue_calculation(40000))
}