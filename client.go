package moexalgopackgo

import (
	"github.com/DBoyara/MoexAlgoPackGo/models"
)

type IAlgoClient interface {
	GetSecurity() (models.Security, error)
}