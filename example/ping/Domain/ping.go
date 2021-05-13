package Domain

type Ping struct {
	Value string
}

func CreatePing() Ping {
	return Ping{"pong"}
}

type PingRepository interface {
	Get() Ping
	Create(Ping) error
}
