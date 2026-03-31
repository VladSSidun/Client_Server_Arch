package order

import "fmt"

type ProductType string

const (
	Electronics ProductType = "Електроніка"
	Clothing    ProductType = "Одяг"
	Food        ProductType = "Їжа"
)

type Product struct {
	Name        string
	Type        ProductType
	Price       float64
	DiscountPct float64
}

func NewProduct(name string, pType ProductType, price float64) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("ціна товару '%s' повинна бути більше 0", name)
	}
	return &Product{Name: name, Type: pType, Price: price, DiscountPct: 0}, nil
}

func (p *Product) SetDiscount(pct float64) error {
	if pct < 0 || pct > 100 {
		return fmt.Errorf("знижка повинна бути від 0 до 100%%")
	}
	p.DiscountPct = pct
	fmt.Printf("[Знижка] На товар '%s' встановлено знижку %.0f%%\n", p.Name, pct)
	return nil
}

func (p *Product) FinalPrice() float64 {
	return p.Price * (1 - p.DiscountPct/100)
}

func (p *Product) Print() {
	fmt.Printf("  %-20s | %-12s | Ціна: %.2f грн | Знижка: %.0f%% | Підсумок: %.2f грн\n",
		p.Name, p.Type, p.Price, p.DiscountPct, p.FinalPrice())
}

type Order struct {
	ID       int
	Products []*Product
}

func NewOrder(id int) *Order {
	return &Order{ID: id}
}

func (o *Order) AddProduct(p *Product) error {
	if p == nil {
		return fmt.Errorf("неможливо додати порожній товар")
	}
	o.Products = append(o.Products, p)
	fmt.Printf("[Замовлення #%d] Додано товар: %s\n", o.ID, p.Name)
	return nil
}

func (o *Order) TotalPrice() float64 {
	total := 0.0
	for _, p := range o.Products {
		total += p.FinalPrice()
	}
	return total
}

func (o *Order) TotalByType(pType ProductType) float64 {
	total := 0.0
	for _, p := range o.Products {
		if p.Type == pType {
			total += p.FinalPrice()
		}
	}
	return total
}

func (o *Order) Print() {
	fmt.Printf("\n===== ЗАМОВЛЕННЯ #%d =====\n", o.ID)
	if len(o.Products) == 0 {
		fmt.Println("Замовлення порожнє")
		return
	}
	for _, p := range o.Products {
		p.Print()
	}
	fmt.Printf("  Загальна вартість: %.2f грн\n", o.TotalPrice())
	fmt.Println("==========================")
}