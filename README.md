# NBA Basket Match Simulator

This golang web api simulate and serve statistic data to preview from UI.

## Endpoints

POST /start -> Begin simulation, creates random teams and players.

GET /real-time-data -> While simualtion is running, get real time statistics of matches
 
## Models

```golang
type Team struct {
	Name        string
	Players     []*Player
	TotalScore  int
	TotalAssist int
	AttackCount int
}

type Player struct {
	Name                      string
	AssistCount               int
	ScoredCount               int
	ShotSuccessRate           int
	TwoPointsSuccessRateVal   float64
	ThreePointsSuccessRateVal float64
	Attempts                  []*Attempt
}

type Game struct {
	HomeTeam *Team
	AwayTeam *Team
	Lock     sync.Mutex
}

type Attempt struct {
	TryingShotPoint int
	IsSuccess       bool
}

type Statistic struct {
	Teams       string
	TopScorer   string
	MostAssists string
	AssistCount int
	ScoreCount  int
	Players     []*Player
}

```


## How is it work ?

nba.scoreboard api allows you to start the simulation and pull instant data as json.
The hierarchy is created as Game > Team > Players > Attempts.

When creating each player, the success rate of scoring a basket is randomly given between 40 and 90. Based on this data, the hit rate of the shot is calculated.

In the simulation, the home team and the opposing team play a match, their attacks start randomly, and scorers and assists are selected according to their shooting ability.

All data as TopScorer, MostAssists, AssistCount, ScoreCount and Player Statistics are returned as json.
 
