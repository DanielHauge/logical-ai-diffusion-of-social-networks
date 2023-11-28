package main

func neighborsCondition(model Model, p int, property int, value int, proposition float64) bool {
	neighborsWithCondition := 0
	for _, r := range model.R[p] {
		switch property {
		case I:
			if model.V[r.To].I == value {
				neighborsWithCondition++
			}
		case A:
			if model.V[r.To].A == value {
				neighborsWithCondition++
			}
		case S:
			if model.V[r.To].S == value {
				neighborsWithCondition++
			}
		}
	}
	if proposition == 0.0 {
		return neighborsWithCondition > 0
	}
	return float64(neighborsWithCondition)/float64(len(model.R[p])) >= proposition
}

func applyPartialAssignment(model Model, p int, v V_Partial) Model {
	f := model.V[p]
	switch v.Property {
	case I:
		f.I = v.Value
	case A:
		f.A = v.Value
	case S:
		f.S = v.Value
	default:
		panic("Invalid property")
	}

	model.V[p] = f
	return model
}

var (
	JustOneNeighbor      = 0.0 // This is just a constant used to infer a condition for just one neightbor
	InvolvementMedium    = 1.0 / 3.0
	InvolvementLow       = 2.0 / 3.0
	SusceptibilityMedium = 1.0 / 2.0
	SusceptibilityLow    = 3.0 / 4.0
)

func transform(model Model) (Model, bool) {
	changed := false
	newModel := DeepCopyModel(model)
	for i := 0; i < model.Count; i++ {
		if model.V[i].I == No && model.V[i].S == High && neighborsCondition(model, i, I, Yes, JustOneNeighbor) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{I, Yes})
			changed = true
		}
		if model.V[i].I == No && model.V[i].S == Medium && neighborsCondition(model, i, I, Yes, InvolvementMedium) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{I, Yes})
			changed = true
		}
		if model.V[i].I == No && model.V[i].S == Low && neighborsCondition(model, i, I, Yes, InvolvementLow) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{I, Yes})
			changed = true
		}
		if model.V[i].I == No && model.V[i].A == Undecided && model.V[i].S == None && neighborsCondition(model, i, I, Yes, 0.0) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, NotActed})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == High {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, Acted})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == Medium && neighborsCondition(model, i, A, Acted, SusceptibilityMedium) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, Acted})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == Medium && neighborsCondition(model, i, A, NotActed, SusceptibilityMedium) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, NotActed})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == Low && neighborsCondition(model, i, A, NotActed, SusceptibilityLow) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, NotActed})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == Low && neighborsCondition(model, i, A, Acted, 1.0-SusceptibilityLow) {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, Acted})
			changed = true
		}
		if model.V[i].I == Yes && model.V[i].A == Undecided && model.V[i].S == None {
			newModel = applyPartialAssignment(newModel, i, V_Partial{A, NotActed})
			changed = true
		}
	}

	return newModel, changed
}

func TransformRecursively(model Model, k int) (Model, bool, int) {
	var changed bool
	for i := 0; i < k; i++ {
		model, changed = transform(model)
		PrintCounts(i, model)
		if !changed {
			return model, changed, i
		}
	}
	return model, changed, k
}
