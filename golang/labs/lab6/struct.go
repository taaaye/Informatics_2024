package labs
//9 вариант
import (
	"fmt"
	"os"
)

type AirPlane struct {
	AirSpeed        float64
	TypePlane       string
	CountFreePlaces int16
   }
   
   func GetSpeed()  float64{
		var speed float64
		fmt.Print("Введите скорость самолета: ")
		fmt.Fscan(os.Stdin, &speed)
		return speed;
   }
   

   func GetTypePlane() string {
		var typePlane string
		fmt.Print("Введите модель самолета: ")
		fmt.Fscan(os.Stdin, &typePlane)
		return typePlane
   }
   
   func GetFreePlaces() int16 {
	var allCountPlace int16
    var countReservedPlaces int16
	fmt.Print("Введите общее кол-во мест ")
	fmt.Fscan(os.Stdin, &allCountPlace)

	fmt.Print("Введите общее кол-во занятых мест ")
	fmt.Fscan(os.Stdin, &countReservedPlaces)

	return allCountPlace - countReservedPlaces
   }
   
   func RunStruct() {
	var a = AirPlane{}

	a.AirSpeed=GetSpeed()
   
	a.TypePlane = GetTypePlane()
   
	a.CountFreePlaces=GetFreePlaces()

	fmt.Print("Скорость самолета " , a.AirSpeed, " Модель ",a.TypePlane, " Кол-во свободных мест ", a.CountFreePlaces)
   }
