package types

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}
