package main

import (
	"fmt"
	"os"
)

const (
	usage         = "[Username] [Password]"
	errUser       = "Access denied for %q.\n"
	errPwd        = "Invalid Password for %q. \n"
	accessOk      = "Access Granted %q \n"
	user, newuser = "jack", "sagar"
	pass, newpass = "1888", "1879"
)

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Printf(usage)
		return
	}
	u, p := args[1], args[2]
	if u != user && u != newuser {
		fmt.Printf(errUser, u)
		return
	} else if p == pass && u == user {
		fmt.Printf(accessOk, u)
	} else if p == newpass && u == newuser {
		fmt.Printf(accessOk, u)
	} else {
		fmt.Printf(errPwd, p)
		return
	}
}
