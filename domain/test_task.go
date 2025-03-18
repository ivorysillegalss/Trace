package domain

type TestTaskContext struct {
	BusinessId int
	Message    string
}

func (t *TestTaskContext) Data() {}
