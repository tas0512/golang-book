package main

import (
	"fmt"
)

func main() {
	//vm := NewVendingMachine(coins, items)
	vm := NewVendingMachine(0,0)

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	can := vm.SelectSD()
	fmt.Println(can) // SD
	// Buy CC(canned coffee) without exact change
	//vm.InsertCoin("T")
	//vm.InsertCoin("T")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	//can = vm.SelectCC()
	//fmt.Println(can) // CC, F, TW, O
	// Start adding change but hit coin return
	//vm.InsertCoin("T")
	//vm.InsertCoin("T")
	//vm.InsertCoin("F")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	//change := vm.CoinReturn()
	//fmt.Println(change) // T, T, F
}

type VendingMachine struct {
	c int
	i string
}

func NewVendingMachine(c int,i string) *VendingMachine {
	return &VendingMachine{c,i}
}

func (v *VendingMachine) InsertCoin(coinType string) {
	if coinType=="T" {
		v.c = v.c + 10
	}
	if coinType=="F" {
		v.c = v.c + 5
	}
	if coinType=="TW" {
		v.c = v.c + 2
	}
	if coinType=="O" {
		v.c = v.c + 1
	}
}

func (v VendingMachine) GetInsertedMoney() int {
	return v.c
}

func (v VendingMachine) SelectSD() {
	v.i = "Soft Drink"
	return v.i
}