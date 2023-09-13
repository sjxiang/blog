package serializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/blog/pkg/errno"
)

// 序列化器


type ErrResponse struct {
    // 业务错误码.
    Code string `json:"code"`

    // 直接对外展示的错误信息.
    Message string `json:"message"`
}


func BuildResponse(c *gin.Context, err error, data interface{}) {
    if err != nil {
        hcode, code, message := errno.Decode(err)
        c.JSON(hcode, ErrResponse{
            Code:    code,
            Message: message,
        })

        return
    }

    c.JSON(http.StatusOK, data)
}



type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`  // 和 error 通用
	Data    interface{} `json:"data,omitempty"`
}

