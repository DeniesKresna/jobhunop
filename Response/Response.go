package Response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func Json(w *gin.Context, status int, payload interface{}) {
	var res ResponseData
	res.Data = payload
	res.Status = status

	w.JSON(status, res)
}
