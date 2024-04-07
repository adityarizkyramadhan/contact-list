package middleware

import (
	"strings"

	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		log.Error().Msgf("Request METHOD [%s]  PATH [%s] STATUS [%d] IP [%s] ERROR [%v]", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), ctx.ClientIP(), err)
		if err != nil {
			errMsg := err.Err.Error()
			if strings.Contains(errMsg, "bad request") {
				utils.ResponseFail(ctx, 400, errMsg, err.Err)
			} else if strings.Contains(errMsg, "internal server error") {
				utils.ResponseFail(ctx, 500, errMsg, err.Err)
			} else {
				utils.ResponseFail(ctx, 500, errMsg, err.Err)
			}
			ctx.Abort()
		}
	}
}
