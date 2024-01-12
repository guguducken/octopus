package auth

import (
	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/user"
)

func GetAuthenticatedUser(cfg *config.Config) (u *common.User, err error) {
	return user.GetSelf(cfg)
}
