package book

import "testing"

func Test_parseRating(t *testing.T) {
	cases := []struct {
		i string
		e float64
	}{
		{"rating0-t", 0.0},
		{"rating1-t", 1.0},
		{"rating2-t", 2.0},
		{"rating3-t", 3.0},
		{"rating4-t", 4.0},
		{"rating5-t", 5.0},
		{"", 0.0},
		{"abc", 0.0},
		{"42", 4.0},
	}

	for _, c := range cases {
		if r := parseRating(c.i); r != c.e {
			t.Errorf("parse %s failed %.2f, expected %.2f", c.i, r, c.e)
		}
	}
}
