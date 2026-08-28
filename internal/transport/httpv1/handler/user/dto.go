package user

type RegisterInput struct {
	Body struct {
		FirstName string `json:"firstname" minLength:"1" maxLength:"50" pattern:"^[a-zA-Z]+$"`
		LastName  string `json:"lastname"  minLength:"1" maxLength:"50" pattern:"^[a-zA-Z]+$"`
		Username  string `json:"username"  minLength:"3" maxLength:"12" pattern:"^[a-zA-Z0-9_]+$"`
		Email     string `json:"email"     format:"email"`
		Password  string `json:"password"  minLength:"7" maxLength:"64"`
	}
}

type RegisterOutput struct {
	Body struct {
		Message     string `json:"message"`
		AccessToken string `json:"accessToken"`
	}
}

type LoginInput struct {
	Body struct {
		Email    string `json:"email"     format:"email"`
		Password string `json:"password"  minLength:"7" maxLength:"64"`
	}
}

type LoginOutput struct {
	Body struct {
		AccessToken string `json:"accessToken"`
	}
}
