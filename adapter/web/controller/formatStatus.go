package controller

import (
	"github.com/renatospaka/scheduler/core/dto"
)

func FormatStatus(status string, code int, msg string) dto.StandardStatusOutputDto {
	statusDto := dto.StandardStatusOutputDto{}
	statusErrDto := dto.StandardErrorOutputDto{}

	if status == "" {
		statusDto.Status = REQUEST_FAILURE
	} else {
		statusDto.Status = status
	}
	
	statusErrDto.Code = code
	statusErrDto.Message = msg
	statusDto.Error = statusErrDto
	return statusDto
}
