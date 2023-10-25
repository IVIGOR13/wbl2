package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализовать простейший telnet-клиент.
// Примеры вызовов:
// go-telnet --timeout=10s host port
// go-telnet mysite.ru 8080
// go-telnet --timeout=3s 1.1.1.1 123
// Требования:
// 1. Программа должна подключаться к указанному хосту (ip или
// доменное имя + порт) по протоколу TCP. После подключения
// STDIN программы должен записываться в сокет, а данные
// полученные и сокета должны выводиться в STDOUT
// 2. Опционально в программу можно передать таймаут на
// подключение к серверу (через аргумент --timeout, по
// умолчанию 10s)
// 3. При нажатии Ctrl+D программа должна закрывать сокет и
// завершаться. Если сокет закрывается со стороны сервера,
// программа должна также завершаться. При подключении к
// несуществующему сервер, программа должна завершаться
// через timeout

func main() {
	timeoutSeconds := flag.Int("timeout", 60, "timeout connection")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("Enter port and host")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	timeout := time.Duration(*timeoutSeconds) * time.Second

	// connection to the server
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		time.Sleep(timeout)
		log.Fatal(err)
	}
	defer conn.Close()

	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGQUIT)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				conn.Close()
				log.Fatal("Connection is broken")
			}
			_, err = conn.Write([]byte(text))
			if err != nil {
				conn.Close()
				log.Fatal(err)
			}
		}
	}()

	go func() {
		reader := bufio.NewScanner(conn)
		for reader.Scan() {
			fmt.Println(reader.Text())
		}
	}()

	select {
	case <-sigQuit:
		conn.Close()
	}
}
