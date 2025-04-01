package handler

import (
	"github.com/RakhimovAns/Person-Service/internal/domain"
	"github.com/RakhimovAns/Person-Service/internal/service"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	service service.PersonService
	logger  logging.Logger
}

func NewPersonHandler(service service.PersonService, logger logging.Logger) *PersonHandler {
	return &PersonHandler{
		service: service,
		logger:  logger,
	}
}

func (h *PersonHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/people", h.GetAll)
		api.POST("/people", h.Create)
		api.GET("/people/:id", h.GetByID)
		api.PUT("/people/:id", h.Update)
		api.DELETE("/people/:id", h.Delete)
	}
}

// @Summary Create a new person
// @Description Create a new person with enriched data
// @Tags people
// @Accept json
// @Produce json
// @Param input body domain.PersonInput true "Person input"
// @Success 201 {object} domain.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /people [post]
func (h *PersonHandler) Create(c *gin.Context) {
	var input domain.PersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Debug("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	person, err := h.service.Create(input)
	if err != nil {
		h.logger.Error("Failed to create person: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

// GetAll godoc
// @Summary Get all people
// @Description Get list of people
// @Tags people
// @Produce json
// @Success 200 {array} domain.Person
// @Router /people [get]
func (h *PersonHandler) GetAll(c *gin.Context) {
	filter := domain.PersonFilter{
		Name:        getStringPointer(c.Query("name")),
		Surname:     getStringPointer(c.Query("surname")),
		Patronymic:  getStringPointer(c.Query("patronymic")),
		Gender:      getStringPointer(c.Query("gender")),
		Nationality: getStringPointer(c.Query("nationality")),
	}

	if ageStr := c.Query("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			h.logger.Debug("Invalid age parameter: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age parameter"})
			return
		}
		filter.Age = &age
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	people, err := h.service.GetAll(filter, page, limit)
	if err != nil {
		h.logger.Error("Failed to get people: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get people"})
		return
	}

	c.JSON(http.StatusOK, people)
}

// @Summary Get person by ID
// @Description Get person by ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} domain.Person
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /people/{id} [get]
func (h *PersonHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid ID parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	person, err := h.service.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get person by ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

// @Summary Update person
// @Description Update person by ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param input body domain.PersonInput true "Person input"
// @Success 200 {object} domain.Person
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /people/{id} [put]
func (h *PersonHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid ID parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	var input domain.PersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Debug("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	person, err := h.service.Update(id, input)
	if err != nil {
		h.logger.Error("Failed to update person with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

// @Summary Delete person
// @Description Delete person by ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /people/{id} [delete]
func (h *PersonHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid ID parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		h.logger.Error("Failed to delete person with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete person"})
		return
	}

	c.Status(http.StatusNoContent)
}

func getStringPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
