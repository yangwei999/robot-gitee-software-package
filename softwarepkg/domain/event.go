package domain

import "encoding/json"

const PlatformGitee = "gitee"

type PrCIFinishedEvent struct {
	PkgId        string `json:"pkg_id"`
	RelevantPR   string `json:"relevant_pr"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (e *PrCIFinishedEvent) Message() ([]byte, error) {
	return json.Marshal(e)
}

func NewPRCIFinishedEvent(
	pr *PullRequest, failedReason, repoLink string,
) PrCIFinishedEvent {
	return PrCIFinishedEvent{
		PkgId:        pr.Pkg.Id,
		RelevantPR:   pr.Link,
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

func NewRepoCreatedEvent(pr *PullRequest, url, reason string) RepoCreatedEvent {
	return RepoCreatedEvent{
		PkgId:        pr.Pkg.Id,
		Platform:     PlatformGitee,
		RepoLink:     url,
		FailedReason: reason,
	}
}

type CodePushedEvent = RepoCreatedEvent
