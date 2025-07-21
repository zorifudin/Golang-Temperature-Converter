package main

import "fmt"

func main() {
	fmt.Println("Masukkan kalkulator suhu yang ingin dipakai")
	fmt.Println("1. Celcius to Fahrenheit")
	fmt.Println("2. Celcius to Kelvin")
	fmt.Println("3. Fahrenheit to celcius")
	fmt.Println("4. Fahrenheit to Kelvin")
	fmt.Println("5. Kelvin to Celcius")
	fmt.Println("6. Kelvin to Fahrenheit")
	fmt.Println("Input your choice: ")

	var pilihan int
	fmt.Scanf("%d\n", &pilihan)

	for pilihan < 1 || pilihan > 6 {
		fmt.Println("Calculator not available, input your choice: ")
		fmt.Scanf("%d\n", &pilihan)
	}

	var suhu float64
	fmt.Println("Input first temperature:")
	fmt.Scanf("%f\n", &suhu)

	var suhuAkhir float64
	if pilihan == 1 {
		suhuAkhir = CelciusToFahrenheit(suhu)
	} else if pilihan == 2 {
		suhuAkhir = CelciusToKelvin(suhu)
	} else if pilihan == 3 {
		suhuAkhir = FahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		suhuAkhir = FahrenheitToKelvin(suhu)
	} else if pilihan == 5 {
		suhuAkhir = KelvinToCelcius(suhu)
	} else {
		suhuAkhir = KelvinToFahrenheit(suhu)
	}

	fmt.Printf("Suhu akhir hasil konversi adalah: %.2f\n", suhuAkhir)
}

func CelciusToFahrenheit(suhuCelcuis float64) float64 {
	suhuFahrenheit := 9.0/5.0*suhuCelcuis + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelcuis float64) float64 {
	suhuKelvin := suhuCelcuis + 273.15
	return suhuKelvin
}

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	suhuCelcuis := (suhuFahrenheit - 32) * 5.0 / 9.0
	return suhuCelcuis
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	suhuKelvin := (suhuFahrenheit-32)*5.0/9.0 + 273.15
	return suhuKelvin
}

func KelvinToCelcius(suhuKelvin float64) float64 {
	suhuCelcuis := (suhuKelvin - 273.15)
	return suhuCelcuis
}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	suhuFahrenheit := (suhuKelvin-273.15)*9.0/5.0 + 32
	return suhuFahrenheit
}
