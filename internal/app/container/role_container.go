package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/handler"
	"github.com/umardev500/pos/internal/app/repository"
	"github.com/umardev500/pos/internal/app/service"
	"github.com/umardev500/pos/pkg"
)

type roleContainer struct {
	roleHandler contract.RoleHandler
}

func NewRoleContainer(db *pkg.GormInstance, v pkg.Validator) pkg.Container {
	rr := repository.NewRoleRepository(db)
	rs := service.NewRoleService(rr, v)
	rh := handler.NewRoleHandler(rs)

	return &roleContainer{
		roleHandler: rh,
	}
}

func (rc *roleContainer) HandleApi(api fiber.Router) {
	role := api.Group("/roles")
	role.Post("/", rc.roleHandler.CreateRole)
	role.Delete("/", rc.roleHandler.DeleteRoles)
	role.Get("/", rc.roleHandler.GetRoles)
	role.Get("/:id", rc.roleHandler.GetRoleById)
}

func (rc *roleContainer) HandleWeb(web fiber.Router) {}
