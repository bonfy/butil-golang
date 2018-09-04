package butil

import (
	"encoding/json"
	"net/http"
)

//
// Web 通用方法
//

// View Model

// RespStatus struct
type RespStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// RespObjectResult struct
type RespObjectResult struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

// Write JSON Status

// writeJsonStatus func
func writeJsonStatus(w http.ResponseWriter, status string, msg string) {
	result := RespStatus{Status: status}

	if msg != "" {
		result.Message = msg
	}

	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// WriteJsonOK func
func WriteJsonOK(w http.ResponseWriter) {
	writeJsonStatus(w, "ok", "")
}

// WriteJsonErr func
func WriteJsonErr(w http.ResponseWriter, msg string) {
	writeJsonStatus(w, "err", msg)
}

// WriteJsonObject func
// obj can be stuct or list of struct
func WriteJsonObject(w http.ResponseWriter, obj interface{}) {
	result := RespObjectResult{Status: "ok", Result: obj}
	js, err := json.Marshal(result)
	if err != nil {
		WriteJsonErr(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
