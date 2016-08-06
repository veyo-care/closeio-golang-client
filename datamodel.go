package closeio

import "time"

type Contact struct {
	Emails []Email `json:"emails"`
	Phones []Phone `json:"phones,omitempty"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
}

type Email struct {
	Email string `json:"email"`
}

type Phone struct {
	Phone string `json:"phone"`
}

type Activity struct {
	UserId            string     `json:"user_id"`
	ID                string     `json:"id"`
	DateCreated       *time.Time `json:"date_created"`
	UserName          string     `json:"user_name"`
	Type              string     `json:"_type"`
	OldStatus         string     `json:"old_status_label"`
	OldStatusID       string     `json:"old_status_id"`
	NewStatus         string     `json:"new_status_label"`
	NewStatusID       string     `json:"new_status_id"`
	CallDuration      int64      `json:"duration"`
	VoiceMailDuration int64      `json:"voicemail_duration"`
	Direction         string     `json:"direction"`
}

type Lead struct {
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Tasks         []Task            `json:"tasks"`
	CreationDate  *time.Time        `json:"date_created,omitempty"`
	DisplayName   string            `json:"display_name,omitempty"`
	StatusID      string            `json:"status_id,omitempty"`
	StatusLabel   string            `json:"status_label,omitempty"`
	Custom        map[string]string `json:"custom"`
	Opportunities []Opportunity     `json:"opportunities,omitempty"`
	Contacts      []Contact         `json:"contacts,omitempty"`
	Addresses     []Address         `json:"addresses,omitempty"`
}

type Task struct {
}

type Opportunity struct {
	Confidence   int        `json:"confidence"`
	CreationDate *time.Time `json:"date_created,omitempty"`
	ID           string     `json:"id"`
	LeadID       string     `json:"lead_id"`
	LeadName     string     `json:"lead_name"`
	StatusID     string     `json:"status_id"`
	StatusType   string     `json:"status_type"`
	Status       string     `json:"status_label"`
	UserID       string     `json:"user_id"`
	UserName     string     `json:"user_name"`
}
