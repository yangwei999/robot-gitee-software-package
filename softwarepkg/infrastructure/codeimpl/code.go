package codeimpl

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/infrastructure/pullrequestimpl"
)

func NewCodeImpl(r pullrequestimpl.RobotConfig, c Config, org string) *codeImpl {
	return &codeImpl{
		robot:     r,
		cfg:       c,
		pkgSrcOrg: org,
	}
}

type codeImpl struct {
	robot     pullrequestimpl.RobotConfig
	cfg       Config
	pkgSrcOrg string
}

func (impl *codeImpl) Push(pr *domain.PullRequest) error {
	repoUrl := fmt.Sprintf(
		"https://%s:%s@gitee.com/%s/%s.git",
		impl.robot.Username,
		impl.robot.Token,
		impl.pkgSrcOrg,
		pr.Pkg.Name,
	)

	cmd := exec.Command(impl.cfg.ShellScript, repoUrl, pr.Pkg.Name,
		pr.ImporterName, pr.ImporterEmail, pr.SrcCode.SpecURL, pr.SrcCode.SrcRPMURL,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(output))
	}

	return nil
}
