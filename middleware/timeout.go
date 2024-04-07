package middleware

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout() gin.HandlerFunc {
	timeLimit, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))
	if timeLimit == 0 {
		timeLimit = 15
	}
	return timeout.New(
		timeout.WithTimeout(time.Duration(timeLimit)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	utils.ResponseFail(c, 503, "Request Timeout", errors.New("request timeout"))
}
