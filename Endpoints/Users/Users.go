package Users

import (
	"QuaverGo"
)

type Users struct {
	APIClient         *QuaverGo.Client
	EndpointExtension string
}

func Init(apiClient *QuaverGo.Client) *Users {
	return &Users{APIClient: apiClient, EndpointExtension: "/user/"}
}
