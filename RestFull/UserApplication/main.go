package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""
	mux := http.NewServeMux()

	//SingUp
	mux.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("passwordConfirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "Error code is -10 : There is empt data")
		}

		if pwd == pwdConfirm {
			fmt.Fprint(w, "kayıt başarılı")

		} else {
			fmt.Println(w, "Password info must be the same")
		}

	})

	//Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helper.IsEmpty("email")
		pwdCheck := helper.IsEmpty("password")

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is a empty data")
			return
		}

		dbPwd := "12356"
		dbEmail := "myavuz53@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprint(w, "Login susccesfull")
		} else {
			fmt.Fprintf(w, "Login failed")
		}
	})

	http.ListenAndServe(":8080", mux)
}

// for k, v := range r.Form {
// 	fmt.Printf("%s =  %s\n", k, v)
// }
