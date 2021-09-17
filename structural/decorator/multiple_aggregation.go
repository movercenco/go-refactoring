package main

import "fmt"

type Agger interface {
	Age() int
	SetAge(int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (b *Lizard) Age() int {
	return b.age
}

func (b *Lizard) SetAge(age int) {
	b.age = age
}

func (b *Lizard) Crawl() {
	if b.age < 10 {
		fmt.Println("Crawling")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (b *Dragon) Age() int {
	return b.bird.age
}

func (b *Dragon) SetAge(age int) {
	b.bird.SetAge(age)
	b.lizard.SetAge(age)
}

func (b *Dragon) Fly() {
	b.bird.Fly()
}

func (b *Dragon) Crawl() {
	b.lizard.Crawl()
}

func main() {
	dragon := Dragon{}
	dragon.SetAge(10)
	dragon.Fly()
	dragon.Crawl()
	fmt.Println(dragon.Age())
}
