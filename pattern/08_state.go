package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type state interface {
	TurnOn()
	TurnOff()
	getRequest(body string)
	returnResponse()
	process()
}

type server struct {
	serverOff  state
	waiting    state
	processing state

	currentState state
	body         string
}

func newServer() *server {
	server := &server{}
	serverOff := serverOff{serv: server}
	waiting := waiting{serv: server}
	processing := processing{serv: server}
	server.setCurrentState(&serverOff)
	server.waiting = &waiting
	server.processing = &processing
	server.serverOff = &serverOff

	return server
}
func (ser *server) setCurrentState(st state) {
	ser.currentState = st
}
func (ser *server) TurnOn() {
	ser.currentState.TurnOn()
}
func (ser *server) TurnOff() {
	ser.currentState.TurnOff()
}
func (ser *server) getRequest(str string) {
	ser.currentState.getRequest(str)
}
func (ser *server) returnResponse() {
	ser.currentState.returnResponse()
}
func (ser *server) process() {
	ser.currentState.process()
}

type serverOff struct {
	serv *server
}

func (serOff *serverOff) TurnOn() {
	serOff.serv.setCurrentState(serOff.serv.waiting)
	fmt.Println("Server is started")
}
func (serOff *serverOff) TurnOff() {
	fmt.Println("Zzzz")
}
func (serOff *serverOff) getRequest(str string) {
	fmt.Println("Zzzz")
}
func (serOff *serverOff) returnResponse() {
	fmt.Println("Zzzz")
}
func (serOff *serverOff) process() {
	fmt.Println("Zzzz")
}

type waiting struct {
	serv *server
}

func (ser *waiting) TurnOn() {
	fmt.Println("Server already started")
}
func (ser *waiting) TurnOff() {
	ser.serv.setCurrentState(ser.serv.serverOff)
	fmt.Println("Server is down")
}
func (ser *waiting) getRequest(str string) {
	ser.serv.setCurrentState(ser.serv.processing)
	ser.serv.body = str
	ser.serv.currentState.process()
	fmt.Println("Запрос получен и обрабатывается")
}
func (ser *waiting) returnResponse() {
	fmt.Println("Nothing to return")
}
func (ser *waiting) process() {
	fmt.Println("Nothing to process")
}

type processing struct {
	serv *server
}

func (ser *processing) TurnOn() {
	fmt.Println("Server already started")
}
func (ser *processing) TurnOff() {
	ser.serv.setCurrentState(ser.serv.serverOff)
	fmt.Println("Server is down")
}
func (ser *processing) getRequest(str string) {
	fmt.Println("Request already process")
}
func (ser *processing) returnResponse() {
	fmt.Println(ser.serv.body)
	ser.serv.setCurrentState(ser.serv.waiting)
}
func (ser *processing) process() {
	ser.serv.body = "i processed " + ser.serv.body
}

/*
func main() {
	server := newServer()
	server.getRequest("GGG")
	server.TurnOn()
	server.getRequest("GGG")
	server.getRequest("GgG")
	server.returnResponse()
	server.getRequest("GgGg")
}*/
