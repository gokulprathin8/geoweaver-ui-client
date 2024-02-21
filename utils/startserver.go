package utils

import (
	"net/http"
	"os/exec"
	"time"
)

func StartServer() {

	_ = exec.Command("java", "-jar", "assets/geoweaver.jar")

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
		}

	}()

}
