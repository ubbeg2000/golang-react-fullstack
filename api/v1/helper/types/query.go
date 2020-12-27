package types

type UserQuery struct {
	Limit int   `query:"limit"`
	Page  int   `query:"page"`
	Role  int   `query:"role"`
	Ids   []int `query:"ids"`
}

func DefaultUserQuery() *UserQuery {
	return &UserQuery{20, 1, 0, []int{}}
}
