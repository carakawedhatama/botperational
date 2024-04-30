package bootstrap

import (
	"botperational/internal/adapter/repository"
	"botperational/internal/adapter/repository/sql"
	"botperational/internal/adapter/rest"
)

func RegisterDatabase() {
	appContainer.RegisterService("database", new(repository.Sqlx))
}

func RegisterCache() {
}

func RegisterRest() {
	appContainer.RegisterService("fiber", new(rest.Fiber))
}

func RegisterToggleService() {
}

func RegisterRepository() {
	appContainer.RegisterService("onLeaveRepository", new(sql.OnLeaveRepository))
	appContainer.RegisterService("onBirthdayRepository", new(sql.OnBirthdayRepository))
}
