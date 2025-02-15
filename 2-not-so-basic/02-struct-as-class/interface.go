package main

import(
	"fmt"
)

// interface is a collection of method signatures
type Animal interface {
	Born()
	Eat(food string)
	Talk()string
	HuntIt(target string)bool

	// numberFeet int8 // interface doesnt have attributes, only methods
}

// we don't explicity relates interface with struct, we just create struct that has these methods implemented 
// and the compiler knows that this struct is this interface
// the struct needs to implement ALL methods
type Cat struct{
	name string
}
type Whale struct{
	race string
	wheight int
}
type Plant struct{ // plant is not a Animal interface because doesnt implements all functions, just some of them
	color string
}

func (c Cat)Born(){ fmt.Println("A cat borns and walks ways") }
func (c Cat)Eat(food string){ fmt.Println("A cat eats a", food, "and it's dirty") }
func (c Cat)Talk()string{ return "miau" }
func (c Cat)HuntIt(target string)bool{ return target == "bird" || target == "insect" || target == "mouse" }

func (w Whale)Born(){ fmt.Println("A whale borns and swims away") }
func (w Whale)Eat(food string){ fmt.Println("A whale eats", food, "and it's clean") }
func (w Whale)Talk()string{ return "wooooooon" }
func (w Whale)HuntIt(target string)bool{ return target == "fish" }

func (p Plant)Born(){ fmt.Println("A plant borns") }
func (p Plant)HuntIt(target string)bool{ return false }

// function that works with any kind of animal
func Hunting(a Animal, food string){
	if(a.HuntIt(food)){
		a.Eat(food)
	}else{
		fmt.Println("I'm still hunting")
	}
}

func MyName(a Animal){
	switch animal := a.(type) {
	case Cat:
		fmt.Println("My name is ", animal.name)
	case Whale:
		fmt.Println("My name is ", animal.race)
	case Animal:
		fmt.Println("I'm a undefined animal")
	default:
		fmt.Println("it's not an animal")
	}
}

func main(){
	cat := Cat{"bichano"}
	whale := Whale{"jubarte", 100}
	plant := Plant{"laranjeira"}

	cat.Born()
	cat.Eat("rat")
	fmt.Println( cat.Talk() )
	fmt.Println( "A cat hunt it? ", cat.HuntIt("bird") )

	whale.Born()
	whale.Eat("fish")
	fmt.Println( whale.Talk() )
	fmt.Println( "A whale hunt it? ", whale.HuntIt("bird") )

	plant.Born()
	fmt.Println( "A plant hunt it? ", plant.HuntIt("bird") )

	fmt.Println("----------- checking who is animal")

	if _, pass := interface{}(cat).(Animal); pass {
		fmt.Println(cat, "is an Animal")
	}else{
		fmt.Println(cat, "is not an Animal")
	}

	if _, pass := interface{}(whale).(Animal); pass {
		fmt.Println(whale, "is an Animal")
	}else{
		fmt.Println(whale, "is not an Animal")
	}
	
	if _, pass := interface{}(plant).(Animal); pass {
		fmt.Println(plant, "is an Animal")
	}else{
		fmt.Println(plant, "is not an Animal")
	}

	fmt.Println("----------- detecting the kind of animal")
	MyName(cat)
	MyName(whale)
	// MyName(plant) // it doesnt work because plant is not an animal

	fmt.Println("----------- calling a function that accepts any implementation")

	Hunting(cat, "mouse")
	Hunting(whale, "mouse")
	// Hunting(plant, "mouse") // it doesnt work because plant is not an animal
	
}