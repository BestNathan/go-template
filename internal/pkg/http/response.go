package http

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Detail  string      `json:"detail,omitempty"`
}

func (r *Response) H() gin.H {
	return gin.H{
		"code":    r.Code,
		"message": r.Message,
		"data":    r.Data,
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{Code: 0, Message: "SUCCESS", Data: data}
}

func NewFailResponse(code int, err error) *Response {
	var m, d string

	m = err.Error()
	d = m

	if e, ok := err.(ErrorWithDetail); ok {
		d = e.Detail()
	}

	return &Response{Code: code, Message: m, Detail: d}
}

type ListData struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

// data, page, size, total
func NewListData(list interface{}) *ListData {

	// ll := len(i)
	// if ll > 0 {
	// 	page = i[0]
	// }

	// if ll > 1 {
	// 	size = i[1]
	// }

	// if ll > 2 {
	// 	total = i[2]
	// }

	return &ListData{List: list}
}

func (l *ListData) Pagination(p, s int) *ListData {
	l.Page = p
	l.Size = s
	return l
}

func (l *ListData) SetTotal(t int) *ListData {
	l.Total = t
	return l
}

func Success(ctx *gin.Context, resp *Response) {
	ctx.JSON(http.StatusOK, resp)
}

func SuccessList(ctx *gin.Context, list *ListData) {
	Success(ctx, NewSuccessResponse(list))
}

func Fail(ctx *gin.Context, err error) {
	if e, ok := err.(*Error); ok {
		ctx.JSON(http.StatusBadRequest, NewFailResponse(int(e.Code), e))
		return
	}

	if e, ok := err.(Error); ok {
		ctx.JSON(http.StatusBadRequest, NewFailResponse(int(e.Code), e))
		return
	}

	ctx.JSON(http.StatusInternalServerError, NewFailResponse(-1, err))
}
