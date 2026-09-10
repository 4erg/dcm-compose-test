// Minimal web service for the DCM Docker Compose smoke test. It connects to the
// "redis" service over the Compose network and reports the result.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func redisPing(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("PING\r\n")); err != nil {
		return err
	}
	buf := make([]byte, 16)
	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	if string(buf[:n]) != "+PONG\r\n" {
		return fmt.Errorf("unexpected reply %q", string(buf[:n]))
	}
	return nil
}

func main() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "redis:6379"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := redisPing(addr); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "Hello from DCM — redis unreachable: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Hello from DCM — redis PONG ok")
	})
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
