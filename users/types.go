package users

type ApiKey struct {
	Id   string
	Hash string
}

type UserRole int

const (
	admin UserRole = iota
	user
)

type User struct {
	Id       string
	Username string

	//bcrypt encoded password
	EncodedUserPassword []byte

	UserRole UserRole

	IsDeleted bool
	IsEnabled bool

	ApiKeys []ApiKey
}
