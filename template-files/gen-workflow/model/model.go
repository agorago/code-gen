package model


type __PackageName__ struct {
	ID          string
	State string
	// Other attributes relevant to __PackageName__ type
}

func (ord *__PackageName__) GetState() string {
	return ord.State
}

func (ord *__PackageName__) SetState(newState string) {
	ord.State = newState
}


