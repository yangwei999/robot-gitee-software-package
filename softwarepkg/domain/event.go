package domain

import "encoding/json"

const PlatformGitee = "gitee"

type CodePushedEvent struct {
	PkgId        string `json:"pkg_id"`
	Platform     string `json:"platform"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (e *CodePushedEvent) Message() ([]byte, error) {
	return json.Marshal(e)
}

func NewCodePushedEvent(pkgId, link string) CodePushedEvent {
	return CodePushedEvent{
		PkgId:    pkgId,
		Platform: PlatformGitee,
		RepoLink: link,
	}
}
