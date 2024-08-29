package handlers

import (
	"net/http"
	"social-sys/internal/api/dto"
	"social-sys/internal/models"
	"social-sys/internal/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service service.PostServiceInterface
}

func NewPostHandler(service service.PostServiceInterface) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var post models.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.service.CreatePost(&post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &post)
}

func (h *PostHandler) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	post, err := h.service.GetPost(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *PostHandler) ListPosts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	posts, total, err := h.service.ListPosts(page, pageSize)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	postListRes := dto.PostListRes{
		Posts: posts,
		Total: total,
	}

	ctx.JSON(http.StatusOK, postListRes)

}

func (h *PostHandler) UpdatePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	var post models.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	post.ID = uint(id)

	if err := h.service.UpdatePost(&post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, post)

}

func (h *PostHandler) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	if err := h.service.DeletePost(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
