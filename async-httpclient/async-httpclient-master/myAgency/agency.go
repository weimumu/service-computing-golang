package myAgency

import "time"

type Quoting struct{}
type Customer struct{}
type Weather struct{}
type Destination struct{}

type MyInfo struct {
	Destinations Destination
	Quote        Quoting
	Weather      Weather
}

func GetDestinations(customers Customer) [30]Destination {
	time.Sleep(300 * time.Millisecond)
	return [30]Destination{}
}

func GetWeather(destinations Destination) Weather {
	time.Sleep(350 * time.Millisecond)
	return Weather{}
}

func GetCustomer() Customer {
	time.Sleep(150 * time.Millisecond)
	return Customer{}
}

func GetQuote(destination Destination) Quoting {
	time.Sleep(200 * time.Millisecond)
	return Quoting{}
}
