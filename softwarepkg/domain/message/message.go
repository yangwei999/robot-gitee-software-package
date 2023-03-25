package message

type EventMessage interface {
	Message() ([]byte, error)
}

type SoftwarePkgMessage interface {
	NotifyRepoCreatedResult(EventMessage) error
	NotifyPRClosed(EventMessage) error
	NotifyPRMerged(EventMessage) error
}
