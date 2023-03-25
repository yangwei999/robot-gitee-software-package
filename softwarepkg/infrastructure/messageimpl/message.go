package messageimpl

import (
	"github.com/opensourceways/robot-gitee-software-package/kafka"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/message"
)

func NewMessageImpl(c Config) *MessageImpl {
	return &MessageImpl{
		cfg: c,
	}
}

type MessageImpl struct {
	cfg Config
}

func (m *MessageImpl) NotifyRepoCreatedResult(e message.EventMessage) error {
	return send(m.cfg.TopicsToNotify.CreatedRepo, e)
}

func (m *MessageImpl) NotifyPRClosed(e message.EventMessage) error {
	return send(m.cfg.TopicsToNotify.ClosedPR, e)
}

func (m *MessageImpl) NotifyPRMerged(e message.EventMessage) error {
	return send(m.cfg.TopicsToNotify.MergedPR, e)
}

func send(topic string, v message.EventMessage) error {
	body, err := v.Message()
	if err != nil {
		return err
	}

	return kafka.Instance().Publish(topic, body)
}
