package models

// CreateUserRequest is what the user sends us
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"` // Validates date format YYYY-MM-DD
}

// UserResponse is what we send back (includes Age)
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age,omitempty"` // omitempty means if age is 0, it might be hidden, but here we want it shown usually
}