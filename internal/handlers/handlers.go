package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) CreateUser(req http.Request, res http.ResponseWriter) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		s.logger.Printf("Error reading request body: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	body := CreateUserRequest{}
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		s.logger.Printf("Error unmarshalling request body: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CreateUserResponse{Data: map[string]string{"status": "ok"}}
	serializedResponse, err := json.Marshal(response)
	if err != nil {
		s.logger.Printf("Error marshalling response: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(serializedResponse)
}
