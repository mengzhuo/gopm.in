package gopm

var cfg *Config

type Config struct {
	Domain string
}

func SetDomain(domain string) {
	if cfg == nil {
		cfg = &Config{}
	}
	cfg.Domain = domain
}
