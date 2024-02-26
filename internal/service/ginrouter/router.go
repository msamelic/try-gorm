package ginrouter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"try-gorm/internal/model"
	"try-gorm/internal/util/envconf"
)

func New(repo model.Producer, env *envconf.Spec, log *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	NewHandler(repo, r, env, log)

	return r
}

type Handler struct {
	Router *gin.Engine
	Repo   model.Producer
	Env    *envconf.Spec
	Log    *zap.Logger
}

func NewHandler(repo model.Producer, r *gin.Engine, env *envconf.Spec, log *zap.Logger) {
	h := &Handler{
		Router: r,
		Repo:   repo,
		Env:    env,
		Log:    log,
	}
	r.GET("/health", h.HealthCheck())

	r.GET("/products", h.ListProducts())
	r.GET("/products/:id", h.GetProduct())
	r.POST("/products", h.CreateProduct())
	r.PUT("/products", h.UpdateProduct())
	r.DELETE("/products", h.DeleteProduct())
}

func (h *Handler) HealthCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}

func (h *Handler) ListProducts() func(c *gin.Context) {
	return func(c *gin.Context) {
		products, err := h.Repo.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, products)
	}
}

func (h *Handler) GetProduct() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		product, err := h.Repo.Get(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, nil)
				return
			}
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func (h *Handler) CreateProduct() func(c *gin.Context) {
	return func(c *gin.Context) {
		var product model.Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		product, err = h.Repo.Create(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func (h *Handler) UpdateProduct() func(c *gin.Context) {
	return func(c *gin.Context) {
		var product model.Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		product, err = h.Repo.Update(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func (h *Handler) DeleteProduct() func(c *gin.Context) {
	return func(c *gin.Context) {
		var product model.Product
		err := c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := h.Repo.Delete(product); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusNoContent, product)
	}
}
