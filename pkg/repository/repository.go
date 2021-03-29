package repository

// IUserRepository defines user operations
type IUserRepository interface {
	AddUser(UserEntity) (string, error)
	GetUser(string) (UserDocument, error)
	UpdateUser(string, UserEntity) error
}
