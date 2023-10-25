package handler

import (
	"CRUD-GO/internal/product"
	"CRUD-GO/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProdService product.IService 
}

func (h *ProductHandler) GetById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return
	}

	productFound, errFound := h.ProdService.GetByIdentifier(id)
	if errFound != nil {
		if apiError, ok := errFound.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiError.Status, apiError)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &productFound)
}