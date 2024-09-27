package jarchi

import (
	"fmt"
	"log"

	"github.com/shoshtari/divar-notifier/internal/divar"
)

func (j JarchiImp) processPost(post divar.Post) {

	var msg string
	msg += post.Title + "\n"
	msg += post.Price + "\n"
	msg += fmt.Sprintf("[link](%s)", post.PostURL)

	_, err := j.notifier.SendPhoto(msg, post.ImageURL)
	if err != nil {
		log.Println(err)
	}

}
