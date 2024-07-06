package moexalgopackgo

import (
	"moexalgopackgo/models"
)

type IAlgoClient interface {
	GetSecurity() (models.Security, error)
}