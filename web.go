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

// RespSingleResult struct
type RespSingleResult struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

// RespListResult struct
type RespListResult struct {
	Status string        `json:"status"`
	Result []interface{} `json:"result"`
}

// Write JSON Status

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

// WriteSingleObjectStatus func
func WriteSingleObjectStatus(w http.ResponseWriter, obj interface{}) {
	result := RespSingleResult{Status: "ok", Result: obj}
	js, err := json.Marshal(result)
	if err != nil {
		WriteJsonErr(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// WriteListObjectStatus func
func WriteListObjectStatus(w http.ResponseWriter, obj []interface{}) {
	result := RespListResult{Status: "ok", Result: obj}
	js, err := json.Marshal(result)
	if err != nil {
		WriteJsonErr(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
