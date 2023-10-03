package enum

// Category should map to category_type_enum
type Category string

const (
	Electricity Category = "electricity"
	Water       Category = "water"
	Postpaid    Category = "postpaid"
	Loan        Category = "loan"
)

// AllCategories returns all possible values for Category
func (c Category) All() []string {
	return []string{
		string(Electricity),
		string(Water),
		string(Postpaid),
		string(Loan),
	}
}

func (c Category) Name() string {
	return "category_name_enum"
}
