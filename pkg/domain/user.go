package domain

type Users struct {
	Email        string `bson:"email" json:"email" validate:"required,email"`
	Password     string `bson:"password" json:"password" validate:"required,min=6"`
	UserName     string `bson:"name" json:"name" validate:"required"`
	FirstName    string `bson:"first_name" json:"first_name" `
	LastName     string `bson:"last_name" json:"last_name"`
	Profile      string `bson:"profile" json:"profile"`
	Verification bool   `bson:"verification" json:"verification"`
}

type UserResponse struct {
	ID           *string `bson:"_id,omitempty" json:"id,omitempty"`
	Email        string  `bson:"email" json:"email"`
	Password     string  `bson:"password" json:"password"`
	UserName     string  `bson:"name" json:"name"`
	FirstName    string  `bson:"first_name" json:"first_name"`
	LastName     string  `bson:"last_name" json:"last_name"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	Verification bool    `bson:"verification" json:"verification"`
}
