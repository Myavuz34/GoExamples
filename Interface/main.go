package main

import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Araç bilgileri : \n" + c.Information())
	fmt.Println("\n")

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "Çalışıyor"
	} else {
		msg = "Çalışmıyor"
	}
	fmt.Println("Araç : " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "Durdu"
	} else {
		msg = "Durmadı araç gidiyor"
	}
	fmt.Println("Araç : " + msg + ".")
}

func main() {

	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true
	fmt.Println(ferr.Information())

	bmw := new(Bmw)
	bmw.Brand = "BMW"
	bmw.Model = "520i"
	bmw.Color = "BLACK"
	bmw.Speed = 260
	bmw.Price = 0.3
	fmt.Println(bmw.Information())

	CarExecute(ferr)
	CarExecute(bmw)

}

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color:" + x.Color + "\n" + "\t" + "Speed :" + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price :" + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}

	return ret
}

type Bmw struct {
	Car
	SpecialProduction
}

func (_ Bmw) Run() bool {
	return true
}

func (_ Bmw) Stop() bool {
	return true
}

func (x *Bmw) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color" + "\n" + "\t" + "Speed :" + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price :" + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"

	return ret
}

type Mercedes struct {
	Car
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color" + "\n" + "\t" + "Speed :" + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price :" + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"

	return ret
}
