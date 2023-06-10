package routes

import (
	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/main-logic"
	"github.com/gin-gonic/gin"
)

func RouteSetup(fetchData mainlogic.FetchData) *gin.Engine {
	r := gin.Default()

	r.Use(Middleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/health", HealthCheck)
		v1.GET("/matches", MatchesGET(fetchData))
	}

	return r
}
