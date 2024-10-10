package main
import "fmt"

func main() {
	var celsius float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := celsius*9/5 + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}