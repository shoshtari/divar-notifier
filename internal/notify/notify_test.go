package notify

import (
	"testing"

	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/test"
	"github.com/stretchr/testify/assert"
)

var conf configs.JarchiConfig
var notifier Notifier

func TestMain(m *testing.M) {
	conf, err := test.GetTestConfig()
	if err != nil {
		panic(err)
	}
	notifier = NewNotifier(conf.Notifier)
	m.Run()

}
func TestMessage(t *testing.T) {
	messageID, err := notifier.SendMessage("Salam")

	assert.Nil(t, err)
	assert.Greater(t, messageID, 0)

	err = notifier.EditMessage(messageID, "bye")
	assert.Nil(t, err)
}
