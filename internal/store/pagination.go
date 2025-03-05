package store

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Search string   `json:"search" validate:"max=100"`
	Tags   []string `json:"tags" validate:"max=10"`
	Since  string   `json:"since" validate:"max=15"`
	Until  string   `json:"until" validate:"max=15"`
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

	tags := q.Get("tags")
	if tags != "" {
		fp.Tags = strings.Split(tags, ",")
	}

	search := q.Get("search")
	if search != "" {
		fp.Search = search
	}

	since := q.Get("since")
	if since != "" {
		fp.Since = ParseTime(since)
	}

	until := q.Get("until")
	if until != "" {
		fp.Until = ParseTime(until)
	}

	return fp, nil

}

func ParseTime(t string) string {
	parsedTime, err := time.Parse(time.DateTime, t)
	if err != nil {
		return "Invalid time format"
	}
	return parsedTime.Format(time.DateTime)
}
