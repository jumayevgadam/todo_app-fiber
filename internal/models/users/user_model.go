package users

// We use DTO and DAO models again
// SignUpReq is DTO model
type SignUpReq struct {
	Username string `form:"userName" json:"userName"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// SignUpRes is DAO model
type SignUpRes struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// ToServer is
func (s *SignUpRes) ToServer() *SignUpReq {
	return &SignUpReq{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}
}

// ToStorage is
func (s *SignUpReq) ToStorage() *SignUpRes {
	return &SignUpRes{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}
}
