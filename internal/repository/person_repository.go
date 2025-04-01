package repository

import (
	"context"
	"fmt"
	"github.com/RakhimovAns/Person-Service/internal/config"
	"github.com/RakhimovAns/Person-Service/internal/domain"
	"github.com/RakhimovAns/Person-Service/pkg/client/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

type PersonRepository interface {
	Create(ctx context.Context, person domain.Person) (int, error)
	GetAll(ctx context.Context, filter domain.PersonFilter, page, limit int) ([]domain.Person, error)
	GetByID(ctx context.Context, id int) (domain.Person, error)
	Update(ctx context.Context, id int, person domain.Person) error
	Delete(ctx context.Context, id int) error
}

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

type personRepository struct {
	db     *sqlx.DB
	logger logging.Logger
}

func NewPersonRepository(db *sqlx.DB, logger logging.Logger) PersonRepository {
	return &personRepository{
		db:     db,
		logger: logger,
	}
}

func (r *personRepository) Create(ctx context.Context, person domain.Person) (int, error) {
	query := `INSERT INTO people (name, surname, patronymic, age, gender, nationality) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query,
		person.Name,
		person.Surname,
		person.Patronymic,
		person.Age,
		person.Gender,
		person.Nationality,
	).Scan(&id)

	if err != nil {
		r.logger.Error("Failed to create person: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *personRepository) GetAll(ctx context.Context, filter domain.PersonFilter, page, limit int) ([]domain.Person, error) {
	query := `SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	if filter.Name != nil {
		query += ` AND name = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Name)
		argPos++
	}

	if filter.Surname != nil {
		query += ` AND surname = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Surname)
		argPos++
	}

	if filter.Patronymic != nil {
		query += ` AND patronymic = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Patronymic)
		argPos++
	}

	if filter.Age != nil {
		query += ` AND age = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Age)
		argPos++
	}

	if filter.Gender != nil {
		query += ` AND gender = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Gender)
		argPos++
	}

	if filter.Nationality != nil {
		query += ` AND nationality = $` + strconv.Itoa(argPos)
		args = append(args, *filter.Nationality)
		argPos++
	}

	query += ` LIMIT $` + strconv.Itoa(argPos) + ` OFFSET $` + strconv.Itoa(argPos+1)
	args = append(args, limit, (page-1)*limit)

	var people []domain.Person
	err := r.db.SelectContext(ctx, &people, query, args...)
	if err != nil {
		r.logger.Error("Failed to get all people: %v", err)
		return nil, err
	}

	return people, nil
}

func (r *personRepository) GetByID(ctx context.Context, id int) (domain.Person, error) {
	query := `SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE id = $1`

	var person domain.Person
	err := r.db.GetContext(ctx, &person, query, id)
	if err != nil {
		r.logger.Error("Failed to get person by ID %d: %v", id, err)
		return domain.Person{}, err
	}

	return person, nil
}

func (r *personRepository) Update(ctx context.Context, id int, person domain.Person) error {
	query := `UPDATE people SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query,
		person.Name,
		person.Surname,
		person.Patronymic,
		person.Age,
		person.Gender,
		person.Nationality,
		id,
	)

	if err != nil {
		r.logger.Error("Failed to update person with ID %d: %v", id, err)
		return err
	}

	return nil
}

func (r *personRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM people WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete person with ID %d: %v", id, err)
		return err
	}

	return nil
}
