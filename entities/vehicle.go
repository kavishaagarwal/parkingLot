package entities

type Vehicle interface {
	GetSize() int
	GetType() string
}

type Car struct{}

func (c *Car) GetSize() int    { return 1 }
func (c *Car) GetType() string { return "Car" }

type Truck struct{}

func (t *Truck) GetSize() int    { return 2 }
func (t *Truck) GetType() string { return "Truck" }

type Motorcycle struct{}

func (m *Motorcycle) GetSize() int    { return 0 }
func (m *Motorcycle) GetType() string { return "Motorcycle" }

type ElectricCar struct{}

func (ec *ElectricCar) GetSize() int    { return 1 }
func (ec *ElectricCar) GetType() string { return "Electric Car" }
