package myAgency

func SyncTask() {
	customers := GetCustomer()
	destinations := GetDestinations(customers)

	var myInfo [30]MyInfo
	// start task by sync mode
	for i, d := range destinations {
		weather := GetWeather(d)
		quote := GetQuote(d)
		myInfo[i] = MyInfo{Destinations: d, Quote: quote, Weather: weather}
	}

}
