package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func StartServer() {

	fmt.Println("Starting Server...")
	cmd := exec.Command("java", "-jar", "assets/geoweaver.jar")

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating stderr pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println("STDOUT:", scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println("STDERR:", scanner.Text())
		}
	}()

	go func() {
		for {
			resp, err := http.Get("http://localhost:8070/Geoweaver")
			if err == nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
				return
			}

			if err == nil {
				resp.Body.Close()
			}

			time.Sleep(1 * time.Second)
			fmt.Println(resp)
		}

	}()

}
