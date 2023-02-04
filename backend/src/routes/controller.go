package routes

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	businesslogic "github.com/MarcBernstein0/match-display/src/business-logic"
	"github.com/MarcBernstein0/match-display/src/models"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

type (
	Date struct {
		Date time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
	}
)

func MatchesGET(customClient *businesslogic.CustomClient, tournamentCache *businesslogic.Tournaments) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {
		var date Date
		matches := make([]models.TournamentMatches, 0)
		if err := ctx.BindQuery(&date); err != nil {
			errResponse := models.ErrorResponse{
				Message:      "did not fill out required 'date' query field",
				ErrorMessage: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		}

		fmt.Printf("%+v\n", date)
		err := tournamentCache.UpdateTournamentCache(date.Date.Format("2006-01-02"), customClient)
		if err != nil {
			errResponse := models.ErrorResponse{
				Message:      "failed to get tournament data",
				ErrorMessage: err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		} else if len(tournamentCache.TournamentInfo) == 0 {
			// no errors but no tournaments
			ctx.JSON(http.StatusOK, matches)
			return
		}
		fmt.Printf("tournaments %+v\n", tournamentCache)
		// get matches
		matches, err = tournamentCache.FetchMatches(customClient)
		if err != nil {
			errResponse := models.ErrorResponse{
				Message:      "failed to get match data",
				ErrorMessage: err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, errResponse)
			return
		}

		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].GameName <= matches[j].GameName
		})

		ctx.JSON(http.StatusOK, matches)

	}

	return fn
}
