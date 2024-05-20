package models

import "time"

type TodoList struct {
    ID                int       `json:"id"`
    UserId            int       `json:"user_id"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
    DeletedAt         *time.Time `json:"deleted_at"`
    CompletionPercent int       `json:"completion_percent"`
}

type TodoMessage struct {
    ID        int       `json:"id"`
    ListID    int       `json:"list_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
    Content   string    `json:"content"`
    IsDone    bool      `json:"is_done"`
}