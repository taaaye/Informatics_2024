package lab7

import "fmt"

type shoes struct {
    Name  string
    Price float64
    Brand string
    Size  string
}

func (s *shoes) getPrice() float64 {
    return s.Price
}

func (s *shoes) setPrice(newPrice float64) {
    s.Price = newPrice
}

func (s *shoes) applyDiscount(discount float64) {
    if discount < 0 || discount > 100 {
        fmt.Println("Скидка должна быть в диапазоне от 0 до 100%")
        return
    }
    s.Price = s.Price * (1 - discount/100)
}

func (s *shoes) getName() string {
    return s.Name
}

func (s *shoes) setName(newName string) {
    s.Name = newName
}

func (s *shoes) setBrand(newBrand string) {
    s.Brand = newBrand
}

func (s *shoes) getProductInfo() string {
    if s.Name == "" || s.Size == "" || s.Brand == "" {
        return emptyFieldsMessage
    }
    return fmt.Sprintf("Обувь: %s, Размер: %s, Бренд: %s, Цена: %.2f", s.Name, s.Size, s.Brand, s.Price)
}