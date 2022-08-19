package ftx

type ftxConfig struct {
	apiKey    string
	secretKey string
}

func NewFtxConfig(apiKey, secretKey string) *ftxConfig {
	return &ftxConfig{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}
