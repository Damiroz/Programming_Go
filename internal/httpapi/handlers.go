package httpapi

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"example.com/pract16/internal/models"
	"example.com/pract16/internal/service"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) CreateNote(c *gin.Context) {
	var n models.Note
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	if err := h.Service.Create(c.Request.Context(), &n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, n)
}

func (h *Handler) GetNote(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	n, err := h.Service.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, n)
}

func (h *Handler) UpdateNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var n models.Note
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	if err := h.Service.Update(c.Request.Context(), id, n.Title, n.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.Service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) ListNotes(c *gin.Context) {
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	offset, _ := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)
	
	notes, err := h.Service.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}