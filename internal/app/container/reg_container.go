package container

import "github.com/umardev500/pos/pkg"

func RegContainers(db *pkg.GormInstance, v pkg.Validator) []pkg.Container {
	return []pkg.Container{
		NewRoleContainer(db, v),
	}
}
