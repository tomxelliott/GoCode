package db

type Item struct {
	// If a structure field name starts with a lowercase letter,
	// Only code within the same package will be able to access them.
	Price float64
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 9.99,
	}
}
