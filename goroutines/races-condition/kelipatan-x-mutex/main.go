package main

import (
	"fmt"
	"sync"
	"time"
)

var start time.Time

type Kelipatan struct {
	Current    int
	KelipatanX int
	m          sync.Mutex
}

func (k *Kelipatan) GetNumber() int {
	k.m.Lock()
	defer k.m.Unlock()

	return k.Current
}

func (k *Kelipatan) GetKelipatan() int {
	k.m.Lock()
	defer k.m.Unlock()

	return k.KelipatanX
}

func (k *Kelipatan) Next(next int) {
	k.m.Lock()
	defer k.m.Unlock()

	nextNumber := k.KelipatanX * next
	k.Current += nextNumber
}

func (k *Kelipatan) SetKelipatan(x int) {
	k.m.Lock()
	defer k.m.Unlock()

	k.KelipatanX = x
}

func IsPositive(number int) bool {
	if number < 1 {
		fmt.Println("Hanya dapat menerima bilangan positif")
	}
	return true
}

func main() {
	var menu int
	kelipatanX := &Kelipatan{}

	for {
		fmt.Println("\n========Bilangan Kelipatan========")
		fmt.Println("Bilangan = ", kelipatanX.GetNumber(), "Kelipatan = ", kelipatanX.GetKelipatan())
		fmt.Println("==================================")
		fmt.Println("||             Menu              ||")
		fmt.Println("==================================")
		fmt.Println("|| 1. input Kelipatan            ||")
		fmt.Println("|| 2. input bilangan selanjutnya ||")
		fmt.Println("==================================")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		switch menu {
		case 0:
			break
		case 1:
			var xKelipatan int
			fmt.Print("Masukan kelipatan : ")
			fmt.Scan(&xKelipatan)
			go func() {
				kelipatanX.SetKelipatan(xKelipatan)
			}()
			time.Sleep(1 * time.Second)
		case 2:
			var next int
			fmt.Print("Masukan jumlah berikutnya : ")
			fmt.Scan(&next)
			go func() {
				kelipatanX.Next(next)
			}()
			time.Sleep(1 * time.Second)
		}
	}
}
