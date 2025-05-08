package request

type IdNumberRequest struct {
	Id int `json:"id" validate:"required"`
}
