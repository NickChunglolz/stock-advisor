package client

type Clients struct {
	TwseApi *TwseApiClient
}

func InitClients() *Clients {
	return &Clients{
		TwseApi: NewTwseApiClient(),
	}
}
