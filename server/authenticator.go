package server

import "github.com/aonx/momonga/configuration"

//type User struct {
//	Name []byte
//}

type Authenticator interface {
	Init(config *configuration.Config)
	Authenticate(user_id, password []byte) (bool, error)
	Shutdown()
}
