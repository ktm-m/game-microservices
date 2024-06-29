package server

import (
	"github.com/Montheankul-K/game-microservices/modules/auth/authHandler"
	"github.com/Montheankul-K/game-microservices/modules/auth/authRepository"
	"github.com/Montheankul-K/game-microservices/modules/auth/authUsecase"
)

func (s *server) authService() {
	repository := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repository)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")
	auth.GET("", s.healthCheckService)
}
