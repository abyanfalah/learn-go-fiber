package general

type testRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"min=3,max=5"`
}
