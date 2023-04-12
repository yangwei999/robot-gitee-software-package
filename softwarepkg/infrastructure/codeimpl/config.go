package codeimpl

type Config struct {
	ShellScript string      `json:"shell_script"`
	PkgSrcOrg   string      `json:"pkg_src_org"`
	Robot       RobotConfig `json:"robot"`
}

func (c *Config) SetDefault() {
	if c.ShellScript == "" {
		c.ShellScript = "/opt/app/code.sh"
	}

	c.PkgSrcOrg = "euler-ttttt"
	c.Robot.Username = "weiwei123444"

}

type RobotConfig struct {
	Username string `json:"username" required:"true"`
	Token    string `json:"token"    required:"true"`
}
