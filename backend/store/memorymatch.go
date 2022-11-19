package store

import "go.uber.org/zap"

type MemoryMatchPersistor struct {
	l *zap.Logger
	m map[Security][]Match
}

func NewMemoryMatchPersistor(l *zap.Logger) *MemoryMatchPersistor {
	return &MemoryMatchPersistor{
		l: l,
		m: make(map[Security][]Match, 5),
	}
}

func (p *MemoryMatchPersistor) AddMatch(*Match) error {
	// TODO: Implement me
	return nil
}

// Security is optional -> if nil return all matches
func (p *MemoryMatchPersistor) GetMatches() ([]*Match, error) {
	// TODO: Implement me
	return nil, nil
}
