package es_ser

type Option struct {
	Page   int    `form:"page"`
	Key    string `form:"key"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Fields []string
	Tag    string `form:"tag"`
}

type SortFiled struct {
	Field     string
	Ascending bool
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}

	return (o.Page - 1) * o.Limit
}
