package utils


import (
	"errors"
	"encoding/json"
	"api-gateway/internal/types"
	"fmt"
	"net/http"
)



const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

func Validate(u *types.User) error {
	switch {
	case len(u.Email) == 0:
		return errors.New(ErrNameEmpty)
	case len(u.Password) == 0:
		return errors.New(ErrPasswordEmpty)
	default:
		return nil
	}
}

func Encode(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}

	return nil
}

func Decode(r *http.Request, v any) (error) {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return  fmt.Errorf("decode json: %w", err)
	}
	return  nil
}


