package modules

import (
	"EchoAPI/core/modules/users"
	"EchoAPI/core/repositories"
)

type Modules struct {
	Users *users.Users
}

func New(repoFactory repositories.RepoFactoryI) *Modules {
	return &Modules{
		Users: users.New(repoFactory),
	}
}
