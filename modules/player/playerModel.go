package player

import "time"

type (
	Profile struct {
		Id        string    `json:"id"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Claims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}

	CreatePlayerReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
		Username string `json:"username" form:"username" validate:"required,max=64"`
	}

	CreatePlayerTransactionReq struct {
		PlayerId string  `json:"player_id" validate:"required,max=64"`
		Amount   float64 `json:"amount" validate:"required"`
	}
)
