package divar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/internal/notify"
	"github.com/shoshtari/divar-notifier/pkg"
)

var not notify.Notifier

type DivarClient interface {
	GetPosts(context.Context, chan<- DivarPost) error
}

type DivarClientImp struct {
	config configs.SectionDivarClient
	client http.Client
}

func (d DivarClientImp) getReqBody(maxPrice int, minSize int, lastPostDate time.Time) []byte {

	const templateFirstPage = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	const template = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData","last_post_date":"%v"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	if lastPostDate.IsZero() {
		ans := (fmt.Sprintf(templateFirstPage, maxPrice, minSize))
		return []byte(ans)
	}
	return []byte(fmt.Sprintf(template, lastPostDate, maxPrice, minSize))

}

func (d DivarClientImp) getPage(ctx context.Context, lastTime time.Time) (*DivarResponse, error) {
	reqBody := bytes.NewReader(d.getReqBody(d.config.MaxPrice, d.config.MinSize, lastTime))

	req, err := http.NewRequest(http.MethodPost, d.config.URL, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't make request")
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "Application/Json")

	res, err := d.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't send request")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status is %v instead of 200", res.StatusCode)
	}

	var divarRes DivarResponse
	err = json.NewDecoder(res.Body).Decode(&divarRes)

	return &divarRes, nil
}

func (d DivarClientImp) GetPosts(ctx context.Context, postChan chan<- DivarPost) error {
	var lastTime time.Time
	running := true

	process := func() chan error {
		errChan := make(chan error, 1)

		go func() {
			res, err := d.getPage(ctx, lastTime)

			if err != nil {
				errChan <- errors.Wrap(err, "couldn't get requests from page")
				return
			}

			for _, widget := range res.ListWidgets {
				postChan <- widget.Post
			}

			lastTime = res.Pagination.Data.LastDate
			running = res.Pagination.HasNext
			errChan <- nil
		}()
		return errChan
	}
	for running {
		select {
		case <-ctx.Done():
			return pkg.ErrCanceled
		case err := <-process():
			not.SendMessage(fmt.Sprint(err, "SALAM"))
			if err == nil {
				continue
			}
			return err
		}
	}

	close(postChan)
	return nil
}

func NewDivarClient(c configs.SectionDivarClient) DivarClient {
	return DivarClientImp{
		config: c,
		client: http.Client{},
	}
}
