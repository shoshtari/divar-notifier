package jarchi

import (
	"fmt"
	"log"

	"github.com/shoshtari/divar-notifier/internal/divar"
)

func (j JarchiImp) processPost(post divar.DivarPost) {

	msg := fmt.Sprintf(`
	Title: %s
	Price: %s
	Post: %s
	`, post.Title, post.Price, post.PostURL)
	_, err := j.notifier.SendPhoto(msg, post.ImageUrl)
	if err != nil {
		log.Println(err)
	}

}
