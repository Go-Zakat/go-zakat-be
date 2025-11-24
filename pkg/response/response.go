package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response adalah format standar response API
type Response struct {
	Status  string      `json:"status"`            // "success" atau "error"
	Message string      `json:"message"`           // Pesan deskriptif
	Data    interface{} `json:"data,omitempty"`    // Data payload (untuk success)
	Error   interface{} `json:"error,omitempty"`   // Detail error (untuk error)
}

// Success mengirimkan response sukses standar
func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error mengirimkan response error standar
func Error(c *gin.Context, code int, message string, errDetail interface{}) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
		Error:   errDetail,
	})
}

// BadRequest mengirimkan response 400 Bad Request
func BadRequest(c *gin.Context, message string, errDetail interface{}) {
	Error(c, http.StatusBadRequest, message, errDetail)
}

// Unauthorized mengirimkan response 401 Unauthorized
func Unauthorized(c *gin.Context, message string, errDetail interface{}) {
	Error(c, http.StatusUnauthorized, message, errDetail)
}

// InternalServerError mengirimkan response 500 Internal Server Error
func InternalServerError(c *gin.Context, message string, errDetail interface{}) {
	Error(c, http.StatusInternalServerError, message, errDetail)
}
