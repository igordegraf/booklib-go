package utils

import (
	"encoding/json"
	"net/http"
)

// ответ вида {"id":1}
func CreateEntityResultMessage(id uint) map[string]interface{} {
	return map[string]interface{}{"id": id}
}

func ErrorResultMessage(message string) map[string]interface{} {
	return map[string]interface{}{"error": message}
}

// func Message(status bool, message string) (map[string]interface{}) {
//   return map[string]interface{} {"status" : status, "message" : message}
// }

func JsonResponse(w http.ResponseWriter, data interface{}, httpCode int) {

	if httpCode == 0 {
		httpCode = http.StatusOK
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func UnderConstructionResponse(w http.ResponseWriter, r *http.Request) {

	JsonResponse(w, "функция пока недоступна", http.StatusForbidden)

}

func UnknownApiCallResponse(w http.ResponseWriter, r *http.Request) {

	JsonResponse(w, "unknown function", http.StatusBadRequest)

}
