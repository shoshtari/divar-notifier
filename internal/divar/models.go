package divar

import "time"

type GetPostRequest struct {
	City           []int      `json:"city_ids"`
	PaginationData Pagination `json:"pagination_data"`
	SearchData     Search     `json:"search_data"`
}

type Pagination struct {
	PageType     string    `json:"@type"`
	LastPostDate time.Time `json:"last_post_date"`
}

type Search struct {
	FormData Form `json:"form_data"`
}

type Form struct {
	Data Filters `json:"data"`
}

type Filters struct {
	// "category": {
	//   "str": {
	//     "value": "residential-sell"
	//   }
	// },
	// "price": {
	//   "number_range": {
	//     "maximum": "4000000000"
	//   }
	// },
	// "size": {
	//   "number_range": {
	//     "minimum": "65"
	//   }
	// },
	// "sort": {
	//   "str": {
	//     "value": "sort_date"
	//   }
	// }
}

type DivarPost struct {
}
