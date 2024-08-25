curl 'https://api.divar.ir/v8/postlist/w/search' \
--data-raw \
'
{
  "city_ids": [
    "1"
  ],
  "pagination_data": {
    "@type": "type.googleapis.com/post_list.PaginationData",
    "last_post_date": "2024-08-23T16:44:14.191730Z"
  },
  "search_data": {
    "form_data": {
      "data": {
        "category": {
          "str": {
            "value": "residential-sell"
          }
        },
        "price": {
          "number_range": {
            "maximum": "4000000000"
          }
        },
        "size": {
          "number_range": {
            "minimum": "65"
          }
        },
        "sort": {
          "str": {
            "value": "sort_date"
          }
        }
      }
    }
  }
}
'
