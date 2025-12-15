package main

import (
	"os/exec"
	"runtime"
)

// Открыть браузер
func OpenBrowser(url string) {
	switch runtime.GOOS { //хранит строку с названием ОС
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start() // Start - метод, который запускает программу в отдельном процессе, не блокируя основной поток
	default:
		exec.Command("xdg-open", url).Start()
	}
}
