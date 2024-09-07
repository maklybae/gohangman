package domain

type State int

const (
	Initial State = iota
	Head
	Body
	LeftArm
	RightArm
	LeftLeg
	RightLeg
)
