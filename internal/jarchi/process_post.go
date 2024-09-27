package jarchi

import (
	"fmt"
	"log"

	"github.com/shoshtari/divar-notifier/internal/divar"
)

func (j JarchiImp) processPost(post divar.DivarPost) {

	var msg string
	msg += post.Title + "\n"
	msg += post.Price + "\n"
	msg += fmt.Sprintf("[link](%s)", post.PostURL)

	_, err := j.notifier.SendPhoto(msg, post.ImageUrl)
	if err != nil {
		log.Println(err)
	}

}
