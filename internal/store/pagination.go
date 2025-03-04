package store

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fp PaginatedFeedQuery) Parser(r *http.Request) (PaginatedFeedQuery, error) {

	q := r.URL.Query()

	limit := q.Get("limit")
	if limit != "" {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return fp, err
		}
		fp.Limit = lim
	}

	offset := q.Get("offset")
	if offset != "" {
		off, err := strconv.Atoi(offset)
		if err != nil {
			return fp, err
		}
		fp.Offset = off
	}

	sort := q.Get("sort")
	if sort != "" {
		fp.Sort = sort
	}
	return fp, nil

}
