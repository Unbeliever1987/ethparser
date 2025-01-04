package ethblocksyncer

import (
	"github.com/Unbeliever1987/ethparser/internal/ethgateway"
	"github.com/Unbeliever1987/ethparser/internal/repository"
)

type Runner struct {
	repository repository.Repository
	ethGateway ethgateway.Gateway
}

func New(repository repository.Repository, ethGateway ethgateway.Gateway) Runner {
	return Runner{
		repository: repository,
		ethGateway: ethGateway,
	}
}
