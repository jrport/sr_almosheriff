package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, s any) error{
	return json.NewEncoder(w).Encode(s)
}

