package abstract_factory

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}
