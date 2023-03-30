package domain

import "encoding/json"

const PlatformGitee = "gitee"

type PRCIFinishedEvent struct {
	PkgId        string `json:"pkg_id"`
	RelevantPR   string `json:"relevant_pr"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (e *PRCIFinishedEvent) Message() ([]byte, error) {
	return json.Marshal(e)
}

func NewPRCIFinishedEvent(
	pkg *SoftwarePkg, failedReason, repoLink string,
) PRCIFinishedEvent {
	return PRCIFinishedEvent{
		PkgId:        pkg.Id,
		RelevantPR:   pkg.PullRequest.Link,
		RepoLink:     repoLink,
		FailedReason: failedReason,
	}
}

type RepoCreatedEvent struct {
	PkgId        string `json:"pkg_id"`
	Platform     string `json:"platform"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (e *RepoCreatedEvent) Message() ([]byte, error) {
	return json.Marshal(e)
}

func NewRepoCreatedEvent(pkg *SoftwarePkg, url, reason string) RepoCreatedEvent {
	return RepoCreatedEvent{
		PkgId:        pkg.Id,
		Platform:     PlatformGitee,
		RepoLink:     url,
		FailedReason: reason,
	}
}

type CodePushedEvent = RepoCreatedEvent
