package Models

type Table struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Capacity   int `json:"capacity"`
	Location   string `json:"location"`
}

func (b *Table) TableName() string {
	return "table"
}
