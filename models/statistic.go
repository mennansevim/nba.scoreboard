package models

type Statistic struct {
	Teams       string
	TopScorer   string
	MostAssists string
	AssistCount int
	ScoreCount  int
	Players     []*Player
}
