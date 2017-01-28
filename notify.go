package main

import (
	"runtime"
	"fmt"
	"os/exec"
	// "github.com/gorilla/websocket"
	// "flag"
	// "log"
	// "net/url"
	// "os"
	// "os/signal"
	// "time"
)

func notify(text string, title string, sound string, icon string) {
	switch runtime.GOOS {
	case "darwin":
		notification := fmt.Sprintf("display notification \"%s\" with title \"%s\" sound name \"%s\"", text, title, sound)
		cmd := exec.Command("osascript", "-e", notification)
		cmd.Start()
	case "linux":
		cmd := exec.Command("notify-send", "-i", icon, title, text)
		cmd.Start()
	case "windows":
		fmt.Printf("Soon\n")
	}
	
}

//var addr = flag.String("addr", "localhost:5000", "http service address")

func main() {
	fmt.Printf("Start\n")
	//notify("Alexandre send you a message !", "Message", "Glass", "")


	/*flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// err = c.WriteMessage(websocket.TextMessage, []byte("test"))
	// if err != nil {
	// 	log.Println("write:", err)
	// 	return
	// }

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		default <-done:
			err := c.WriteMessage(websocket.TextMessage, []byte("salut"))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			}
			c.Close()
			return
		}
	}*/
}