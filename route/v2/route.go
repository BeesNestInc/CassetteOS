package v2

import (
	"github.com/BeesNestInc/CassetteOS/codegen"
	"github.com/BeesNestInc/CassetteOS/service"
)

type CasaOS struct {
	fileUploadService *service.FileUploadService
}

func NewCasaOS() codegen.ServerInterface {
	return &CasaOS{
		fileUploadService: service.NewFileUploadService(),
	}
}
