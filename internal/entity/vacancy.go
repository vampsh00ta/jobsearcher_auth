package entity

type Vacancy struct {
	ID             int      `json:"id" db:"id"`
	SlugID         string   `json:"slug_id" db:"slug_id"`
	Name           string   `json:"name" db:"name"`
	Cities         []string `json:"city" db:"city"`
	CitiesIDs      []string `json:"city_id" db:"city_id"`
	Category       string   `json:"category" db:"category"`
	CategorySlug   string   `json:"category_slug" db:"category_slug"`
	Description    string   `json:"description" db:"description"`
	ExperienceSlug string   `json:"experience_slug" db:"experience_slug"`
	ExperienceName string   `json:"experience_name" db:"experience_name"`
	CompanySlug    string   `json:"company_slug" db:"company_slug"`
	CompanyName    string   `json:"company_name" db:"company_name"`
	Link           string   `json:"link" db:"link"`
	Rank           float64  `json:"rank" db:"rank"`
}
