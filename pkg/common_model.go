package pkg

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IDsReq struct {
	IDs []uuid.UUID `json:"ids" validate:"required,min=1"`
}

type Response struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Code       string      `json:"code,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
	Ref        interface{} `json:"ref,omitempty"`
}

type TimeCommon struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type ValidationErr struct {
	Tag     string      `json:"tag"`
	Kind    string      `json:"kind"`
	Path    string      `json:"path"`
	Options interface{} `json:"options,omitempty"`
	Message string      `json:"message"`
}
