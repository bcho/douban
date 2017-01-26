package movie

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bcho/douban/crawler"
	"github.com/bcho/douban/people"
)

const (
	basePath = "http://movie.douban.com"

	COLLECT_DATE_FMT = "2006-01-02"
)

type sortBy string

const (
	SortByTime   sortBy = "time"
	SortByRating        = "rating"
	SortByTitle         = "title"
)

type CollectOpt struct {
	SortBy sortBy // sort by method
}

var defaultCollectOpt = &CollectOpt{
	SortBy: SortByTime,
}

func mergeCollectOpts(opts ...*CollectOpt) *CollectOpt {
	c := &CollectOpt{}

	opts = append([]*CollectOpt{defaultCollectOpt}, opts...)

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.SortBy != "" {
			c.SortBy = opt.SortBy
		}
	}

	return c
}

// Collect crawls douban user's movie collection
func Collect(people people.People, opt *CollectOpt) ([]*Subject, *http.Response, error) {
	opt = mergeCollectOpts(opt)

	url, err := url.Parse(fmt.Sprintf("%s/people/%s/collect", basePath, people.Username))
	if err != nil {
		return nil, nil, err
	}
	q := url.Query()
	q.Set("start", "0")
	q.Set("filter", "all")
	q.Set("mode", "list")
	q.Set("sort", string(opt.SortBy))
	url.RawQuery = q.Encode()

	resp, err := crawler.DefaultClient.Get(url.String())
	if err != nil {
		return nil, resp, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, resp, err
	}

	var albums []*Subject

	doc.Find(".list-view .item").Each(func(_ int, s *goquery.Selection) {
		album := &Subject{}
		titleNode := s.Find(".title > a")
		url := titleNode.AttrOr("href", "")
		if url == "" {
			return
		}
		title := strings.TrimSpace(titleNode.Text())
		dateNode := s.Find(".date")
		date, err := time.Parse(COLLECT_DATE_FMT, strings.TrimSpace(dateNode.Text()))
		if err != nil {
			return
		}
		rating := parseRating(strings.TrimSpace(s.Find(".date > span").AttrOr("class", "")))

		album = &Subject{
			Url:         url,
			Title:       title,
			Rating:      rating,
			CollectDate: date,
		}

		albums = append(albums, album)
	})

	return albums, resp, nil
}
