package usersdto

type CreateUserRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	NPP           int    `json:"npp"`
	NPPSupervisor int    `json:"npp_supervisor"`
	Password      string `json:"password"`
}
