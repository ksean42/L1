package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////////

// Адаптируемый неизменяемый объект
type TermometerF interface {
	getTempF(float32) float32
}

type TermometerFahrenheit struct {
}

// Принимает температуру по фаренгейту, возвращает температуру по цельсию
func (tf *TermometerFahrenheit) getTempF(fahrenheit float32) float32 {
	return (fahrenheit - 32) / 1.8
}

//////////////////////////////////////////
// Наш объект и интерфейс
type TermometerK interface {
	getTemp(float32) float32
}

type TermometerKelvin struct {
}

// Принимает температуру по кельвину, возвращает температуру по цельсию
func (tk *TermometerKelvin) getTemp(kelvin float32) float32 {
	return kelvin - 273
}

///////////////////////
// Адаптер интерфейса TermometerF к TermometerK
type Adapter struct {
	*TermometerFahrenheit
}

func (a *Adapter) getTemp(temp float32) float32 {
	return a.getTempF(temp)
}

func main() {
	var fahrenheit TermometerF = &TermometerFahrenheit{}
	fmt.Println(fahrenheit.getTempF(-30))

	var kelvin TermometerK = &TermometerKelvin{}
	fmt.Println(kelvin.getTemp(-30))

	var adapter TermometerK = &Adapter{&TermometerFahrenheit{}}
	printTemp(adapter)
	printTemp(kelvin)

}

func printTemp(t TermometerK) {
	fmt.Println("Temp:", t.getTemp(-30))
}
