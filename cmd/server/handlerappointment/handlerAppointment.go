package handlerappointment

import (
	"CRUD-FINAL/internal/appointments"
	"CRUD-FINAL/internal/domain"
	"CRUD-FINAL/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointHandler struct {
	ApointmentService appointments.IService
}

func (s *AppointHandler) GetAll(ctx *gin.Context ) {
  appointmentsFound, errFound := s.ApointmentService.GetAllAppointments()
	if errFound != nil {
		if apiError, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiError.Status, apiError)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
		return
	}
  ctx.JSON(http.StatusOK, &appointmentsFound)
}


func (h *AppointHandler) GetById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	appointmentFound, errFound := h.ApointmentService.GetByIdentifier(id)
	if errFound != nil {
		if apiError, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiError.Status, apiError)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
		return
	}

	ctx.JSON(http.StatusOK, &appointmentFound)
}

// func (h *DentistHandler) DeleteById(ctx *gin.Context){
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
// 		return
// 	}

// 	idFound, errFound := h.DentistService.DeleteDentist(id)
// 	if errFound != nil {
// 		if apiError, ok := errFound.(*web.ErrorApi); ok {
// 			ctx.AbortWithStatusJSON(apiError.Status, apiError)
// 			return
// 		}
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errFound)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, &idFound)

// }

func (h *AppointHandler) SaveAppointment(ctx *gin.Context){
	 var requestBody domain.Appointment 
	err := ctx.ShouldBindJSON(&requestBody)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    } 

	h.ApointmentService.SaveAppointment(&requestBody)
    ctx.JSON(http.StatusOK, gin.H{"message": "Resource created successfully"})
}

// func (h *DentistHandler) UpdateDentist(ctx * gin.Context){
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
// 		return
// 	}
// 	var requestBody domain.Dentist 
// 	err2 := ctx.ShouldBindJSON(&requestBody)
//     if err2 != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
// 	if id == requestBody.Id {
// 	  idUpdated, err3 := h.DentistService.UpdateDentist(requestBody)
// 	  if err3 != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
// 		return
// 	  }

// 	 ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Resource with id %v updated successfully", idUpdated)})
// 	}

// }