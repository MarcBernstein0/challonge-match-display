package cache

import (
	"fmt"
	"time"

	"github.com/MarcBernstein0/challonge-match-display/backend/src/models"
)

type Cache struct {
	tournamentsAndParticipants []models.TournamentParticipants
	cacheTimer                 time.Duration
	timeStamp                  time.Time
}

func NewCache(cacheTimer time.Duration) *Cache {
	return &Cache{
		tournamentsAndParticipants: []models.TournamentParticipants{},
		cacheTimer:                 cacheTimer,
		timeStamp:                  time.Now(),
	}
}

func (c *Cache) UpdateCache(listTournamentParticipants []models.TournamentParticipants) {
	fmt.Println("Cache is updating")
	c.tournamentsAndParticipants = listTournamentParticipants
	c.timeStamp = time.Now()
}

func (c *Cache) GetData() []models.TournamentParticipants {
	fmt.Println("Getting data from cache")
	return c.tournamentsAndParticipants
}

func (c *Cache) ShouldUpdate() bool {
	timeSince := time.Since(c.timeStamp)
	return timeSince >= c.cacheTimer
}

func (c *Cache) IsCacheEmpty() bool {
	return len(c.tournamentsAndParticipants) == 0
}
