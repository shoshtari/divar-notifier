package divar

import "time"

type DivarPost struct {
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
	// action.payload.token is neccessary to obtain original link
	Action struct {
		Payload struct {
			Token string `json:"token"`
		} `json:"payload"`
	} `json:"action"`

	Price   string `json:"middle_description_text"`
	PostURL string // we must fill this part by ourselves
}

type Widget struct {
	Post DivarPost `json:"data"`
}

type DivarResponse struct {
	ListWidgets []Widget `json:"list_widgets"`
	Pagination  struct {
		HasNext bool `json:"has_next_page"`
		Data    struct {
			LastDate time.Time `json:"last_post_date"`
			Page     int       `json:"page"`
		} `json:"data"`
	} `json:"pagination"`
}
