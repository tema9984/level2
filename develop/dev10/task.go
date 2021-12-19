package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeout int
	var host, port string
	flag.IntVar(&timeout, "timeout", 10, "set time for connections")
	flag.StringVar(&host, "h", "127.0.0.1", "set host")
	flag.StringVar(&port, "p", "9090", "set port")
	flag.Parse()
	err := telnet(host, port, timeout)
	fmt.Println(err)
}
func telnet(host, port string, timeout int) error {
	d := net.Dialer{Timeout: time.Second * time.Duration(timeout)}
	conn, err := d.Dial("tcp", host+":"+port)
	if err != nil {
		return err
	}
	tcpconn := conn.(*net.TCPConn)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(tcpconn, os.Stdin)
		fmt.Fprintln(os.Stderr, "conn 2 os.Stdin done")
		tcpconn.CloseWrite()
		wg.Done()
	}()
	go func() {
		io.Copy(os.Stdout, tcpconn)
		fmt.Fprintln(os.Stderr, "os.Stdout 2 conn done")
		tcpconn.CloseRead()
		wg.Done()
		wg.Done()
	}()
	wg.Wait()
	return err
}
