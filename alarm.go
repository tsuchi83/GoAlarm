package main

import (
	"fmt"
	"os"
	"os/exec"
)

// アラーム音鳴らす
func start() {
	//out, err := exec.Command("sh", "-c", "mpg321 -l 0 SIREN.mp3 > /dev/null &").Output()
	out, err := exec.Command("play", "-t", "mp3", "ALARM.mp3").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func main() {
	start()
}
