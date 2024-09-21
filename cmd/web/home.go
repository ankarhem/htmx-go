package web

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func RandomNumberHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.IntN(100)
	_, err := w.Write([]byte(strconv.Itoa(number)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
