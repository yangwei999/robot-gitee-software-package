package codeimpl

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/utils"
)

func NewCodeImpl(c Config) *codeImpl {
	return &codeImpl{
		cfg: c,
	}
}

type codeImpl struct {
	cfg Config
}

func (impl *codeImpl) Push(pr *domain.PullRequest) error {
	repoUrl := fmt.Sprintf(
		"https://%s:%s@gitee.com/%s/%s.git",
		impl.cfg.Robot.Username,
		impl.cfg.Robot.Token,
		impl.cfg.PkgSrcOrg,
		pr.Pkg.Name,
	)

	params := []string{
		impl.cfg.ShellScript,
		repoUrl,
		pr.Pkg.Name,
		pr.ImporterName,
		pr.ImporterEmail,
		pr.SrcCode.SpecURL,
		pr.SrcCode.SrcRPMURL,
	}

	_, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, params=%v",
			err.Error(), params[:len(params)-1],
		)
	}

	return err
}
