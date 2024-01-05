package controller

import (
	"gotodo-app/config"
	"gotodo-app/delivery/middleware"
	"gotodo-app/shared/common"
	"gotodo-app/shared/shared_model"
	"gotodo-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC       usecase.AuthorUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (a *AuthorController) listHandler(ctx *gin.Context) {
	author := ctx.MustGet("author").(string)
	authors, err := a.authorUC.FindAllAuthor(author)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range authors {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, shared_model.Paging{}, "Ok")
}

func (a *AuthorController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	author, err := a.authorUC.FindAuthorByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "author with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, author, "Ok")
}

func (a *AuthorController) Route() {
	a.rg.GET(config.AuthorGetList, a.authMiddleware.RequireToken("admin", "user"), a.listHandler)
	a.rg.GET(config.AuthorGetById, a.authMiddleware.RequireToken("admin", "user"), a.getHandler)
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *AuthorController {
	return &AuthorController{
		authorUC:       authorUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
