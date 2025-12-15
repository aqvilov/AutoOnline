package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//url
	var Url string
	// Просим ввести верный url
	for {
		fmt.Println("Введите url сайта в формате 'https://example.com/path': ")
		fmt.Scan(&Url)

		if IsRealUrl(Url) {
			break
		} else {
			fmt.Println("Вы ввели неправильный URL!")
			fmt.Println("")
		}
	}

	// Интервал между запросами

	var RequestTime int
	fmt.Println("Введите интервал между запросами (в минутах): ")
	for {
		fmt.Scan(&RequestTime)
		if RequestTime == 0 {
			log.Printf("Интервал между запросами не может быть равен 0!\nПожалуйста, введите другое значение: ")
		} else {
			break
		}
	}

	// Время работы
	var allRequestTime int
	fmt.Println("Введите сколько времени будут отправляться запросы (в минутах):")
	fmt.Println("(Введите '0' если хотите бесконечное время запросов): ")
	fmt.Scan(&allRequestTime)

	//get cookie
	fmt.Println("Введите cookie, следуя инструкции:")
	fmt.Print(`
╔═══════════════════════════════════════════════════════════════╗
║                   ИНСТРУКЦИЯ ПО ПОЛУЧЕНИЮ КУКИ                ║
╠═══════════════════════════════════════════════════════════════╣
║                                                               ║
║  ШАГ 1: ОТКРОЙТЕ ИНСТРУМЕНТЫ РАЗРАБОТЧИКА                     ║
║  • Chrome/Edge: Нажмите F12 или Ctrl+Shift+I                  ║
║  • Firefox: Нажмите F12 или Ctrl+Shift+I                      ║
║  • Safari: Нажмите Cmd+Opt+I (включите "Develop" в настройках)║
║                                                               ║
║  ШАГ 2: НАЙДИТЕ ВКЛАДКУ С COOKIES:                            ║
║  • Chrome/Edge: Вкладка "Application" → "Cookies" в левом меню║
║  • Firefox: Вкладка "Storage" → "Cookies"                     ║
║  • Safari: Вкладка "Storage" → "Cookies"                      ║
║                                                               ║
║  ШАГ 3: СКОПИРУЙТЕ ВСЕ КУКИ:                                  ║
║  • Найдите домен вашего сайта                                 ║
║  • Скопируйте ВСЕ пары "Name=Value"                           ║
║  • Формат: "name1=value1; name2=value2; name3=value3"         ║
║                                                               ║
║  ПРИМЕР:                                                      ║
║    _ga=GA1.1.123456789.1234567890;                            ║
║    sessionid=abc123def456ghi789;                              ║
║    csrftoken=xyz789abc456def123                               ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝
`)

	fmt.Scanln()

	time.Sleep(4 * time.Second)
	OpenBrowser(Url)

	reader := bufio.NewReader(os.Stdin)

	client := &http.Client{ // для тестового запроса
		Timeout: 30 * time.Second,
	}

	var cookie string

	for {
		fmt.Println("Введите cookie, следуя инструкции:")
		cookie, _ = reader.ReadString('\n')
		cookie = strings.TrimSpace(cookie)

		if cookie == "" {
			continue
		}

		//Проверка тестовым запросом
		if TestReq := MakeRequestStatus(client, Url, cookie); TestReq {
			// Тестовый запрос успешен --> куки верные
			fmt.Println("Тестовый запрос успешен! Начинаем работу программы.")
			break
		} else {
			fmt.Println("Тестовый запрос не удался. Cookie неверны, либо устарели")
			fmt.Println("Хотите ввести Cookie заново (y/n)?")

			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))

			if response == "y" || response == "yes" || response == "д" || response == "да" {
				continue // переход на следующую итерацию цикла
			} else { // пытаемся работать с исходными куки
				return
			}
		}
	}

	time.Sleep(4 * time.Second)
	OpenBrowser(Url)

	if RequestTime > allRequestTime && allRequestTime != 0 {
		log.Printf("Ваш интервал между запросами в %d минут, превышает время работы программы в %d минут", RequestTime, allRequestTime)
		return
	} else {

		timeout := time.After(time.Duration(allRequestTime) * time.Minute) // канал, который отправит сигнал через указанное время ( например через 10 минут)

		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		MakeRequest(client, Url, cookie)

		ticker := time.NewTicker(time.Duration(RequestTime) * time.Minute) // 5 minutes default
		defer ticker.Stop()

		fmt.Printf("Запускаем запросы каждые %d минут\n", RequestTime)
		fmt.Printf("Нажмите Ctrl+C для остановки\n")

		if allRequestTime == 0 { // работает бесконечно ( пока не нажмут ctrl+C)
			fmt.Println("\nВы ввели '0', запросы будут продолжаться пока вы не нажмете Ctrl+C")
			for range ticker.C {
				MakeRequest(client, Url, cookie)
			}
		} else {
			fmt.Printf("\nЗапросы будут отправляться %d минут!\n", allRequestTime)
			for {
				select {
				case <-ticker.C:
					MakeRequest(client, Url, cookie)
				case <-timeout:
					fmt.Println("Время вышло. Программа завершена")
					return
				}
			}
		}
	}
}
