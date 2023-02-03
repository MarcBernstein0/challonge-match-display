package routes

import (
	"fmt"
	"net/http"
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

	}

	return fn
}

// func MatchesGET(fetchData mainlogic.FetchData) gin.HandlerFunc {
// 	fn := func(c *gin.Context) {
// 		var date Date
// 		matches := make([]models.TournamentMatches, 0)
// 		if err := c.BindQuery(&date); err != nil {
// 			errResponse := models.ErrorResponse{
// 				Message:      "did not fill out required 'date' query field",
// 				ErrorMessage: err.Error(),
// 			}
// 			c.JSON(http.StatusBadRequest, errResponse)
// 			return
// 		}
// 		// fmt.Printf("%+v\n", date)
// 		// get date
// 		// call tournaments
// 		tournaments, err := fetchData.FetchTournaments(date.Date.Format("2006-01-02"))
// 		if err != nil {
// 			errResponse := models.ErrorResponse{
// 				Message:      "failed to get tournament data",
// 				ErrorMessage: err.Error(),
// 			}
// 			c.JSON(http.StatusInternalServerError, errResponse)
// 			return
// 		} else if len(tournaments) == 0 {
// 			// no errors but empty tournaments
// 			c.JSON(http.StatusOK, matches)
// 			return
// 		}
// 		// fmt.Printf("tournaments %+v\n", tournaments)
// 		// call particiapnts
// 		participants, err := fetchData.FetchParticipants(tournaments)
// 		if err != nil {
// 			errResponse := models.ErrorResponse{
// 				Message:      "failed to get participant data",
// 				ErrorMessage: err.Error(),
// 			}
// 			c.JSON(http.StatusInternalServerError, errResponse)
// 			return
// 		}
// 		// fmt.Printf("list of participants %+v\n", participants)
// 		// call matches
// 		matches, err = fetchData.FetchMatches(participants)
// 		if err != nil {
// 			errResponse := models.ErrorResponse{
// 				Message:      "failed to get match data",
// 				ErrorMessage: err.Error(),
// 			}
// 			c.JSON(http.StatusInternalServerError, errResponse)
// 			return
// 		}
// 		// fmt.Printf("list of matches %+v\n", matches)
// 		// return matches
// 		sort.SliceStable(matches, func(i, j int) bool {
// 			return matches[i].GameName <= matches[j].GameName
// 		})

// 		c.JSON(http.StatusOK, matches)
// 	}
// 	return fn
// }
