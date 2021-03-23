package main

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
}
func main(){
// can inherit the property from one struct to another
 bird := Bird {}
 bird.Name = "Sparrow"
 bird.Origin = "Asia"
 bird.CanFly = true

}
