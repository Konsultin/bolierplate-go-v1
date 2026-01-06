package dto

// ===== Login with Password =====

type LoginPassword_Payload struct {
	Identifier string         `json:"identifier" validate:"required,min=3,max=255"` // username, email, or phone
	Password   string         `json:"password" validate:"required,min=6,max=128"`
	Device     *DeviceSession `json:"device,omitempty" validate:"omitempty"`
}

// ===== Login with OAuth =====

type LoginOAuth_Payload struct {
	Provider AuthProvider_Enum `json:"provider" validate:"required,oneof=2 3 4"` // 2=GOOGLE, 3=FACEBOOK, 4=APPLE
	IdToken  string            `json:"idToken" validate:"required"`              // OAuth ID token from provider
	Device   *DeviceSession    `json:"device,omitempty" validate:"omitempty"`
}

// ===== Register User =====

type RegisterUser_Payload struct {
	FullName string `json:"fullName" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"omitempty,email,max=255"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
	Username string `json:"username" validate:"omitempty,min=3,max=100,alphanum"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

// ===== OAuth User Info (from provider) =====

type OAuthUserInfo struct {
	ProviderId    string `json:"providerId"` // Provider's user ID
	Email         string `json:"email"`      // User's email from provider
	Name          string `json:"name"`       // User's name from provider
	Picture       string `json:"picture"`    // Profile picture URL
	EmailVerified bool   `json:"emailVerified"`
}
