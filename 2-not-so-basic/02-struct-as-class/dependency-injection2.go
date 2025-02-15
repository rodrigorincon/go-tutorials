package main

import(
	"fmt"
	"math"
)

//// motor interface and 2 implementations 
type Motor interface{
	GetHorsePower()int // define the motor power
	OilSpent(kmDriven int)int // define how much km/liter
	Move(minutes int)int // define how fast the motor is
}
type BmwV8Motor struct{ // powerful motor, runs fast and high oil consume
	horsePower int
}
func (m BmwV8Motor) GetHorsePower()int{
	return m.horsePower
}
func (m BmwV8Motor) OilSpent(kmDriven int)int{
	eficience := float64(m.horsePower / 23)
	return int(math.Ceil( float64(kmDriven) / eficience)) // 7km per lter
}
func(m BmwV8Motor) Move(minutes int)int{
	return minutes * m.horsePower/60 // 3km per min
}

type RenaultV1Motor struct{ // regular motor, runs slow and small oil consume
	horsePower int
}
func (m RenaultV1Motor) GetHorsePower()int{
	return m.horsePower
}
func (m RenaultV1Motor) OilSpent(kmDriven int)int{
	eficience := float64(m.horsePower / 10)
	return int(math.Ceil( float64(kmDriven) / eficience)) // 12km per liter
}
func(m RenaultV1Motor) Move(minutes int)int{
	return minutes * m.horsePower/120 // 1km per min
}

//// vehicle interace and 2 implementations
type Vehicle interface{
	Drive(minutes int)int // how much km moved considering the vehicle weight
	SetMotor(m Motor)
	ReportTravel(minutes int)
}
type BmwX7 struct {
	m Motor // dependency injection starts here
}
func (v BmwX7) Drive(minutes int)int{
	return int(math.Ceil( float64(v.m.Move(minutes)) / 1.1)) // lighweight
}
func (v *BmwX7) SetMotor(m Motor){
	v.m = m
}
func (v BmwX7) ReportTravel(minutes int){
	spaceMoved := v.Drive(minutes)
	liters := v.m.OilSpent(spaceMoved)
	fmt.Println("you moved", spaceMoved, "km and spent", liters, "liters of oil in a super comfort car.")
}
type Sandero struct {
	m Motor
}
func (v Sandero) Drive(minutes int)int{
	return int(math.Ceil( float64(v.m.Move(minutes)) / 2)) // heavyhweight
}
func (v *Sandero) SetMotor(m Motor){
	v.m = m
}
func (v Sandero) ReportTravel(minutes int){
	spaceMoved := v.Drive(minutes)
	liters := v.m.OilSpent(spaceMoved)
	fmt.Println("you moved", spaceMoved, "km and spent", liters, "liters of oil.")
}


func main(){
	sandero := Sandero{}
	sandero.SetMotor(RenaultV1Motor{horsePower: 120})
	sandero.ReportTravel(10)

	x7 := BmwX7{}
	x7.SetMotor(BmwV8Motor{horsePower: 180})
	x7.ReportTravel(10)
}