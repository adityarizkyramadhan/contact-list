package response

type FindAll struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}
