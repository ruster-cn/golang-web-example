package ping

type PingService struct{}

func NewPingService() *PingService {
	return &PingService{}
}

func (ping *PingService) Ping() (string, error) {
	return "ok", nil
}
