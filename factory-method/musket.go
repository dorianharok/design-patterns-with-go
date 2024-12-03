package factory_method

type musket struct {
	Gun
}

func NewMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
