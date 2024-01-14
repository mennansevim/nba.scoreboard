package models

import (
	"math/rand"
	"time"
)

type Team struct {
	Name        string
	Players     []*Player
	TotalScore  int
	TotalAssist int
	AttackCount int
}

func (t *Team) TryShootingBasket() {
	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	scoredPlayerIndex := random.Intn(len(t.Players))
	scoredPlayer := t.Players[scoredPlayerIndex]
	score := scoredPlayer.TryAShot()
	t.TotalScore += score

	if score == 0 {
		return
	}

	// -- Choose unique number except scorerPlayer
	asistedPlayerIndex := scoredPlayerIndex
	for asistedPlayerIndex == scoredPlayerIndex {
		asistedPlayerIndex = random.Intn(len(t.Players))
	}
	asistedPlayer := t.Players[asistedPlayerIndex]
	asistedPlayer.AssistCount++
	t.TotalAssist++
}
