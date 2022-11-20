package store

import (
	"context"
)

type DummpyMatchPersistor struct {
	Matches []*Match
	Added   []*Match
}

func NewDummyMatchPersistor(matches []*Match) *DummpyMatchPersistor {
	return &DummpyMatchPersistor{
		Matches: matches,
	}
}

func (p *DummpyMatchPersistor) AddMatch(ctx context.Context, match *Match) error {
	p.Added = append(p.Added, match)
	p.Matches = append(p.Matches, match)
	return nil
}

func (p *DummpyMatchPersistor) GetMatches(ctx context.Context, stonk string) ([]*Match, error) {
	allMatches := make([]*Match, 0, len(p.Matches))
	for _, m := range p.Matches {
		if stonk != "" && m.Stonk != stonk {
			continue
		}
		allMatches = append(allMatches, m)
	}

	return allMatches, nil
}
