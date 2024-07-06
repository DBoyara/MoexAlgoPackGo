package moexalgopackgo

type IAlgoClient interface {
	Authenticate() error
}

type AlgoClient struct {
	auth_url string
	login    string
	password string
	cert     string
}

func NewAlgoClient(login, password string) IAlgoClient {
	return &AlgoClient{
		auth_url: "https://passport.moex.com/authenticate",
		login:    login,
		password: password,
	}
}
