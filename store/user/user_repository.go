package user

type Repository interface {
	CreateUser(data User) (id string, createdAt string, err error)
}
