package labs
//9 вариант
import (
  "fmt"
  "math"
)

func CalculateY(x, a, b float64) float64 {
  numerator := math.Log10(x*x - 1)
  denominator := math.Log(a*x*x-b) / math.Log(5)
  y := numerator / denominator
  return y
}

func TaskA(xn, xk, xdelta, a, b float64) []float64 {
  var y_a = []float64{}
  for i := xn; i < xk; i = i + xdelta {
    y_a = append(y_a, CalculateY(i, a, b))
  }
  return y_a
}

func TaskB(xs []float64, a, b float64) []float64 {
  var y_b = []float64{}
  for _, x_b := range xs {
    y_b = append(y_b, CalculateY(x_b, a, b))
  }
  return y_b  
}

func Laba4run() {
  fmt.Print("Задача A\n")
  fmt.Print(TaskA(1.2, 2.2, 0.2, 1.1, 0.09), "\n")
  fmt.Print("Задача B\n")
  var x = []float64{1.21, 1.76, 2.53, 3.48, 4.52}
  fmt.Print(TaskB(x, 1.1, 0.09), "\n")
}