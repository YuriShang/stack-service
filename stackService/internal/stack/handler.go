package stack

import (
	"encoding/json"
	"net/http"

	"stackService/dto"
	"stackService/utils"

	"stackService/pkg/logging"

	"github.com/gorilla/mux"
)

const (
	push = "/push"
	pop  = "/pop"
)

type Handler struct {
	Logger       logging.Logger
	StackService Service
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(push, h.Push).Methods("POST")
	router.HandleFunc(pop, h.Pop).Methods("DELETE")
}

func (h *Handler) Push(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("PUSH")

	decoder := json.NewDecoder(r.Body)
	createData := dto.StackData{}
	err := decoder.Decode(&createData)

	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Failed to parse passed data")
		return
	} else {
		if err := utils.Validate(&createData); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	data, err := h.StackService.Push(r.Context(), createData)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.ResponseWithJSON(w, http.StatusCreated, data)
	}
}

func (h *Handler) Pop(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("POP")
	data, err := h.StackService.Pop(r.Context())
	if err != nil && data.Data == 0 {
		utils.ResponseWithError(w, http.StatusNotFound, "Stack is empty!")
	} else if err != nil && data.Data != 0 {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.ResponseWithJSON(w, http.StatusOK, data)
	}
}
