package main

import (
	"strconv"
	"fmt"
)

func main() {
	//vm := NewVendingMachine(coins, items)
	vm := NewVendingMachine(0,"")

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	can := vm.SelectSD()
	fmt.Println(can) // SD
	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O
	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F
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

func (v *VendingMachine) SelectSD() string {
	v.i = "SD"
	fmt.Println("select ",v.i)
	num := v.c-18
	returnCoin := strconv.Itoa(num)
	v.c = v.c-18
	//message := calculateChange(v.c)

	return "Return:"+v.i+" Change Amount:"+returnCoin
}

func (v *VendingMachine) SelectCC() string {
	v.i = "CC"
	fmt.Println("select ",v.i)
	num := v.c-12
	returnCoin := strconv.Itoa(num)
	v.c = v.c-12
	//message := calculateChange(v.c)

	return "Return:"+v.i+" Change Amount:"+returnCoin
}

func (v *VendingMachine) CoinReturn() string {
	fmt.Println("Return Coin")
	num := v.c
	returnCoin := strconv.Itoa(num)
	v.c = v.c-num

	return "Change Amount:"+returnCoin
}

/*func calculateChange(n int) string {
	textmsg := ""
	for n>0 {
		if n>10 {
			textmsg = textmsg + "T"
			n = n - 10
		}
		if n>5 {
			textmsg = textmsg + "F"
			n = n - 5
		}
		if n>2 {
			textmsg = textmsg + "TW"
			n = n - 2
		}
		if n>1 {
			textmsg = textmsg + "O"
			n = n - 1
		}
	}
	return textmsg
}*/