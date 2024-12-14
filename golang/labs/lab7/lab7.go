package lab7

import "fmt"

func CalculationSumProduct(listproducts []Product) float64 {
    var sum float64 = 0
    for _, product := range listproducts {
        sum += product.getPrice()
    }
    return sum
}

func Laba7run() {
    tv := &electronics{Price: 24999.99, Name: "Телевизор", Brand: "Samsung", Color: "Черный"}
    hoodys := &hoody{Price: 5000, Name: "Толстовка Nikifilini", Color: "Черный", Size: "XXL"}
    trainers := &shoes{Price: 39000, Name: "Air Jordan 3 'Black Cat'", Brand: "Nike", Size: "42"}

    tv.setPrice(29999.99)
    tv.setName("Монитор")
    hoodys.setColor("Белый")
    trainers.setName("RS-X Efekt S&P")
    trainers.setBrand("Puma")

    listproducts := []Product{tv, hoodys, trainers}

    fmt.Printf("Сумма товаров, без учёта скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))

    tv.applyDiscount(5)
    hoodys.applyDiscount(25)
    trainers.applyDiscount(15)

    fmt.Printf("Сумма товаров, с учётом скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))
}
