package request

import (
	"gf-web/library/global"
)

type IRequest interface {
	Rules() global.Rules
}

type ARequest struct {
}