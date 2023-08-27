package golang

import (
	"net/http"
)

func MainHttp() {
	http.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
	http.ListenAndServe(":3005", nil)
}
