package cache

import (
	"fmt"
	"time"

	"github.com/MarcBernstein0/challonge-match-display/backend/src/models"
)

type cacheData struct {
	tournamentsAndParticipants []models.TournamentParticipants
	timeStamp                  time.Time
}

type Cache struct {
	data                map[string]cacheData
	cacheTimer          time.Duration
	lastClearCacheTimer time.Duration
	lastClearCache      time.Time
}

func NewCache(cacheTimer, lastClearCacheTimer time.Duration) *Cache {
	return &Cache{
		data:                map[string]cacheData{},
		cacheTimer:          cacheTimer,
		lastClearCacheTimer: lastClearCacheTimer,
	}
}

func (c *Cache) UpdateCache(listTournamentParticipants []models.TournamentParticipants, date string) {
	fmt.Println("Cache is updating")
	c.data[date] = cacheData{
		tournamentsAndParticipants: listTournamentParticipants,
		timeStamp:                  time.Now(),
	}
}

func (c *Cache) GetData(date string) []models.TournamentParticipants {
	fmt.Println("Getting data from cache")
	return c.data[date].tournamentsAndParticipants
}

func (c *Cache) ShouldUpdate(date string) bool {
	if data, ok := c.data[date]; ok {
		timeSince := time.Since(data.timeStamp)
		return timeSince >= c.cacheTimer
	}
	return true

}

func (c *Cache) IsCacheEmptyDate(date string) bool {
	if data, ok := c.data[date]; ok {
		return len(data.tournamentsAndParticipants) == 0
	}
	return true
}

func (c *Cache) ShouldClearCacheData() bool {
	timeSince := time.Since(c.lastClearCache)
	return timeSince >= c.lastClearCacheTimer
}

func (c *Cache) ClearCache() {
	c.data = map[string]cacheData{}
	c.lastClearCache = time.Now()
}
