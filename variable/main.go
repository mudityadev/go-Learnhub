package main

import "fmt"

var LoginToken string = "mkrihjtl@r0811" // public variable

func main() {
	var username string = "mudityadev"
	fmt.Println(username)
	fmt.Printf("[INFO] Type of the Variable : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("[INFO] Type of the Variable : %T \n", isLoggedIn)

	var valueFloat float32 = 21.0
	fmt.Println(valueFloat)
	fmt.Printf("[INFO] Type of the Variable : %T \n", valueFloat)

	helloSexy := "Muditya Raghav"
	fmt.Println(helloSexy)

	// LoginToken Public Global Var
	fmt.Println(LoginToken)
	fmt.Printf("[INFO] Type of the Variable : %T \n", LoginToken)
}
