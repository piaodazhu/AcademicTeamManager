package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Page struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func Result(code int, data interface{}, c *gin.Context) {
	message := msg[code]
	c.JSON(http.StatusOK, Response{code, message, data})
}

func PageResult(code int, data interface{}, rows int64, c *gin.Context) {
	message := msg[code]
	page := &Page{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{code, message, page})
}
