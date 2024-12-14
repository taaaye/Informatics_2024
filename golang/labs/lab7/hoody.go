package lab7

import "fmt"

type hoody struct {
    Size  string
    Price float64
    Name  string
    Color string
}

func (h *hoody) getPrice() float64 {
    return h.Price
}

func (h *hoody) setPrice(newPrice float64) {
    h.Price = newPrice
}

func (h *hoody) applyDiscount(discount float64) {
    if discount < 0 || discount > 100 {
        fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
        return
    }
    h.Price = h.Price * (1 - discount/100)
}

func (h *hoody) getName() string {
    return h.Name
}

func (h *hoody) setName(newName string) {
    h.Name = newName
}

func (h *hoody) setColor(newColor string) {
    h.Color = newColor
}

func (h *hoody) getProductInfo() string {
    if h.Name == "" || h.Size == "" || h.Color == "" {
        return "Присутствуют пустые поля"
    }
    return fmt.Sprintf("Худи: %s, Размер: %s, Цвет: %s, Цена: %.2f", h.Name, h.Size, h.Color, h.Price)
}
