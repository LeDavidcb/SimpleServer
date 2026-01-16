package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (self *LoginRequest) ParseFromRequest(w *http.ResponseWriter, r *http.Request) error {

	// try to get the request parameters
	body, err := io.ReadAll(r.Body)
	if err != nil {
		switch err {
		case io.EOF:
			http.Error(*w, "Body must be provided", http.StatusBadRequest)
			return errors.New("The body was not provided")

		default:
			http.Error(*w, "Error reading the body", http.StatusInternalServerError)
			return errors.New("Error reading the body")
		}
	}

	err = json.Unmarshal(body, self)

	if err != nil || self.Email == "" || self.Password == "" {
		http.Error(*w, "Error reading the body", http.StatusInternalServerError)
		return fmt.Errorf("Error Unmarshaling body: %s", err)
	}

	return nil
}
