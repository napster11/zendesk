package zendeskModel

import (
	"time"
)

//APIResponse struct for the API Response of the GetTicketList API
type APIResponse struct {
	Tickets []TicketList `json:"tickets"`
}

//TicketList struct for the object of API Response List
type TicketList struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	AssigneeID  int       `json:"assignee_id"`
	BrandID     int       `json:"brand_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tags        []string  `json:"tags"`
	SubmitterID int       `json:"submitter_id"`
	Status      string    `json:"status"`
	URL         string    `json:"url"`
	RequesterID int       `json:"requester_id"`
}
