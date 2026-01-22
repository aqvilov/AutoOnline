package tests

import (
	MainFuncions "FunPayAutoOnline/mainFunc"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	// создаем сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")
		userAgent := r.Header.Get("User-Agent")

		if cookie == "" || cookie == "invalid" {
			w.WriteHeader(http.StatusBadRequest) // code 400-- неверный запрос
			w.Write([]byte("неверные cookie"))
			return
		}

		if !strings.Contains(userAgent, "Mozilla") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Неверный User-Agent"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	}))
	defer server.Close()

	client := &http.Client{} // удовлетворяет нужному интерфейсу

	err := MainFuncions.MakeRequest(client, server.URL, "")

	if err == nil {
		t.Error("ожидалась ошибка при неверно введенных cookie, но ее нет")
	}
}
