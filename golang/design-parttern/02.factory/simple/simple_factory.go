package factory

type Car interface {
	Drive() string
}

type BMW struct{}

func (c BMW) Drive() string {
	return "BMW"
}

type Cadillac struct{}

func (c Cadillac) Drive() string {
	return "Cadillac"
}

type Geely struct {
}

func (c Geely) Drive() string {
	return "Geely"
}

type CarFactory struct{}

func (c CarFactory) New(name string) Car {
	switch name {
	case "bmw":
		return &BMW{}
	case "cadillac":
		return &Cadillac{}
	case "geely":
		return &Geely{}
	default:
		return nil
	}
}
