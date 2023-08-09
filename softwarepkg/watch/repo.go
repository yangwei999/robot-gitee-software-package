package watch

import (
	"time"

	sdk "github.com/opensourceways/go-gitee/gitee"
	"github.com/opensourceways/robot-gitee-lib/client"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/app"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/repository"
)

func NewWatchingImpl(
	cfg Config,
	repo repository.SoftwarePkg,
	service app.PackageService,
) *WatchingImpl {
	cli := client.NewClient(func() []byte {
		return []byte(cfg.RobotToken)
	})

	return &WatchingImpl{
		cfg:     cfg,
		cli:     cli,
		repo:    repo,
		service: service,
		stop:    make(chan struct{}),
		stopped: make(chan struct{}),
	}
}

type iClient interface {
	GetRepo(org, repo string) (sdk.Project, error)
	GetGiteePullRequest(org, repo string, number int32) (sdk.PullRequest, error)
}

type WatchingImpl struct {
	cfg     Config
	cli     iClient
	repo    repository.SoftwarePkg
	service app.PackageService
	stop    chan struct{}
	stopped chan struct{}
}

func (impl *WatchingImpl) Start() {
	go impl.watch()
}

func (impl *WatchingImpl) Stop() {
	close(impl.stop)

	<-impl.stopped
}

func (impl *WatchingImpl) watch() {
	interval := impl.cfg.IntervalDuration()

	checkStop := func() bool {
		select {
		case <-impl.stop:
			return true
		default:
			return false
		}
	}

	for {
		prs, err := impl.repo.FindAll()
		if err != nil {
			logrus.Errorf("find all storage pr failed, err: %s", err.Error())
		}

		for _, pr := range prs {
			impl.handle(pr)

			if checkStop() {
				close(impl.stopped)

				return
			}
		}

		time.Sleep(interval)
	}
}

func (impl *WatchingImpl) handle(pkg domain.SoftwarePkg) {
	switch pkg.Status {
	case domain.PkgStatusNew:
		if err := impl.service.HandleCreatePR(&pkg); err != nil {
			logrus.Errorf("handle retry pkg err: %s", err.Error())
		}

	case domain.PkgStatusInitialized:
		pr, err := impl.cli.GetGiteePullRequest(impl.cfg.CommunityOrg,
			impl.cfg.CommunityRepo, int32(pkg.PullRequest.Num))
		if err != nil {
			logrus.Errorf("get pr %d err: %s", pkg.PullRequest.Num, err.Error())

			return
		}

		if pr.State == sdk.StatusOpen {
			if err = impl.handleCILabel(pkg, pr); err != nil {
				logrus.Error(err.Error())
			}

			return
		}

		if err = impl.handlePRState(pr); err != nil {
			logrus.Error(err.Error())
		}

		return

	case domain.PkgStatusPRMerged:
		v, err := impl.cli.GetRepo(impl.cfg.PkgOrg, pkg.Name)
		if err != nil {
			return
		}

		if err = impl.service.HandleRepoCreated(&pkg, v.HtmlUrl); err != nil {
			logrus.Errorf("handle repo created err: %s", err.Error())
		}

	case domain.PkgStatusRepoCreated:
		if err := impl.service.HandlePushCode(&pkg); err != nil {
			logrus.Errorf("handle push code err: %s", err.Error())
		}
	}
}

func (impl *WatchingImpl) handleCILabel(pkg domain.SoftwarePkg, pr sdk.PullRequest) error {
	cmd := app.CmdToHandleCI{
		PRNum: int(pr.Number),
	}

	for _, l := range pr.Labels {
		if l.Name == impl.cfg.CISuccessLabel {
			return impl.service.HandleCI(&cmd)
		}

		if l.Name == impl.cfg.CIFailureLabel {
			cmd.FailedReason = "ci check failed"

			if v, err := impl.cli.GetRepo(impl.cfg.PkgOrg, pkg.Name); err == nil {
				cmd.RepoLink = v.HtmlUrl
				cmd.FailedReason = "package already exists"
			}

			return impl.service.HandleCI(&cmd)
		}
	}

	return nil
}

func (impl *WatchingImpl) handlePRState(pr sdk.PullRequest) error {
	switch pr.State {
	case sdk.StatusMerged:
		cmd := app.CmdToHandlePRMerged{
			PRNum: int(pr.Number),
		}

		return impl.service.HandlePRMerged(&cmd)

	case sdk.StatusClosed:
		cmd := app.CmdToHandlePRClosed{
			PRNum:      int(pr.Number),
			RejectedBy: "maintainer",
		}

		return impl.service.HandlePRClosed(&cmd)
	}

	return nil
}
