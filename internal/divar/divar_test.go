package divar

import (
	"testing"

	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/test"
	"github.com/stretchr/testify/assert"
)

var conf configs.ServiceConfig
var divarClient DivarClient

func TestMain(m *testing.M) {
	conf, err := test.GetTestConfig()
	if err != nil {
		panic(err)
	}
	divarClient = NewDivarClient(conf.DivarClient)
	m.Run()
}
func TestDivar(t *testing.T) {
	posts, err := divarClient.GetPosts()
	assert.Nil(t, err)
	assert.NotNil(t, posts)
}
