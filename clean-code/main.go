//Old Code (not Clean Code)
/*
package main

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}
*/

// Refactor (Clean Code)
package main

type Vehicle struct {
	TotalWheels  int
	SpeedPerHour int
}

type Car struct {
	Vehicle
}

func (c *Car) Move() {
	c.IncreseSpeed(10)
}

func (c *Car) IncreseSpeed(newSpeed int) {
	c.SpeedPerHour += newSpeed //current speed plus the new speed
}

func main() {
	fastCar := Car{}
	fastCar.Move()
	fastCar.Move()
	fastCar.Move()

	slowCar := Car{}
	slowCar.Move()
}
