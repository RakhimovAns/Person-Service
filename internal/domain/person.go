package domain

type Person struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Patronymic  *string `json:"patronymic,omitempty"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Nationality string  `json:"nationality"`
}

type PersonInput struct {
	Name       string  `json:"name" validate:"required"`
	Surname    string  `json:"surname" validate:"required"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type PersonFilter struct {
	Name        *string `json:"name,omitempty"`
	Surname     *string `json:"surname,omitempty"`
	Patronymic  *string `json:"patronymic,omitempty"`
	Age         *int    `json:"age,omitempty"`
	Gender      *string `json:"gender,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
}
