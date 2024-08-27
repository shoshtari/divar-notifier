package jarchi

import (
	"fmt"
	"log"

	"github.com/shoshtari/divar-notifier/internal/divar"
)

func (j JarchiImp) processPost(post divar.DivarPost) {

	msg := fmt.Sprintf(`
	Title: %s
	Image: %s
	Price: %s
	Post: %s
	`, post.Title, post.ImageUrl, post.Price, post.PostURL)
	_, err := j.notifier.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}

}
