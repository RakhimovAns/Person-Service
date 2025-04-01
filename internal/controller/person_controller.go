package controller

import (
	"github.com/RakhimovAns/Person-Service/internal/domain"
	"github.com/RakhimovAns/Person-Service/internal/service"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
)

type PersonController interface {
	Create(person domain.PersonInput) (domain.Person, error)
	GetAll(filter domain.PersonFilter, page, limit int) ([]domain.Person, error)
	GetByID(id int) (domain.Person, error)
	Update(id int, person domain.PersonInput) (domain.Person, error)
	Delete(id int) error
}

type personController struct {
	service service.PersonService
	logger  logging.Logger
}

func NewPersonController(service service.PersonService, logger logging.Logger) PersonController {
	return &personController{
		service: service,
		logger:  logger,
	}
}

func (c *personController) Create(person domain.PersonInput) (domain.Person, error) {
	c.logger.Debug("Creating person: %+v", person)
	return c.service.Create(person)
}

func (c *personController) GetAll(filter domain.PersonFilter, page, limit int) ([]domain.Person, error) {
	c.logger.Debug("Getting all persons with filter: %+v, page: %d, limit: %d", filter, page, limit)
	return c.service.GetAll(filter, page, limit)
}

func (c *personController) GetByID(id int) (domain.Person, error) {
	c.logger.Debug("Getting person by ID: %d", id)
	return c.service.GetByID(id)
}

func (c *personController) Update(id int, person domain.PersonInput) (domain.Person, error) {
	c.logger.Debug("Updating person with ID: %d, data: %+v", id, person)
	return c.service.Update(id, person)
}

func (c *personController) Delete(id int) error {
	c.logger.Debug("Deleting person with ID: %d", id)
	return c.service.Delete(id)
}
