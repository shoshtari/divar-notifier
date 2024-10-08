package divar

//TODO: get post pic as a io reader or byte array to be able to send it to notifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shoshtari/divar-notifier/internal/configs"
	"github.com/shoshtari/divar-notifier/pkg"
)

type DivarClient interface {
	GetPosts(context.Context, chan<- Post) error
}

type DivarClientImp struct {
	config configs.SectionDivarClient
	client http.Client
}

func (d DivarClientImp) getReqBody(maxPrice int, minSize int, lastPostDate time.Time) []byte {

	const templateFirstPage = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	const template = `{"city_ids":["1"],"pagination_data":{"@type":"type.googleapis.com/post_list.PaginationData","last_post_date":"%s"},"search_data":{"form_data":{"data":{"category":{"str":{"value":"residential-sell"}},"price":{"number_range":{"maximum":"%d"}},"size":{"number_range":{"minimum":"%d"}},"sort":{"str":{"value":"sort_date"}}}}}}`
	if lastPostDate.IsZero() {
		ans := (fmt.Sprintf(templateFirstPage, maxPrice, minSize))
		return []byte(ans)
	}
	lastPostDateStr := lastPostDate.Format("2006-01-02T15:04:05.999999Z")
	return []byte(fmt.Sprintf(template, lastPostDateStr, maxPrice, minSize))

}

func (d DivarClientImp) getListPage(ctx context.Context, lastTime time.Time) (*DivarResponse, error) {
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

func (d DivarClientImp) getPostPage(post DivarPost) (Post, error) {
	ans := Post{
		Title:    post.Title,
		Price:    post.Price,
		ImageURL: post.ImageUrl,
		PostURL:  post.PostURL,
	}
	return ans, nil
}

func (d DivarClientImp) getPostUrl(post DivarPost) string {
	title := strings.Replace(post.Title, " ", "-", -1)
	return fmt.Sprintf("https://divar.ir/v/%s/%s", title, post.Action.Payload.Token)
}

func (d DivarClientImp) GetPosts(ctx context.Context, postChan chan<- Post) error {
	var lastTime time.Time
	running := true
	maxTime := time.Now().Add(-1 * d.config.MaxDate)

	process := func() chan error {
		errChan := make(chan error, 1)

		go func() {
			res, err := d.getListPage(ctx, lastTime)

			if err != nil {
				errChan <- errors.Wrap(err, "couldn't get requests from page")
				return
			}

			for _, widget := range res.ListWidgets {
				widget.Post.PostURL = d.getPostUrl(widget.Post)

				post, err := d.getPostPage(widget.Post)
				if err != nil {
					errChan <- errors.Wrap(err, "couldn't get post page")
				}

				postChan <- post
			}

			lastTime = res.Pagination.Data.LastDate
			running = res.Pagination.HasNext
			errChan <- nil
		}()
		return errChan
	}
	for running && (maxTime.Before(lastTime) || lastTime.IsZero()) {
		select {
		case <-ctx.Done():
			return pkg.ErrCanceled
		case err := <-process():
			if err == nil {
				continue
			}
			return err
		}
	}

	return nil
}

func NewDivarClient(c configs.SectionDivarClient) DivarClient {
	return DivarClientImp{
		config: c,
		client: http.Client{},
	}
}
