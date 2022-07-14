package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Read request data and covert to map data.
func RequestJson(r *http.Request) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		ThrowException(BAD_REQUEST)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		ThrowException(BAD_JSON)
	}
	return data
}

// Response template
type JsonResponse struct {
	Status string
	Data   interface{}
}

// Use this function to send a response to client
func (resp *JsonResponse) Response(w http.ResponseWriter, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if resp.Data != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": resp.Status, "data": resp.Data})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": resp.Status})
	}
}
