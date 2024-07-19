package model

type Projects struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Client string `json:"client"`
	//CreatedAt time.Time `json:"createdAt"`
	Amount  float32 `json:"amount"`
	Tracked float32 `json:"tracked"`
}
