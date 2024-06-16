package ch5

import (
	"fmt"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	timeout := time.Minute * 1
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server doesnot respond %s, %d", url, timeout)
}
