package Model

type Todo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `bool:"completed"`
}

func (b *Todo) TableName() string {
	return "todo"
}
