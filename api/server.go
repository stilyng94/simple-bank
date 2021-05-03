package api

import (
	"simple-bank/ent"
	"simple-bank/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	iAccountRepo repository.IAccountRepo
}

func NewServer(dbClient *ent.Client) *Server {
	server := &Server{}
	router := gin.Default()
	iAccountRepo := repository.NewAccountRepo(dbClient)
	server.iAccountRepo = iAccountRepo

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.getAccounts)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
