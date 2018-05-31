package myAgency

func AsyncTask() {

	customer := GetCustomer()
	destinations := GetDestinations(customer)
	var myInfo [30]MyInfo

	//Test Array Length is 30
	quotes := [30]chan Quoting{}
	weathers := [30]chan Weather{}

	//use chan
	for i := range weathers {
		weathers[i] = make(chan Weather)
	}

	for i := range quotes {
		quotes[i] = make(chan Quoting)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		// start task by async mode
		go func() {
			quotes[i] <- GetQuote(d)
		}()

		go func() {
			weathers[i] <- GetWeather(d)
		}()
	}

	for index, d := range destinations {
		myInfo[index] = MyInfo{Destinations: d, Quote: <-quotes[index], Weather: <-weathers[index]}
	}
}
