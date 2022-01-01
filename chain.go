package main

import "log"

type IChainHandler interface {
	Next(Domain) error
	Join(*IChainHandler)
}

// TASK1

type Task1Handler struct {
	Handler IChainHandler
	Cfg     Task1Config
}

func (g *Task1Handler) Join(h *IChainHandler) {
	g.Handler = *h
}

func (g *Task1Handler) Next(domain Domain) error {
	if domain.Key != "Task1" {
		return g.Handler.Next(domain)
	}

	log.Println("Sucsess Task1")

	return nil
}

// TASK 2

type Task2Handler struct {
	Handler IChainHandler
	Cfg     Task2Config
}

func (g *Task2Handler) Join(h *IChainHandler) {
	g.Handler = *h
}

func (g *Task2Handler) Next(domain Domain) error {
	if domain.Key != "Task2" {
		return g.Handler.Next(domain)
	}

	log.Println("Sucsess Task2")

	return nil
}

// TASK 3

type Task3Handler struct {
	Handler IChainHandler
	Cfg     Task3Config
}

func (g *Task3Handler) Join(h *IChainHandler) {
	g.Handler = *h
}

func (g *Task3Handler) Next(domain Domain) error {
	if domain.Key != "Task3" {
		return g.Handler.Next(domain)
	}

	log.Println("Sucsess Task3")

	return nil
}
