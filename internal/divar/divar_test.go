package divar

import (
	"context"
	"testing"
	"time"

	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/internal/notify"
	"github.com/shoshtari/divar-notifier/test"
	"github.com/stretchr/testify/assert"
)

var conf configs.ServiceConfig
var divarClient DivarClient

func TestMain(m *testing.M) {
	var err error
	conf, err = test.GetTestConfig()
	if err != nil {
		panic(err)
	}
	divarClient = NewDivarClient(conf.DivarClient)
	m.Run()
}
func TestDivar(t *testing.T) {
	posts := make(chan DivarPost, 100)

	var err error
	var postCount int

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	not = notify.NewNotifier(conf.Notifier)
	go func() {
		for {
			_, exists := <-posts
			if !exists {
				return
			}
			postCount++
		}
	}()
	err = divarClient.GetPosts(ctx, posts)
	assert.Nil(t, err)
	assert.NotZero(t, postCount)
}
