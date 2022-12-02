package handlers

import (
	authdto "backend/dto/auth"
	dto "backend/dto/result"
	usersdto "backend/dto/users"
	"backend/models"
	"backend/pkg/bcrypt"
	jwtToken "backend/pkg/jwt"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) RegisterSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(usersdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	users := models.Users{
		Nama:          request.Nama,
		Email:         request.Email,
		Password:      password,
		NPP:           time.Now().UnixNano() / int64(time.Millisecond),
		NPPSupervisor: 0,
	}

	data, _ := h.AuthRepository.RegisterSV(users)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	nppSV := int64(adminInfo["nppSV"].(float64))

	if nppSV == 0 {
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: "Prohibited!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	request := new(usersdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	users := models.Users{
		Nama:          request.Nama,
		Email:         request.Email,
		Password:      password,
		NPP:           time.Now().UnixNano() / int64(time.Millisecond),
		NPPSupervisor: nppSV,
	}

	data, _ := h.AuthRepository.RegisterUser(users)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.AuthRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	users := models.Users{
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.AuthRepository.Login(users.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: "Email not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, data.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: "Wrong password!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	gnrtToken := jwt.MapClaims{}
	gnrtToken["id"] = data.ID
	gnrtToken["exp"] = time.Now().Add(time.Hour * 3).Unix()
	gnrtToken["nppSV"] = data.NPP
	gnrtToken["nama"] = data.Nama

	token, err := jwtToken.GenerateToken(&gnrtToken)
	if err != nil {
		fmt.Println("Unauthorize")
		return
	}

	AuthResponse := authdto.AuthResponse{
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
		Token:    token,
		ID:       data.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Message: "Success Get Data", Data: AuthResponse}
	json.NewEncoder(w).Encode(response)
}
