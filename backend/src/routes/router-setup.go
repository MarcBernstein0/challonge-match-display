package routes

import (
	businesslogic "github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic"
	"github.com/gin-gonic/gin"
)

func RouteSetup(client *businesslogic.CustomClient, tournamentCache *businesslogic.Tournaments) *gin.Engine {
	r := gin.Default()

	r.Use(Middleware())
	r.GET("/health", HealthCheck)
	r.GET("/matches", MatchesGET(client, tournamentCache))

	return r
}
