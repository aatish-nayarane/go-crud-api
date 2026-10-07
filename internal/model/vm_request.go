package model

type CreateVMRequest struct {
	Name string `json:"name"`
	CPU  int64  `json:"cpu"`
	RAM  int64  `json:"ram"`
}

type UpdateVMRequest struct {
	Name string `json:"name"`
	CPU  int64  `json:"cpu"`
	RAM  int64  `json:"ram"`
}
