package datastruct

type Person struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Verified bool   `json:"isVerified"`
	Role     Role   `json:"role"`
}

type Role string

const (
	ADMIN    Role = "admin"
	EMPLOYEE Role = "employee"
	USER     Role = "user"
)