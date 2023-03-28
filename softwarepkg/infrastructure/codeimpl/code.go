package codeimpl

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
)

func NewCodeImpl(c Config, org string) *codeImpl {
	return &codeImpl{
		cfg:       c,
		pkgSrcOrg: org,
	}
}

type codeImpl struct {
	cfg       Config
	pkgSrcOrg string
}

func (impl *codeImpl) Push(pr *domain.PullRequest) error {
	repoUrl := fmt.Sprintf(
		"https://%s:%s@gitee.com/%s/%s.git",
		impl.cfg.Robot.Username,
		impl.cfg.Robot.Token,
		impl.pkgSrcOrg,
		pr.Pkg.Name,
	)

	cmd := exec.Command(
		impl.cfg.ShellScript, repoUrl, pr.Pkg.Name, pr.ImporterName,
		pr.ImporterEmail, pr.SrcCode.SpecURL, pr.SrcCode.SrcRPMURL,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(output))
	}

	return nil
}
