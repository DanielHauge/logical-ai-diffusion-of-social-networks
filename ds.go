package main

type Relation struct {
	From int
	To   int
}

const (
	Yes = iota
	No
)

const (
	Acted = iota
	NotActed
	Undecided
)

const (
	High = iota
	Medium
	Low
	None
)

// Properties
const (
	I = iota // Involvement
	A        // Acted
	S        // Susceptibility
)

// Full assignment V_full for an agent
type V struct {
	I int // Involvement: {Yes, No}
	A int // Acted: {Acted, NotActed, Undecided}
	S int // Suspecptibility: {High, Medium, Low, None}
}

type V_Partial struct {
	Property int // {I, A, S}
	Value    int // {Yes, No}, {Acted, NotActed, Undecided}, {High, Medium, Low, None}
}

type Relations []Relation

// Binary relations: ≍
type RelationsMap map[int]Relations

// Assignments: v
type Assignments map[int]V

type Model struct {
	R     RelationsMap
	V     Assignments
	Count int
}
