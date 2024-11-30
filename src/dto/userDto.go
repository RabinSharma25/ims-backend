package dto

type AddUserDetailsReqDto struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UpdateUserDetailsReqDto struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type UpdateUserEmailReqDto struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type GetUserDetailsResDto struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
