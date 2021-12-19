package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	in := make(chan []string)
	go input(in)
	wg.Wait()

}
func input(ch chan []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("$ ")
	buffer := make(chan string)
	for {

		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		str := strings.Split(cmdString, " | ")
		for k, v := range str {
			arr := strings.Split(v, " ")
			if k != 0 {
				arr = append(arr, <-buffer)
			}

			go handler(ch, buffer)
			ch <- arr
		}

		s := <-buffer
		fmt.Printf(s)
		fmt.Print("$ ")

	}

}
func handler(in chan []string, out chan string) {
	for {
		arr := <-in
		switch arr[0] {
		case "cd":
			err := os.Chdir(arr[1])
			if err != nil {
				out <- err.Error()
			} else {
				out <- ""
			}

		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				out <- err.Error()
			} else {
				out <- pwd + "\n"
			}
		case "echo":
			out <- arr[1] + "\n"
		case "ps":
			cmd := exec.Command("ps")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
			out <- ""
		case "kill":
			i, _ := strconv.Atoi(arr[1])
			proc, err := os.FindProcess(i)
			proc.Kill()
			if err != nil {
				out <- err.Error()
			} else {
				out <- ""
			}
		case "nc":
			err := nc(arr[1], arr[2])
			if err != nil {
				out <- err.Error()

			} else {
				out <- "ok"
			}
		default:
			out <- arr[0] + ": command not found\n"

		}

	}
}

func output(ch chan string) {
	for {
		s := <-ch
		fmt.Printf(s)
		fmt.Print("$ ")
	}
}

func nc(host, port string) error {
	conn, err := net.Dial("tcp", host+":"+port)
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
	}()
	wg.Wait()
	return err
}
