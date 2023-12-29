package messageimpl

type Config struct {
	TopicsToNotify TopicsToNotify `json:"topics_to_notify"`
}

type TopicsToNotify struct {
	PushedCode string `json:"pushed_code"        required:"true"`
}
