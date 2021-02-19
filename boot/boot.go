package boot

import (
	"gf-web/library/validator"
	_ "gf-web/packed"
)

func init() {
	//validator.InitRequiredValidator()
	validator.InitArrayValidator()
}
