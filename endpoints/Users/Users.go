package Users

import "quaverGo"

type Users struct {
	APIClient         *quaverGo.Client
	EndpointExtension string
}

func Init(apiClient *quaverGo.Client) *Users {
	return &Users{APIClient: apiClient, EndpointExtension: "/user/"}
}
