package tools

import (
	"api/utils"
	"net/http"
)

// IPAddr - get user IP Address
func IPAddr(w http.ResponseWriter, r *http.Request) {
	// ========= CORS =========
	utils.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// ========================
	// all done
	res := struct {
		IPAddr string `json:"ipAddr"`
	}{
		IPAddr: r.RemoteAddr,
	}
	utils.SendJSON(&w, res)
	return
}
