package main

func main() {

	domain := Domain{
		Key:   "Task3",
		Value: "Hello",
	}

	cfg := Config{
		TaskCfg1: Task1Config{Key: "Task1"},
		TaskCfg2: Task2Config{Key: "Task2"},
		TaskCfg3: Task3Config{Key: "Task3"},
	}

	NewHandlerDirector(cfg).Start(domain)

}
