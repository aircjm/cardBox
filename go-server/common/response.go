package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  GetStatusMsg(errCode),
		"data": data,
	})
	return
}

func (g *Gin) ResponseParamError() {
	g.C.JSON(http.StatusBadRequest, gin.H{
		"code": ErrorParamsError,
		"msg":  GetStatusMsg(ErrorParamsError),
		"data": nil,
	})
}

// 入参从Json转换成Bean
func (g *Gin) ScanRequestToBean(params interface{}) error {
	paramsBody := g.C.ShouldBindJSON(params)
	if paramsBody != nil {
		g.ResponseParamError()
	}
	return paramsBody
}
