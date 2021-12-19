package model

type Event struct {
	Id         int
	EventName  string
	EventDate  string
	Descripton string
}

type Respons struct {
	Result []Event `json:"result"`
}

type ErrRespons struct {
	Err error `json:"error"`
}
