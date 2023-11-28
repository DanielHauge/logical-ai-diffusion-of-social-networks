package main

import "math/rand"

func allocate_model() Model {
	// Allocate new relations ≍ and assignments v
	r := make(RelationsMap)
	v := make(Assignments)

	// Allocate a new model
	result := Model{
		R: r,
		V: v,
	}
	return result
}

func getRandomNumberExcludingN(exclude int, max int) int {
	var randomNumber int

	for {
		randomNumber = rand.Intn(max)
		if randomNumber != exclude {
			break
		}
	}

	return randomNumber
}

func isModelConnected(model Model) bool {
	marked := make([]bool, model.Count)
	var dfs func(int)
	dfs = func(i int) {
		marked[i] = true
		for _, r := range model.R[i] {
			if !marked[r.To] {
				dfs(r.To)
			}
		}
	}
	dfs(0)
	for i := 0; i < model.Count; i++ {
		if !marked[i] {
			return false
		}
	}

	return true
}

func appendRandomRelation(model Model) Model {
	return model
}

func appendRandomRelationOn(model Model, i int, people_count int) Model {
	other := getRandomNumberExcludingN(i, people_count)
	r_i := Relation{
		From: i,
		To:   other,
	}
	model.R[i] = append(model.R[i], r_i)
	r_j := Relation{
		From: other,
		To:   i,
	}
	model.R[other] = append(model.R[other], r_j)

	return model
}

func random_model(people_count int, relations_per_people_count int) Model {
	model := allocate_model()
	model.Count = people_count

	for i := 0; i < people_count; i++ {
		// Assign random features
		a := V{
			I: No,           // Corresponds to Yes/No
			A: Undecided,    // Corresponds to Acted/NotActed/Undecided
			S: rand.Intn(4), // Corresponds to High/Medium/Low/None
		}
		model.V[i] = a

		for j := 0; j < relations_per_people_count; j++ {
			model = appendRandomRelationOn(model, i, people_count)
		}

	}

	return model
}

func Random_model_connected(people_count int, relations_per_people_count int) Model {
	model := random_model(people_count, relations_per_people_count)

	for !isModelConnected(model) {
		model = appendRandomRelation(model)
	}

	return model
}

func DeepCopyModel(model Model) Model {
	var newModel Model
	newModel.Count = model.Count
	newModel.V = make(Assignments)
	newModel.R = make(RelationsMap)

	for i := 0; i < model.Count; i++ {
		newModel.V[i] = model.V[i]
		newModel.R[i] = append(model.R[i][:0:0], model.R[i]...)
	}

	return newModel
}

func Print_model(model Model) {
	for i := 0; i < model.Count; i++ {
		println("Person", i, ": I =", model.V[i].I, " A =", model.V[i].A, " S =", model.V[i].S)
		for _, r := range model.R[i] {
			println("Relation:", r.From, "≍", r.To)
		}
		println()
	}
}
