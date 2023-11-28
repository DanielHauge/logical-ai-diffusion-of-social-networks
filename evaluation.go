package main

func countFeature(model Model, property int, value int) int {
	count := 0
	for _, v := range model.V {
		switch property {
		case I:
			if v.I == value {
				count++
			}
		case A:
			if v.A == value {
				count++
			}
		case S:
			if v.S == value {
				count++
			}
		}
	}
	return count
}

func CountInvolved(model Model) (int, int) {
	return countFeature(model, I, Yes), countFeature(model, I, No)
}

func CountActed(model Model) (int, int, int) {
	return countFeature(model, A, Acted), countFeature(model, A, NotActed), countFeature(model, A, Undecided)
}

func CountSusceptible(model Model) (int, int, int, int) {
	return countFeature(model, S, High), countFeature(model, S, Medium), countFeature(model, S, Low), countFeature(model, S, None)
}

func PrintCounts(i int, model Model) {
	Yesses, Nos := CountInvolved(model)
	Acted, NotActed, Undecided := CountActed(model)
	// High, Medium, Low, None := CountSusceptible(model)

	println(i, Yesses, Nos, Acted, NotActed, Undecided)
}
