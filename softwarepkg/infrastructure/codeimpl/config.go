package codeimpl

type Config struct {
	ShellScript string      `json:"shell_script"`
	Robot       RobotConfig `json:"robot"`
}

func (c *Config) SetDefault() {
	if c.ShellScript == "" {
		c.ShellScript = "/opt/app/code.sh"
	}
}

type RobotConfig struct {
	Username string `json:"username" required:"true"`
	Token    string `json:"token"    required:"true"`
}
