package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}
func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}
func (o *Observable) Notify(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Doctor struct{}

func (d *Doctor) Notify(data interface{}) {
	fmt.Printf("Doctor was notified by %s \n", data.(string))
}

type Manager struct{}

func (d *Manager) Notify(data interface{}) {
	fmt.Printf("Manager was notified by %s \n", data.(string))
}

type Worker struct {
	Obs  Observable
	Name string
}

func (w *Worker) Sick() {
	w.Obs.Notify(w.Name)
}

func NewWorker(name string) *Worker {
	return &Worker{
		Name: name,
		Obs:  Observable{list.New()},
	}
}

func main() {
	jhon := NewWorker("Jhon")

	doctor := &Doctor{}
	manager := &Manager{}

	jhon.Obs.Subscribe(doctor)
	jhon.Obs.Subscribe(manager)

	jhon.Sick()
}
