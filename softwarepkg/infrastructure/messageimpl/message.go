package messageimpl

import (
	kafka "github.com/opensourceways/kafka-lib/agent"

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

func (m *MessageImpl) NotifyCodePushedResult(e message.EventMessage) error {
	return send(m.cfg.TopicsToNotify.PushedCode, e)
}

func send(topic string, v message.EventMessage) error {
	body, err := v.Message()
	if err != nil {
		return err
	}

	return kafka.Publish(topic, nil, body)
}
