package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	stdData, err := ioutil.ReadAll(stdinReader)
	if err != nil {
		fmt.Println("couldn't read input!\n")
		os.Exit(1)
	}

	var f interface{}
	err = json.Unmarshal(stdData, &f)
	if err != nil {
		fmt.Println("couldn't unmarshal json!\n")
		os.Exit(1)
	}

	jsonMap := f.(map[string]interface{})
	for k, v := range jsonMap {
		port, err := strconv.Atoi(k)
		if err != nil {
			fmt.Println("%s not a valid port!\n", k)
			continue
		}
		response := v.(string)
		startService(port, response)
	}

	shutdownCh := make(chan os.Signal)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-shutdownCh)
}

func startService(port int, response string) {
	log.Printf("starting service at %d...\n", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("couldn't listen to port %d", port)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatalf("couldn't listen accept the request in port %d", port)
			}
			go respond(conn, response)
		}
	}()
}

func respond(conn net.Conn, response string) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("couldn't read data")
			return
		}
		readStr := string(buf[:n])
		log.Printf("READ:%s", string(readStr))
		conn.Write([]byte(response + "\n"))
	}
}
