package providers

type ProviderErr string

func (se ProviderErr) Error() string {
	return string(se)
}

const (
	InitialInfoTimeout ProviderErr = "timeout in getting initial info"
)
