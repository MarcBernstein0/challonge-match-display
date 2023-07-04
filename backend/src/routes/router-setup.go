package routes

import (
	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic/main-logic"
	"github.com/gin-gonic/gin"
)

func RouteSetup(fetchData mainlogic.FetchData) *gin.Engine {
	r := gin.Default()

	r.Use(Middleware())
	superGroup := r.Group("/api")
	v1 := superGroup.Group("/v1")
	{
		v1.GET("/health", HealthCheck)
		v1.GET("/matches", MatchesGET(fetchData))
	}

	return r
}
