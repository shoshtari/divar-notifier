package jarchi

//TODO: handle retry/log for functions like func()error
//TODO: generate proper message to give to notifier
//TODO: add a db to avoid processing a message twice

import (
	"context"
	"log"
	"time"

	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/internal/divar"
	"github.com/shoshtari/divar-notifier/internal/notify"
)

type Jarchi interface {
	Start() error
	Stop() error
}

type JarchiImp struct {
	notifier    notify.Notifier
	divarClient divar.DivarClient
	config      configs.SectionCore
}

func (j JarchiImp) processPost(post divar.DivarPost) {

	_, err := j.notifier.SendMessage(post.Title)
	if err != nil {
		log.Println(err)
	}

}

func (j JarchiImp) processPosts() error {
	postChan := make(chan divar.DivarPost)
	defer close(postChan)

	go func() {
		for {
			post, exists := <-postChan
			if !exists {
				return
			}
			j.processPost(post)
		}
	}()

	ctx := context.TODO()
	return j.divarClient.GetPosts(ctx, postChan)

}

func (j JarchiImp) Start() error {
	j.processPosts()
	t := time.NewTicker(j.config.Interval)
	for range t.C {
		j.processPosts()
	}
	return nil
}

// Stop implements Jarchi.
func (j JarchiImp) Stop() error {
	panic("unimplemented")
}

func New(divarClient divar.DivarClient, notifier notify.Notifier, config configs.SectionCore) Jarchi {
	return JarchiImp{
		config:      config,
		divarClient: divarClient,
		notifier:    notifier,
	}
}
