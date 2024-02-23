package httperror

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Handle(ctx *gin.Context, err error, httpStatus int) {
	log.Println(err)
	ctx.JSON(http.StatusBadRequest, gin.H{"status": "error"})
}
