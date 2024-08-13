package usecase

type PasswordHasher interface {
	ComparePassword(hashedPassword, password string) error
	EncryptPassword(password string) (string, error)
}

type TokenGenerator interface {
	CreateToken(userID, role, email string) (string, string, error)
}

type InputValidator interface {
	IsValidEmail(email string) bool
	IsValidPassword(password string) bool
}
