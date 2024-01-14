package models

import (
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	HomeTeam *Team
	AwayTeam *Team
	Lock     sync.Mutex
}

func (g *Game) TotalScore() int {
	return g.HomeTeam.TotalScore + g.AwayTeam.TotalScore
}

func (g *Game) TotalAssist() int {
	return g.HomeTeam.TotalAssist + g.AwayTeam.TotalAssist
}

func (g *Game) AttackCount() int {
	return g.HomeTeam.AttackCount + g.AwayTeam.AttackCount
}

func GetTopScorerPlayerName(games []*Game) *Player {
	var maxScoredPlayer *Player
	maxScored := 0

	for _, game := range games {
		for _, team := range []*Team{game.HomeTeam, game.AwayTeam} {
			for _, player := range team.Players {
				if player.ScoredCount > maxScored {
					maxScored = player.ScoredCount
					maxScoredPlayer = player
				}
			}
		}
	}

	return maxScoredPlayer
}

func GetMostAssistsPlayerName(games []*Game) *Player {
	var maxAssistPlayer *Player
	maxAssists := 0

	for _, game := range games {
		for _, team := range []*Team{game.HomeTeam, game.AwayTeam} {
			for _, player := range team.Players {
				if player.AssistCount > maxAssists {
					maxAssists = player.AssistCount
					maxAssistPlayer = player
				}
			}
		}
	}

	return maxAssistPlayer
}

func (g *Game) SimulateMatch() {
	var matchTimer = 240
	for i := 0; i < matchTimer; i++ {
		time.Sleep(1 * time.Second)
		g.Lock.Lock()

		attackerTeam := g.HomeTeam
		if rand.Intn(2) == 0 {
			attackerTeam = g.AwayTeam
		}
		attackerTeam.AttackCount++
		attackerTeam.TryShootingBasket()

		g.Lock.Unlock()
	}
}
