package usecase

import (
	"gotodo-app/model/dto"
	"gotodo-app/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	authorUc   AuthorUseCase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	author, err := a.authorUc.FindAuthorByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(author)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUseCase(authorUc AuthorUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{authorUc: authorUc, jwtService: jwtService}
}
