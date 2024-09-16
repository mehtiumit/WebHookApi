package entities

import "time"

type Hook struct {
	ID        string    `json:"id" bson:"_id"`
	To        string    `json:"to" bson:"to"`
	ContentId string    `json:"content" bson:"content"`
	Body      string    `json:"body" bson:"body"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type HookRedisModel struct {
	HookId string    `json:"hook_id"`
	SentAt time.Time `json:"sent_at"`
}
