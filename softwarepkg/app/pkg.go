package app

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/useradapter"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/code"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/message"
)

type PackageService interface {
	HandlePushCode(*CmdToHandlePushCode) error
}

func NewPackageService(
	cd code.Code,
	p message.SoftwarePkgMessage,
	u useradapter.UserAdapter,
) *packageService {
	return &packageService{
		code:     cd,
		producer: p,
		user:     u,
	}
}

type packageService struct {
	producer message.SoftwarePkgMessage
	code     code.Code
	user     useradapter.UserAdapter
}

func (s *packageService) HandlePushCode(cmd *CmdToHandlePushCode) error {
	importerEmail, err := s.user.GetEmail(cmd.Importer)
	if err != nil {
		return err
	}

	if !s.code.CheckRepoCreated(cmd.PkgName) {
		return fmt.Errorf("repo %s has not been created", cmd.PkgName)
	}

	pushCode := cmd.toPushCode(importerEmail)
	repoUrl, err := s.code.Push(&pushCode)
	if err != nil {
		logrus.Errorf("pkgId %s push code err: %s", pushCode.PkgId, err.Error())

		return err
	}

	e := domain.NewCodePushedEvent(pushCode.PkgId, repoUrl)
	return s.producer.NotifyCodePushedResult(&e)
}
