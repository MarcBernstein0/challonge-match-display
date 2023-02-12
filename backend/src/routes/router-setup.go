package routes

import (
	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/main-logic"
	"github.com/gin-gonic/gin"
)

func RouteSetup(fetchData mainlogic.FetchData) *gin.Engine {
	r := gin.Default()

	r.Use(Middleware())
	r.GET("/health", HealthCheck)
	r.GET("/matches", MatchesGET(fetchData))

	return r
}
