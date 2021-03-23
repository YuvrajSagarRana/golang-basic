package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}
func main()  {
	var N int
	fmt.Println("Please enter the number")
	fmt.Scanf("%d", &N)
	fmt.Println(N)
}

func printNode(n *Node)  {
	fmt.Printf("Value is %v", n.Value)
	if n.Left != nil {
		fmt.Printf("Left Value is %v", n.Left.Value)
	}

	if n.Right != nil {
		fmt.Printf("Right Value is %v", n.Right.Value)
	}
	fmt.Println()
}
