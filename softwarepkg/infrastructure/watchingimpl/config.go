package watchingimpl

import "time"

type Config struct {
	Org string `json:"org"`
	// unit second
	Interval int `json:"interval"`
}

func (cfg *Config) SetDefault() {

	cfg.Org = "euler-ttttt"

	if cfg.Interval <= 0 {
		cfg.Interval = 10
	}
}

func (cfg *Config) IntervalDuration() time.Duration {
	return time.Second * time.Duration(cfg.Interval)
}
