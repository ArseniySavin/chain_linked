package main

func main() {

	domain := Domain{
		Key:   "Task3",
		Value: "Hello",
	}

	chain := Task1Handler{
		Cfg: Task1Config{Key: "Task1"},
		Handler: &Task2Handler{
			Cfg: Task2Config{Key: "Task2"},
			Handler: &Task3Handler{
				Cfg: Task3Config{Key: "Task3"},
			},
		},
	}

	chain.Next(domain)

	cfg := Config{
		TaskCfg1: Task1Config{Key: "Task1"},
		TaskCfg2: Task2Config{Key: "Task2"},
		TaskCfg3: Task3Config{Key: "Task3"},
	}

	NewHandlerDirector(cfg).Start(domain)

}
