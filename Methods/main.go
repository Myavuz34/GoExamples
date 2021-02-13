package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Kullanıcı oluşturma V1
	fmt.Println("Kullanıcı Oluşturma V1")

	// user1 := &User{
	// 	ID:        1,
	// 	FirstName: "Mustafa",
	// 	LastName:  "Yavuz",
	// 	UserName:  "Myavuz",
	// 	Age:       30,
	// 	Pay: &Payment{
	// 		Salary: 5676,
	// 		Bonus:  1000,
	// 	},
	// }
	// fmt.Println(user1)
	// fmt.Println(user1.GetFullName())
	// fmt.Println(user1.GetSalary())
	// fmt.Println("Maaş : " + strconv.FormatFloat(user1.GetSalary(), 'g', -1, 64))

	//Kullanıcı oluşturma V2
	fmt.Println("Kullanıcı Oluşturma V2")
	user2 := NewUser()
	user2.FirstName = "Mustafa"
	user2.LastName = "Yavuz"
	user2.UserName = "Myavuz"
	user2.Age = 30
	user2.Pay = &Payment{Salary: 755, Bonus: 656}

	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetSalary())
	fmt.Println(("Maaş: " + strconv.FormatFloat(user2.GetSalary(), 'g', -1, 64)))

}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

func (u User) GetFullName() string {

	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetSalary() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
