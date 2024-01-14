package models

import (
	"example.com/m/shared"
	"math/rand"
	"time"
)

type Player struct {
	Name                      string
	AssistCount               int
	ScoredCount               int
	ShotSuccessRate           int
	TwoPointsSuccessRateVal   float64
	ThreePointsSuccessRateVal float64
	Attempts                  []*Attempt
}

func (p *Player) TryAShot() int {
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	tryingScorePoint := random.Intn(2) + 2                // -- Random choose between 2 Points or 3Points
	isShotSuccess := random.Intn(100) < p.ShotSuccessRate // -- SuccessRate can be changed by player
	p.Attempts = append(p.Attempts, &Attempt{IsSuccess: isShotSuccess, TryingShotPoint: tryingScorePoint})

	if isShotSuccess {
		p.ScoredCount += tryingScorePoint
	}

	p.TwoPointsSuccessRateVal = p.TwoPointsSuccessRate()
	p.ThreePointsSuccessRateVal = p.ThreePointsSuccessRate()

	if isShotSuccess {
		return tryingScorePoint
	}

	return 0
}

func (p *Player) TwoPointsSuccessRate() float64 {
	successCount := 0
	attemptCount := 0
	for _, attempt := range p.Attempts {
		if attempt.TryingShotPoint == 2 {
			attemptCount++
			if attempt.IsSuccess {
				successCount++
			}
		}
	}
	if successCount == 0 {
		return 0
	}
	return shared.RoundToTwoDecimals(float64(successCount) / float64(attemptCount) * 100)
}

func (p *Player) ThreePointsSuccessRate() float64 {
	successCount := 0
	attemptCount := 0
	for _, attempt := range p.Attempts {
		if attempt.TryingShotPoint == 3 {
			attemptCount++
			if attempt.IsSuccess {
				successCount++
			}
		}
	}

	if successCount == 0 {
		return 0
	}

	return shared.RoundToTwoDecimals(float64(successCount) / float64(attemptCount) * 100)
}
