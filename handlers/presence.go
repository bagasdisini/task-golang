package handlers

import (
	e_presencedto "backend/dto/e_presence"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerPresence struct {
	PresenceRepository repositories.PresenceRepository
}

func HandlerPresence(PresenceRepository repositories.PresenceRepository) *handlerPresence {
	return &handlerPresence{PresenceRepository}
}

func (h *handlerPresence) ShowPresences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Presences, err := h.PresenceRepository.ShowPresences()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: Presences}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPresence) GetPresenceByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var Presence models.EPresence
	Presence, err := h.PresenceRepository.GetPresenceByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: Presence}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPresence) CreatePresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	adminId := int(adminInfo["id"].(float64))

	request := new(e_presencedto.CreatePresenceRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Presence := models.EPresence{
		Type:      request.Type,
		IDUser:    adminId,
		IsApprove: "false",
		Tanggal:   request.Tanggal,
		Waktu:     request.Waktu,
	}

	Presence, err = h.PresenceRepository.CreatePresence(Presence)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Presence, err = h.PresenceRepository.GetPresenceByID(Presence.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: Presence}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPresence) UpdatePresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := e_presencedto.UpdatePresenceRequest{
		IsApprove: r.FormValue("isApprove"),
	}

	Presence := models.EPresence{}

	if request.IsApprove != "" {
		Presence.IsApprove = request.IsApprove
	}

	data, err := h.PresenceRepository.UpdatePresence(Presence, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Message: "Success Get Data", Data: data}
	json.NewEncoder(w).Encode(response)
}
