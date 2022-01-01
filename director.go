package main

type HandlerDirector struct {
	node *IChainHandler
}

// NewHandlerDirector - the director is using join for then linked stack sequence
func NewHandlerDirector(cfg Config) *HandlerDirector {
	return new(HandlerDirector).Join(&Task3Handler{Cfg: cfg.TaskCfg3}).
		Join(&Task2Handler{Cfg: cfg.TaskCfg2}).
		Join(&Task1Handler{Cfg: cfg.TaskCfg1})
}

func (r *HandlerDirector) Join(node IChainHandler) *HandlerDirector {
	if r.node == nil {
		r.node = &node
		return r
	}
	node.Join(r.node)
	r.node = &node
	return r
}

func (r *HandlerDirector) Start(domain Domain) error {
	n := *r.node
	return n.Next(domain)
}
