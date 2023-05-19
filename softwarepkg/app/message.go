package app

import (
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/pullrequest"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/repository"
)

type MessageService interface {
	NewPkg(*CmdToHandleNewPkg) error
	RetryPkg(*domain.SoftwarePkg) error
}

func NewMessageService(repo repository.SoftwarePkg, prCli pullrequest.PullRequest,
) *messageService {
	return &messageService{
		repo:  repo,
		prCli: prCli,
	}
}

type messageService struct {
	repo  repository.SoftwarePkg
	prCli pullrequest.PullRequest
}

func (s *messageService) NewPkg(cmd *CmdToHandleNewPkg) error {
	pr, err := s.prCli.Create(cmd)
	if err != nil {
		cmd.SetPkgStatusFailed()
	} else {
		cmd.PullRequest = pr
		cmd.SetPkgStatusInitialized()
	}

	return s.repo.Add(cmd)
}

func (s *messageService) RetryPkg(pkg *domain.SoftwarePkg) error {
	pr, err := s.prCli.Create(pkg)
	if err == nil {
		pkg.PullRequest = pr
		pkg.SetPkgStatusInitialized()

		return s.repo.Save(pkg)
	}

	return nil
}
