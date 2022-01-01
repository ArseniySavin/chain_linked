package main

type Domain struct {
	Key string
	Value string
}

type Config struct {
	TaskCfg1 Task1Config
	TaskCfg2 Task2Config
	TaskCfg3 Task3Config
}

type Task1Config struct {
	Key string
}

type Task2Config struct {
	Key string
}

type Task3Config struct {
	Key string
}