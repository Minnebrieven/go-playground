package main

import (
	"fmt"
	"sync"
)

type Product struct {
	Name  string
	Brand string
	Price int
	Unit  int
	m     sync.Mutex
}

var dummyData = map[int]Product{
	0: Product{
		Name:  "G102 Prodigy",
		Brand: "Logitech",
		Price: 210000,
		Unit:  100,
	},
	1: Product{
		Name:  "Mechanical Keyboard K102",
		Brand: "Logitech",
		Price: 300000,
		Unit:  50,
	},
}

func (p *Product) SetProduct(name, brand string, price, unit int) {
	p.m.Lock()
	defer p.m.Unlock()

	p.Name = name
	p.Brand = brand
	p.Price = price
	p.Unit = unit
}

func (p *Product) GetName() string {
	p.m.Lock()
	defer p.m.Unlock()

	nameProduct := p.Name

	return nameProduct
}

func (p *Product) GetBrand() string {
	p.m.Lock()
	defer p.m.Unlock()

	brandProduct := p.Brand

	return brandProduct
}

func (p *Product) GetPrice() int {
	p.m.Lock()
	defer p.m.Unlock()

	priceProduct := p.Price

	return priceProduct
}

func (p *Product) GetUnit() int {
	p.m.Lock()
	defer p.m.Unlock()

	unitProduct := p.Unit

	return unitProduct
}

func (p *Product) IncreaseUnit(unit int) {
	p.m.Lock()
	defer p.m.Unlock()

	p.Unit += unit
	fmt.Println("Unit Increased to ", p.Unit, " by ", unit)
}

func (p *Product) DecreaseUnit(unit int) {
	p.m.Lock()
	defer p.m.Unlock()

	p.Unit -= unit
	fmt.Println("Unit Decreased to ", p.Unit, " by ", unit)
}

func newProduct(inputNama, inputBrand string, inputPrice, inputUnit int) Product {
	var product = Product{
		Name:  inputNama,
		Brand: inputBrand,
		Price: inputPrice,
		Unit:  inputUnit,
	}

	return product
}

func showAllProduct() {
	menuHeader("Show All Product")
	for id, val := range dummyData {
		fmt.Println("\nId Product =>   ", id)
		fmt.Println("Product Name =>   ", val.GetName())
		fmt.Println("Product Brand =>  ", val.GetBrand())
		fmt.Println("Product Price =>  ", val.GetPrice())
		fmt.Println("Product Unit =>   ", val.GetUnit())
	}
	fmt.Println("==================================\n")
}

func addNewProduct() {
	var name, brand string
	var price, unit int

	menuHeader("Add New Product")

	fmt.Print("\nInsert New Product Name :  ")
	fmt.Scan(&name)
	fmt.Print("\nInsert New Product Brand : ")
	fmt.Scan(&brand)
	fmt.Print("\nInsert New Product Price : ")
	fmt.Scan(&price)
	fmt.Print("\nInsert New Product Unit :  ")
	fmt.Scan(&unit)

	idProduct := len(dummyData) + 1
	dummyData[idProduct] = newProduct(name, brand, price, unit)
	fmt.Println("New product created id:", idProduct)
	fmt.Println("==================================")
}

