package routes

import (
	businesslogic "github.com/MarcBernstein0/match-display/src/business-logic"
	"github.com/gin-gonic/gin"
)

func RouteSetup(client *businesslogic.CustomClient, tournamentCache *businesslogic.Tournaments) *gin.Engine {
	r := gin.Default()

	r.Use(Middleware())
	r.GET("/health", HealthCheck)

	return r
}

// import (
// 	businesslogic "github.com/MarcBernstein0/match-display/src/business-logic"
// 	"github.com/gin-gonic/gin"
// )

// func RouteSetup(fetchData businesslogic.FetchData) *gin.Engine {
// 	r := gin.Default()

// 	r.Use(Middleware())
// 	r.GET("/health", HealthCheck)
// 	// r.GET("/matches", MatchesGET(fetchData))

// 	return r
// }
