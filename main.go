package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"

	m "example.com/m/models"
)

func main() {
	gameManager := NewNbaMatchManager()

	http.HandleFunc("/real-time-data", gameManager.GetHandler)
	http.HandleFunc("/start", gameManager.StartHandler)

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type MatchManager struct {
	games []*m.Game
	mu    sync.Mutex
}

func NewNbaMatchManager() *MatchManager {
	return &MatchManager{}
}

func (gm *MatchManager) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	gm.mu.Lock()
	defer gm.mu.Unlock()

	statistics := gm.calculateStatistics()
	json.NewEncoder(w).Encode(statistics)
}

func (gm *MatchManager) StartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	gm.startGames()
}

func (gm *MatchManager) startGames() {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	countOfTeams := 5
	gm.games = make([]*m.Game, 0, countOfTeams)

	for i := 1; i <= countOfTeams; i++ {
		homeTeam := gm.createTeam(fmt.Sprintf("Home Team %d", i))
		awayTeam := gm.createTeam(fmt.Sprintf("Away Team %d", i+10))
		gm.games = append(gm.games, &m.Game{HomeTeam: homeTeam, AwayTeam: awayTeam})
	}

	var wg sync.WaitGroup
	for _, game := range gm.games {
		wg.Add(1)
		go func(g *m.Game) {
			defer wg.Done()
			g.SimulateMatch()
		}(game)
	}
}

func (gm *MatchManager) createTeam(name string) *m.Team {
	numberOfPlayers := 5
	team := &m.Team{Name: name, Players: make([]*m.Player, 0)}
	randomShotSuccessRate := rand.Intn(51) + 40 // -- Random 40 - 90 success rate
	for i := 1; i <= numberOfPlayers; i++ {
		team.Players = append(team.Players, &m.Player{Name: fmt.Sprintf("Player %s%d", name, i), ShotSuccessRate: randomShotSuccessRate})
	}
	return team
}

func (gm *MatchManager) calculateStatistics() []m.Statistic {
	var statistics = []m.Statistic{}

	for _, game := range gm.games {
		statistic := m.Statistic{}
		statistic.Teams = fmt.Sprintf("Match %s vs. %s\n", game.HomeTeam.Name, game.AwayTeam.Name)
		statistic.TopScorer = m.GetTopScorerPlayerName(gm.games).Name
		statistic.MostAssists = m.GetMostAssistsPlayerName(gm.games).Name
		statistic.ScoreCount = game.TotalScore()
		statistic.AssistCount = game.TotalAssist()
		statistic.Players = append(statistic.Players, game.AwayTeam.Players...)
		statistic.Players = append(statistic.Players, game.HomeTeam.Players...)

		statistics = append(statistics, statistic)
	}

	return statistics
}
