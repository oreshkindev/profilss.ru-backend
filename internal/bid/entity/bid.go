package entity

import "time"

type (
	Bid struct {
		ID          uint64    `json:"id"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
		Description string    `json:"description"`
		Email       string    `json:"email,omitempty"`
		Fullname    string    `json:"fullname,omitempty"`
		Phone       string    `json:"phone,omitempty"`
		Status      bool      `json:"status" gorm:"default:false"`
	}

	BidUsecase interface {
		Create(*Bid) (*Bid, error)
		Find() ([]Bid, error)
		First(string) (*Bid, error)
		Update(*Bid, string) (*Bid, error)
		Delete(string) error
	}

	BidRepository interface {
		Create(*Bid) (*Bid, error)
		Find() ([]Bid, error)
		First(string) (*Bid, error)
		Update(*Bid, string) (*Bid, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Bid) NewResponse() *Bid {
	return &Bid{
		ID:          response.ID,
		CreatedAt:   response.CreatedAt,
		Email:       response.Email,
		Fullname:    response.Fullname,
		Phone:       response.Phone,
		Description: response.Description,
		Status:      response.Status,
	}
}
