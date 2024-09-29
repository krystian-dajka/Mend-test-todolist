package util

import (
	"github.com/krystian-dajka/Mend-test-todolist/models"
)

type ResMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResError struct {
	Success bool  `json:"success"`
	Error   error `json:"message"`
}

type ResUser struct {
	Success bool           `json:"success"`
	Message models.UserRes `json:"message"`
}

type ResTodo struct {
	Success bool        `json:"success"`
	Message models.Todo `json:"message"`
}

type ResTodos struct {
	Success bool          `json:"success"`
	Message []models.Todo `json:"message"`
}
