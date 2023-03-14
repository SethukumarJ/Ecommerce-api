package domain

type Admins struct {
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	UserName  string `bson:"name" json:"name"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Verification bool `bson:"verification" json:"verification"`
}

type AdminResponse struct {
	ID        *string `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string  `bson:"email" json:"email"`
	Password  string  `bson:"password" json:"password"`
	UserName  string  `bson:"name" json:"name"`
	FirstName string  `bson:"first_name" json:"first_name"`
	LastName  string  `bson:"last_name" json:"last_name"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Verification bool `bson:"verification" json:"verification"`
}
