package main

import (
	"fmt"
)

type AbsenceRequestType int

const (
	Vacation AbsenceRequestType = iota
	Sick
)

type AbsenceRequest interface {
	Notify()
}

type vacationAbsenceRequest struct {
	user        string
}

func (v vacationAbsenceRequest) Notify() {
	fmt.Println("send email to hr manager and team lead")
}

type sickAbsenceRequest struct {
	user        string
}

func (s sickAbsenceRequest) Notify() {
	fmt.Println("send email to team lead")
}

func newVacationAbsenceRequest(user string) AbsenceRequest {
	return &vacationAbsenceRequest{
		user: user,
	}
}

func newSickAbsenceRequest(user string) AbsenceRequest {
	return &sickAbsenceRequest{
		user: user,
	}
}

func newRequest(user string, requestType AbsenceRequestType) (AbsenceRequest, error) {
	if requestType == Vacation {
		return newVacationAbsenceRequest(user), nil
	}
	if requestType == Sick {
		return newSickAbsenceRequest(user), nil
	}

	return nil, fmt.Errorf("Invalid request type.")
}

func main() {
	v, _ := newRequest("Marin", Vacation)
	v.Notify()

	s, _ := newRequest("Cris", Sick)
	s.Notify()
}
