package codeimpl

type Config struct {
	ShellScript string `json:"shell_script"`
}

func (c *Config) SetDefault() {
	if c.ShellScript == "" {
		c.ShellScript = "/opt/app/code.sh"
	}
}
