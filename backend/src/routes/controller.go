package routes

import (
	"net/http"
	"sort"
	"time"

	"github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic/cache"
	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic/main-logic"
	"github.com/MarcBernstein0/challonge-match-display/backend/src/models"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

type Date struct {
	Date time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
}

func MatchesGET(fetchData mainlogic.FetchData, cache *cache.Cache) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var date Date
		matches := make([]models.TournamentMatches, 0)
		if err := c.BindQuery(&date); err != nil {
			errResponse := models.ErrorResponse{
				Message:      "did not fill out required 'date' query field",
				ErrorMessage: err.Error(),
			}
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}

		var participants []models.TournamentParticipants
		// check if cache is empty or time limit has been exceeded
		if cache.IsCacheEmpty() || cache.ShouldUpdate() {
			// get date
			// call tournaments
			tournaments, err := fetchData.FetchTournaments(date.Date.Format("2006-01-02"))
			if err != nil {
				errResponse := models.ErrorResponse{
					Message:      "failed to get tournament data",
					ErrorMessage: err.Error(),
				}
				c.JSON(http.StatusInternalServerError, errResponse)
				return
			} else if len(tournaments) == 0 {
				// no errors but empty tournaments
				c.JSON(http.StatusOK, matches)
				return
			}
			// call participants
			participants, err = fetchData.FetchParticipants(tournaments)
			if err != nil {
				errResponse := models.ErrorResponse{
					Message:      "failed to get participant data",
					ErrorMessage: err.Error(),
				}
				c.JSON(http.StatusInternalServerError, errResponse)
				return
			}
			cache.UpdateCache(participants)
		} else {
			participants = cache.GetData()
		}

		// call matches
		matches, err := fetchData.FetchMatches(participants)
		if err != nil {
			errResponse := models.ErrorResponse{
				Message:      "failed to get match data",
				ErrorMessage: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		// return matches
		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].GameName <= matches[j].GameName
		})

		c.JSON(http.StatusOK, matches)
	}
	return fn
}
