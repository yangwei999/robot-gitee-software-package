package codeimpl

import (
	"fmt"

	"github.com/opensourceways/server-common-lib/utils"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
)

func NewCodeImpl(cfg Config) *codeImpl {
	gitUrl := fmt.Sprintf(
		"https://%s:%s@gitee.com/%s/",
		cfg.Robot.Username,
		"82bff85208414136c0ef726f6e76d0dc",
		cfg.PkgSrcOrg,
	)

	return &codeImpl{
		gitUrl: gitUrl,
		script: cfg.ShellScript,
	}
}

type codeImpl struct {
	gitUrl string
	script string
}

func (impl *codeImpl) Push(pkg *domain.SoftwarePkg) error {
	repoUrl := fmt.Sprintf("%s%s.git", impl.gitUrl, pkg.Name)

	params := []string{
		impl.script,
		repoUrl,
		pkg.Name,
		pkg.Importer.Name,
		pkg.Importer.Email,
		pkg.Application.SourceCode.SpecURL,
		pkg.Application.SourceCode.SrcRPMURL,
	}

	out, err, _ := utils.RunCmd(params...)
	if err != nil {
		logrus.Errorf(
			"run push code shell, err=%s, params=%v",
			err.Error()+string(out), params[:len(params)-1],
		)
	}

	return err
}
