package main

import "fmt"

func main()  {
	sayMsg("Hello World", "Nothing", 3)
}

func sayMsg(msg, next string, idx int)  {
	fmt.Println(msg)
}

func (greet *Greet) above()  {

}

type Greet struct {
	name string
	hello [3]int


}