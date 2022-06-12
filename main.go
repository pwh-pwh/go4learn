package main

import (
	"booking-app/tool"
	"fmt"
	"strconv"
)


type User struct {
	userName string
	password string
	age int64
}



var a = 10

var userList = make([]User,0)

func main() {

	user:= User {
		userName: "pwh",
		password: "root",
		age: 20,
	}
	userList = append(userList,user)
	for _,u := range userList {
		fmt.Println(u)
	}

	num := strconv.FormatUint(uint64(11),10)
	fmt.Println(num)
	var userData = make(map[string]string)
	userData["userName"] = "coderpwh"
	userData["type"] = "coder"
	fmt.Println(userData)
	fmt.Println(tool.Aan)
	fmt.Println(tool.Add(11, 21))
	var str = "stt"
	fmt.Println("Hello world  ", str, "to learn")

	var userName string
	var userTickets int
	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)
	fmt.Println(userName)
	var n, e = fmt.Scan(&userName)
	fmt.Printf("n is %v e is %v", n, e)
	userTickets = 14
	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)
	fmt.Printf("name type:%T tk type:%T\n", userName, userTickets)

	var booking [10]string
	fmt.Printf("array %v type %T len %v\n", booking, booking, len(booking))

	var bslick []string
	bslick = append(bslick, "test")
	fmt.Printf("type: %T value: %v len: %v getfirst %v\n", bslick, bslick, len(bslick), bslick[0])

	ct := 2
	switch ct {
	case 1, 3:
		fmt.Println("1 or 3")
	case 2, 4:
		fmt.Println("2 or 4")

	default:
		fmt.Println(" default ")

	}

}
