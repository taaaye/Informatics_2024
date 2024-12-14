package lab7

import "fmt"

type electronics struct {
    Name  string
    Price float64
    Brand string
    Color string
}

const emptyFieldsMessage = "Присутствуют пустые поля"

func (e *electronics) getPrice() float64 {
    return e.Price
}

func (e *electronics) setPrice(newPrice float64) {
    e.Price = newPrice
}

func (e *electronics) applyDiscount(discount float64) {
    if discount < 0 || discount > 100 {
        fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
        return
    }
    e.Price = e.Price * (1 - discount/100)
}

func (e *electronics) getName() string {
    return e.Name
}

func (e *electronics) setName(newName string) {
    e.Name = newName
}

func (e *electronics) getProductInfo() string {
    if e.Name == "" || e.Brand == "" || e.Color == "" {
        return emptyFieldsMessage
    }
    return fmt.Sprintf("Электроника: %s, Бренд: %s, Цвет: %s, Цена: %.2f", e.Name, e.Brand, e.Color, e.Price)
}
