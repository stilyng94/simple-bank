package api

import (
	"simple-bank/ent"
	"simple-bank/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router        *gin.Engine
	iAccountRepo  repository.IAccountRepo
	iTransferRepo repository.ITransferRepo
	dbClient      *ent.Client
}

func NewServer(dbClient *ent.Client) *Server {
	server := &Server{}
	router := gin.Default()
	server.dbClient = dbClient
	server.init()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.getAccounts)
	router.GET("/accounts/:id", server.getAccount)

	server.Router = router
	return server
}

func (server *Server) init() {
	server.iAccountRepo = repository.NewAccountRepo(server.dbClient)
	server.iTransferRepo = repository.NewTransferRepo(server.dbClient)
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
