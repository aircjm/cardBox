package common

import (
	"github.com/aircjm/gocard/common/responseStatus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  responseStatus.GetStatusMsg(errCode),
		"data": data,
	})
	return
}

func (g *Gin) ResponseParamError() {
	g.C.JSON(http.StatusBadRequest, gin.H{
		"code": responseStatus.ERROR_PARAMS_ERROR,
		"msg":  responseStatus.GetStatusMsg(responseStatus.ERROR_PARAMS_ERROR),
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
