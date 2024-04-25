package entity

import "time"

type (
	Product struct {
		ID              uint64                    `json:"id"`
		Characteristics []ProductsCharacteristics `json:"characteristics" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		CreatedAt       time.Time                 `json:"created_at" gorm:"default:now()"`
		Description     string                    `json:"description"`
		Published       bool                      `json:"published" gorm:"default:false"`
		Title           string                    `json:"title"`
	}

	ProductsCharacteristics struct {
		ID               uint64         `json:"id"`
		Characteristic   Characteristic `json:"characteristic" gorm:"foreignKey:CharacteristicID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		CharacteristicID uint64         `json:"-"`
		Measure          Measure        `json:"measure" gorm:"foreignKey:MeasureID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		MeasureID        uint64         `json:"-"`
		ProductID        uint64         `json:"-"`
		Value            string         `json:"value"`
	}

	ProductUsecase interface {
		Create(*Product) (*Product, error)
		Find() ([]Product, error)
		First(string) (*Product, error)
		Delete(string) error
	}

	ProductRepository interface {
		Create(*Product) (*Product, error)
		Find() ([]Product, error)
		First(string) (*Product, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
