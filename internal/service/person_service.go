package service

import (
	"context"
	"github.com/RakhimovAns/Person-Service/internal/domain"
	"github.com/RakhimovAns/Person-Service/internal/repository"
	"github.com/RakhimovAns/Person-Service/pkg/client"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
)

type PersonService interface {
	Create(person domain.PersonInput) (domain.Person, error)
	GetAll(filter domain.PersonFilter, page, limit int) ([]domain.Person, error)
	GetByID(id int) (domain.Person, error)
	Update(id int, person domain.PersonInput) (domain.Person, error)
	Delete(id int) error
}

type personService struct {
	repo              repository.PersonRepository
	agifyClient       client.AgifyClient
	genderizeClient   client.GenderizeClient
	nationalizeClient client.NationalizeClient
	logger            logging.Logger
}

func NewPersonService(
	repo repository.PersonRepository,
	agifyClient client.AgifyClient,
	genderizeClient client.GenderizeClient,
	nationalizeClient client.NationalizeClient,
	logger logging.Logger,
) PersonService {
	return &personService{
		repo:              repo,
		agifyClient:       agifyClient,
		genderizeClient:   genderizeClient,
		nationalizeClient: nationalizeClient,
		logger:            logger,
	}
}

func (s *personService) Create(input domain.PersonInput) (domain.Person, error) {
	ctx := context.Background()

	age, err := s.agifyClient.GetAge(input.Name)
	if err != nil {
		s.logger.Error("Failed to get age: %v", err)
		return domain.Person{}, err
	}

	gender, err := s.genderizeClient.GetGender(input.Name)
	if err != nil {
		s.logger.Error("Failed to get gender: %v", err)
		return domain.Person{}, err
	}

	nationality, err := s.nationalizeClient.GetNationality(input.Name)
	if err != nil {
		s.logger.Error("Failed to get nationality: %v", err)
		return domain.Person{}, err
	}

	person := domain.Person{
		Name:        input.Name,
		Surname:     input.Surname,
		Patronymic:  input.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}

	id, err := s.repo.Create(ctx, person)
	if err != nil {
		return domain.Person{}, err
	}

	person.ID = id
	return person, nil
}

func (s *personService) GetAll(filter domain.PersonFilter, page, limit int) ([]domain.Person, error) {
	ctx := context.Background()
	return s.repo.GetAll(ctx, filter, page, limit)
}

func (s *personService) GetByID(id int) (domain.Person, error) {
	ctx := context.Background()
	return s.repo.GetByID(ctx, id)
}

func (s *personService) Update(id int, input domain.PersonInput) (domain.Person, error) {
	ctx := context.Background()

	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Person{}, err
	}

	age, err := s.agifyClient.GetAge(input.Name)
	if err != nil {
		s.logger.Error("Failed to get age: %v", err)
		return domain.Person{}, err
	}

	gender, err := s.genderizeClient.GetGender(input.Name)
	if err != nil {
		s.logger.Error("Failed to get gender: %v", err)
		return domain.Person{}, err
	}

	nationality, err := s.nationalizeClient.GetNationality(input.Name)
	if err != nil {
		s.logger.Error("Failed to get nationality: %v", err)
		return domain.Person{}, err
	}

	person := domain.Person{
		ID:          id,
		Name:        input.Name,
		Surname:     input.Surname,
		Patronymic:  input.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}

	if err := s.repo.Update(ctx, id, person); err != nil {
		return domain.Person{}, err
	}

	return person, nil
}

func (s *personService) Delete(id int) error {
	ctx := context.Background()
	return s.repo.Delete(ctx, id)
}
