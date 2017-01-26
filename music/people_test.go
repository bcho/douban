package music

import "testing"

func Test_mergeCollectOpts_useDefault(t *testing.T) {
	c := mergeCollectOpts(nil, nil)
	if c.SortBy != defaultCollectOpt.SortBy {
		t.Errorf("%s", c.SortBy)
	}
}

func Test_mergeCollectOpts_overwrite(t *testing.T) {
	c := mergeCollectOpts(
		&CollectOpt{
			SortBy: SortByTime,
		},
		&CollectOpt{
			SortBy: SortByTitle,
		},
	)
	if c.SortBy != SortByTitle {
		t.Errorf("%s", c.SortBy)
	}
}
