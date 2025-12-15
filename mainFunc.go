package main

import (
	"log"
	"net/http"
	"time"
)

func MakeRequest(client *http.Client, url, cookie string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Ошибка создания запроса на сервер %v:", err)
		return
	} else {
		log.Print("Запрос создан!")
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка выполнения %v", err)
		return
	}

	defer resp.Body.Close()
	log.Printf("Статус запроса: %s", resp.Status)
}

func MakeRequestStatus(client *http.Client, url, cookie string) bool {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	} else {
		return false
	}

}

func IsRealUrl(url string) bool {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true // подключение есть
	} else {
		return false // недействительная ссылка
	}
}
