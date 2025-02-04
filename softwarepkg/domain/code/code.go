package code

import "github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"

type Code interface {
	Push(pkg *domain.PushCode) (string, error)
	CheckRepoCreated(repo string) bool
}
