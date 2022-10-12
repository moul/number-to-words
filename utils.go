package ntw

func integerToTriplets(number int) []int {
	triplets := []int{}

	for number > 0 {
		triplets = append(triplets, number%1000)
		number = number / 1000
	}

	return triplets
}

func integerToDHybrid(number int) []int {
	hybrid := []int{}

	startHybrid := false
	for number > 0 {
		if !startHybrid{
			hybrid = append(hybrid, number%1000)
			number = number / 1000
			startHybrid = true
		} else {
			hybrid = append(hybrid, number%100)
			number = number / 100
		}
		
	}

	return hybrid
}
