package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/handler"
	"github.com/umardev500/pos/internal/app/repository"
	"github.com/umardev500/pos/internal/app/service"
	"github.com/umardev500/pos/pkg"
)

type userContainer struct {
	userHandler contract.UserHandler
}

func NewUserContainer(db *pkg.GormInstance, v pkg.Validator) pkg.Container {
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur, v)
	uh := handler.NewUserHandler(us)

	return &userContainer{
		userHandler: uh,
	}
}

func (u *userContainer) HandleApi(api fiber.Router) {
	user := api.Group("/users")
	user.Get("/", u.userHandler.GetUsers)
}

func (u *userContainer) HandleWeb(web fiber.Router) {}
