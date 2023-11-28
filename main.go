package main

func main() {
	random_model := Random_model_connected(1000, 5)
	random_model = applyPartialAssignment(random_model, 0, V_Partial{I, Yes})
	random_model = applyPartialAssignment(random_model, 0, V_Partial{S, High})
	TransformRecursively(random_model, 10000)
}
