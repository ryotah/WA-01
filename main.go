package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println(n.Format(time.DateTime))
	openbrowser("https://replit.com/")
}

// https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
