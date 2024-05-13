package entity

type Computer struct {
	Id           int
	PricePerHour int
}

func NewComputer(id int, price int) Computer {
	return Computer{
		Id:           id,
		PricePerHour: price,
	}
}
