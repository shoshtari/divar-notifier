package divar

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/shoshtari/divar-notifier/internal/configs"
)

type DivarClient interface {
	GetPosts() ([]DivarPost, error)
}

type DivarClientImp struct {
	basurl string
	client http.Client
}

func (d DivarClientImp) getReqBody(maxPrice int, minSize int, lastPostDate time.Time) []byte {

	const templateFirstPage = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	const template = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData","last_post_date":"%v"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	if lastPostDate.IsZero() {
		return []byte(fmt.Sprintf(templateFirstPage, maxPrice, minSize))
	}
	return []byte(fmt.Sprintf(template, lastPostDate, maxPrice, minSize))

}

func (d DivarClientImp) GetPosts() ([]DivarPost, error) {
	var t time.Time
	reqBody := bytes.NewReader(d.getReqBody(1, 1, t))
	res, err := d.client.Post(d.basurl, "Application/Json", reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't send request")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status is not ok")
	}

	return nil, nil
}

func NewDivarClient(c configs.SectionDivarClient) DivarClient {
	return DivarClientImp{
		basurl: c.URL,
		client: http.Client{},
	}
}
