package v1

// `POST /login` 接口的请求参数
type LoginRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)" binding:"required"`
	Password string `json:"password" valid:"required,stringlength(6|18)" binding:"required"`
}

// `POST /login` 接口的返回参数
type LoginResponse struct {
	Token string `json:"token"`
}

// `POST /v1/users/{name}/change-password` 接口的请求参数.
type ChangePasswordRequest struct {
	// 旧密码.
	OldPassword string `json:"old_password" valid:"required,stringlength(6|18)"`

	// 新密码.
	NewPassword string `json:"new_password" valid:"required,stringlength(6|18)"`
}

// `POST /v1/users` 接口的请求参数
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)" binding:"required"`
	Password string `json:"password" valid:"required,stringlength(6|18)"           binding:"required"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"          binding:"required"`
	Email    string `json:"email"    valid:"required,email"                        binding:"required"`
	Phone    string `json:"phone"    valid:"required,stringlength(11|11)"          binding:"required"`
}

