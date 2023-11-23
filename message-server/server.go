package messageserver

import (
	"encoding/json"

	kafka "github.com/opensourceways/kafka-lib/agent"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/app"
)

func Init(s app.PackageService, c Config) *MessageServer {
	return &MessageServer{
		cfg:     c,
		service: s,
	}
}

type MessageServer struct {
	cfg     Config
	service app.PackageService
}

func (m *MessageServer) Run() error {
	return kafka.Subscribe(m.cfg.GroupName, m.handlePushCode, []string{m.cfg.Topics.PushCode})
}

func (m *MessageServer) handlePushCode(payload []byte, header map[string]string) error {
	msg := new(msgToHandlePushCode)

	if err := json.Unmarshal(payload, msg); err != nil {
		return err
	}

	return m.service.HandlePushCode(msg)
}
