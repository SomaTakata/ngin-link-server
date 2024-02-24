package httperror

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Handle(ctx *gin.Context, err error, httpStatus int) {
	log.Println(err)
	ctx.JSON(httpStatus, gin.H{"status": "error"})
}
