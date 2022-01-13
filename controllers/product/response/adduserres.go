package response

import "time"

type ResponseProduct struct {
	Id          int       `json:"id"`
	NameProduct string    `json:"productName"`
	Poin        int       `json:"poin"`
	Amount      string    `json:"amount"`
	Img         string    `json:"img"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
