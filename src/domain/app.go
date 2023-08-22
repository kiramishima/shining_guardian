package domain

import "github.com/ServiceWeaver/weaver"

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type App struct {
	weaver.Implements[weaver.Main]
	Server weaver.Listener
}
