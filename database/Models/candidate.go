package Models

type Candidate struct {
	Id         string `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Experience int64  `json:"experience"`
	Timestamps Timestamps
}

type Timestamps struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListCandidatesResponse struct {
	Candidates []Candidate `json:"candidates"`
}
