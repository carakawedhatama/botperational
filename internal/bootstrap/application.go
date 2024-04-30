package bootstrap

import "botperational/internal/application/service"

func RegisterService() {
	appContainer.RegisterService("onLeaveService", new(service.OnLeaveService))
	appContainer.RegisterService("onBirthdayService", new(service.OnBirthdayService))
}
