package response

type TaskResponse struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type ListTaskResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
type GetTaskResponse struct {
	Task TaskResponse `json:"task"`
}

type CreateTaskResponse struct {
	Task TaskResponse `json:"task"`
}

type UpdateTaskResponse struct {
	Task TaskResponse `json:"task"`
}
