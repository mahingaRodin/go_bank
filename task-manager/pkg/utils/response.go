package utils

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// Success sends a successful response
func Success(c *gin.Context, data interface{}, message string) {
    c.JSON(http.StatusOK, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

// Created sends a resource creation response
func Created(c *gin.Context, data interface{}) {
    c.JSON(http.StatusCreated, Response{
        Success: true,
        Message: "Resource created successfully",
        Data:    data,
    })
}

// Error sends an error response
func Error(c *gin.Context, statusCode int, errorMessage string) {
    c.JSON(statusCode, Response{
        Success: false,
        Error:   errorMessage,
    })
}

// NotFound sends a 404 response
func NotFound(c *gin.Context, resource string) {
    Error(c, http.StatusNotFound, resource+" not found")
}

// BadRequest sends a 400 response
func BadRequest(c *gin.Context, message string) {
    Error(c, http.StatusBadRequest, message)
}

// InternalError sends a 500 response
func InternalError(c *gin.Context, message string) {
    Error(c, http.StatusInternalServerError, message)
}