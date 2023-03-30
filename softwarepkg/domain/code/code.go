package code

import "github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"

type Code interface {
	Push(pkg *domain.SoftwarePkg) error
}
