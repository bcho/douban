package book

import (
	"fmt"
	"time"
)

type Subject struct {
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Rating      float64   `json:"rating"`
	CollectDate time.Time `json:"collect_date"`
	WishDate    time.Time `json:"wish_date"`
	DoDate      time.Time `json:"do_date"`
}

func (s Subject) String() string {
	return fmt.Sprintf("<Book:%s(%s) %.2f>", s.Title, s.Url, s.Rating)
}
