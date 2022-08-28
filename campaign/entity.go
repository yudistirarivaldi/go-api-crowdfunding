package campaign

import (
	"crowdfunding/user"
	"time"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string //benefit bagi si donator
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt		 time.Time
	CampaignImages	[]CampaignImages
	User 			user.User
}

type CampaignImages struct {
	ID	int
	CampaignID int
	FileName string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}