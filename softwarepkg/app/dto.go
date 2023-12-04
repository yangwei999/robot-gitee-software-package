package app

import "github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"

type CmdToHandlePushCode struct {
	Importer string `json:"importer"`
	PkgId    string `json:"pkg_id"`
	PkgName  string `json:"pkg_name"`
	Platform string `json:"platform"`
}

func (c *CmdToHandlePushCode) toPushCode(email string) domain.PushCode {
	return domain.PushCode{
		PkgId:    c.PkgId,
		PkgName:  c.PkgName,
		Platform: c.Platform,
		Importer: domain.Importer{
			Name:  c.Importer,
			Email: email,
		},
	}
}
