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
    var tv Product = &Electronics{price: 24999.99, name: "Телевизор", brand: "Samsung", color: "Черный"}
    var hoodys Product = &Hoody{price: 5000, name: "Толстовка Nikifilini", color: "Черный", size: "XXL"}
    var trainers Product = &Shoes{price: 39000, name: "Air Jordan 3 'Black Cat'", brand: "Nike", size: "42"}

    tv.setPrice(29999.99)
    tv.setName("Монитор")
    
    hoodys.(*Hoody).setColor("Белый")
    trainers.(*Shoes).setName("RS-X Efekt S&P")
    trainers.(*Shoes).setBrand("Puma")

    listproducts := []Product{tv, hoodys, trainers}

    fmt.Printf("Сумма товаров, без учёта скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))

    tv.applyDiscount(5)
    hoodys.applyDiscount(25)
    trainers.applyDiscount(15)

    fmt.Printf("Сумма товаров, с учётом скидки, равна: %.2f рублей \n", CalculationSumProduct(listproducts))
}