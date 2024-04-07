package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LogActivity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info().Msgf("Request METHOD [%s]  PATH [%s] STATUS [%d] IP [%s]", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), ctx.ClientIP())
		ctx.Next()
	}
}
