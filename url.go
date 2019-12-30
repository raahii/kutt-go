package kutt

import "time"

type URL struct {
	ID         string    `json:"id"`
	Target     string    `json:"target"`
	ShortURL   string    `json:"shortUrl"`
	Password   bool      `json:"password"`
	Reuse      bool      `json:"reuse"`
	DomainID   string    `json:"domain_id"`
	VisitCount int       `json:"visit_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
