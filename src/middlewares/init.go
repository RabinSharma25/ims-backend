package middlewares

import "github.com/gin-gonic/gin"

func ApplyMiddlewares(engine *gin.Engine) {
	engine.Use(CORSMiddleware())
}
