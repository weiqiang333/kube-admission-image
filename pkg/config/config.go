package config

// FlagVar 定义传参
type FlagVar struct {
	Addrss              string
	Tls                 bool
	CertFile            string
	KeyFile             string
	SourceDefaultPolicy bool
}