func editProduct() {
	var idProduct int

	menuHeader("Edit Product")

	fmt.Print("Masukan ID product : ")
	fmt.Scan(&idProduct)

	if product, exist := dummyData[idProduct]; !exist {
		fmt.Println("Product is not Exist")
	} else {
		var editName, editBrand string
		var editPrice, editUnit int

		fmt.Println("Product ID => ", idProduct)
		fmt.Println("type (skip) to skip this column")
		fmt.Printf("Edit name product (current %s) to :", product.GetName())
		editName = func() string {
			fmt.Scan(&editName)
			if editName == "skip" {
				editName = product.GetName()
				return editName
			}
			return editName
		}()

		fmt.Println("type (skip) to skip this column")
		fmt.Printf("Edit brand product (current %s) to :", product.GetBrand())
		editBrand = func() string {
			fmt.Scan(&editBrand)
			if editBrand == "skip" {
				editBrand = product.GetBrand()
				return editBrand
			}
			return editBrand
		}()

		fmt.Println("type (-1) to skip this column")
		fmt.Printf("Edit price product (current %d) to :", product.GetPrice())
		editPrice = func() int {
			fmt.Scan(&editPrice)
			if editPrice == -1 {
				editPrice = product.GetPrice()
				return editPrice
			}
			return editPrice
		}()

		fmt.Println("type (-1) to skip this column")
		fmt.Printf("Edit unit product (current %d) to:", product.GetUnit())
		editUnit = func() int {
			fmt.Scan(&editUnit)
			if editUnit == -1 {
				editUnit = product.GetUnit()
				return editUnit
			}
			return editUnit
		}()

		product.SetProduct(editName, editBrand, editPrice, editUnit)
		dummyData[idProduct] = product
		fmt.Println("Product ID:", idProduct, " Successfuly Edited")
		fmt.Println("==================================")
	}
}

func increaseUnit() {
	var idProduct int

	menuHeader("Increase Product Unit")

	fmt.Print("Masukan ID product : ")
	fmt.Scan(&idProduct)

	if product, exist := dummyData[idProduct]; !exist {
		fmt.Println("Product is not Exist")
	} else {
		var increasedUnit int

		fmt.Println("Product => ", product.GetName())
		fmt.Printf("Increase Unit (current %d) to :", product.GetUnit())
		increasedUnit = func() int {
			fmt.Scan(&increasedUnit)
			if increasedUnit <= 0 {
				fmt.Println("cant add negative numbers or zero")
			}
			return increasedUnit
		}()
		product.IncreaseUnit(increasedUnit)
		dummyData[idProduct] = product
		fmt.Println("Product Name:", product.GetName(), ". Successfuly Increase the unit")
		fmt.Println("==================================")
	}
}

func decreasedUnit() {
	var idProduct int

	menuHeader("Increase Product Unit")

	fmt.Print("Masukan ID product : ")
	fmt.Scan(&idProduct)

	if product, exist := dummyData[idProduct]; !exist {
		fmt.Println("Product is not Exist")
	} else {
		var decreasedUnit int

		fmt.Println("Product => ", product.GetName())
		fmt.Printf("Decrease Unit (current %d) to :", product.GetUnit())
		decreasedUnit = func() int {
			fmt.Scan(&decreasedUnit)
			if decreasedUnit <= 0 {
				fmt.Println("cant reduce negative numbers or zero")
			}
			return decreasedUnit
		}()
		product.DecreaseUnit(decreasedUnit)
		dummyData[idProduct] = product
		fmt.Println("Product Name:", product.GetName(), ". Successfuly Decrease the unit")
		fmt.Println("==================================")
	}
}

func menuHeader(breadcrumb string) {
	fmt.Println("\n==================================")
	fmt.Println("     Product Management System    ")
	fmt.Println("==================================")
	fmt.Println("        Menu - ", breadcrumb)
	fmt.Println("----------------------------------")
}

func menuBody(menu int) {
	switch menu {
	case 1:
		showAllProduct()

	case 2:
		addNewProduct()

	case 3:
		editProduct()

	case 4:
		increaseUnit()

	case 5:
		decreasedUnit()

	default:
		menuHeader("Main")
		fmt.Println("\n1. Show All Product")
		fmt.Println("2. Add New Product")
		fmt.Println("3. Edit Product")
		fmt.Println("4. Increase Product Unit")
		fmt.Println("5. Decrease Product Unit")
		fmt.Println("9. Exit")
		fmt.Println("----------------------------------")
	}

}

func main() {
	var menu int
	for {
		menuBody(menu)
		fmt.Print("Choose Menu (Back to MainMenu = 0 or 9 to Exit): ")
		fmt.Scan(&menu)
		if menu == 9 {
			break
		}
	}
}
