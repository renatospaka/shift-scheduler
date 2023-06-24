package controller

import (
	pkgController "github.com/renatospaka/scheduler/adapter/web/controller"
	"github.com/renatospaka/scheduler/core/dto"
)

func formatStatus(status string, code int, msg string) dto.StandardStatusOutputDto {
	statusDto := dto.StandardStatusOutputDto{}
	statusErrDto := dto.StandardErrorOutputDto{}

	if status == "" {
		statusDto.Status = pkgController.REQUEST_FAILURE
	} else {
		statusDto.Status = status
	}
	
	statusErrDto.Code = code
	statusErrDto.Message = msg
	statusDto.Error = statusErrDto
	return statusDto
}
