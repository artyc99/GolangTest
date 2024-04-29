package users

type User struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Age           int    `json:"age"`
	RecordingDate int64  `json:"recording_date"`
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type GetListReq struct {
	FromAge *int `json:"from_age,omitempty"`
	ToAge   *int `json:"to_age,omitempty"`

	FromDate *int64 `json:"from_date,omitempty"`
	ToDate   *int64 `json:"to_date,omitempty"`
}

type GetListRes struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type CreateReq struct {
	ID            string `json:"id" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Age           int    `json:"age" validate:"required"`
	RecordingDate int64  `json:"recording_date" validate:"required"`
}

type CreateRes struct {
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
