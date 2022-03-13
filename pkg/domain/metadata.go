package domain

type Metadata struct {
	Total int32 `json:"total"`
	CurrentPage int32 `json:"current_page"`
	HasPrevious bool `json:"has_previous"`
	HasNext bool `json:"has_next"`
	PreviousPage int32 `json:"previous_page"`
	NextPage int32 `json:"next_page"`
}