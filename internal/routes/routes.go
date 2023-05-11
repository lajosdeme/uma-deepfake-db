package routes

import (
	"fmt"

	dbapi "github.com/deepfake-db/internal/api"
	"github.com/deepfake-db/internal/config"
	"github.com/gin-gonic/gin"
)

/*
GET
api/assertions/
- filter by imageId, assertionId, address

GET
api/notifications/
- param address
- filter for resolved & not seen

POST
api/notifications/set-seen/
- param assertionID
*/
func ConfigRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Next()
	})

	api := r.Group("/api")
	api.GET("/assertions", dbapi.GetAssertions)
	api.GET("/notifications/:asserter", dbapi.GetNotifications)
	api.POST("/notifications", dbapi.SetNotificationSeen)

	api.GET("/ping", dbapi.Ping)
	return r
}

func RunRouter(r *gin.Engine) {
	r.Run(fmt.Sprintf(":%d", config.DbConfig.Port))
}
