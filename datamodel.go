package closeio

import "time"

type Contact struct {
	CreatedBy   string     `json:"created_by"`
	DateCreated *time.Time `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated"`
	Emails      []Email    `json:"emails"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Phones      []Phone    `json:"phones,omitempty"`
	Title       string     `json:"title"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
}

type Email struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Phone struct {
	Phone          string `json:"phone"`
	PhoneFormatted string `json:"phone_formatted"`
	Type           string `json:"type"`
}

type Activity struct {
	Attachments       []Attachment `json:"attachments"`
	CallDuration      int64        `json:"duration"`
	DateCreated       *time.Time   `json:"date_created"`
	Direction         string       `json:"direction"`
	ID                string       `json:"id"`
	OldStatus         string       `json:"old_status_label"`
	OldStatusID       string       `json:"old_status_id"`
	NewStatus         string       `json:"new_status_label"`
	NewStatusID       string       `json:"new_status_id"`
	Type              string       `json:"_type"`
	UserID            string       `json:"user_id"`
	UserName          string       `json:"user_name"`
	VoiceMailDuration int64        `json:"voicemail_duration"`
}

type Attachment struct {
	ContentID   string `json:"content_id"`
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
}

type Lead struct {
	Addresses      []Address         `json:"addresses,omitempty"`
	Contacts       []Contact         `json:"contacts,omitempty"`
	CreationDate   *time.Time        `json:"date_created,omitempty"`
	Custom         map[string]string `json:"custom,omitempty"`
	Description    string            `json:"description,omitempty"`
	DisplayName    string            `json:"display_name,omitempty"`
	HtmlURL        string            `json:"html_url,omitempty"`
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	OrganizationID string            `json:"organization_id,omitempty"`
	Opportunities  []Opportunity     `json:"opportunities,omitempty"`
	StatusID       string            `json:"status_id,omitempty"`
	StatusLabel    string            `json:"status_label,omitempty"`
	Tasks          []Task            `json:"tasks"`
	UpdatedBy      string            `json:"updated_by,omitempty"`
	UpdatedByName  string            `json:"updated_by_name,omitempty"`
	URL            string            `json:"url,omitempty"`
}

type Task struct {
	Type           string     `json:"_type"`
	AssignedTo     string     `json:"assigned_to,omitempty"`
	AssignedToName string     `json:"assigned_to_name,omitempty"`
	ContactID      string     `json:"contact_id,omitempty"`
	ContactName    string     `json:"contact_name,omitempty"`
	CreatedBy      string     `json:"created_by"`
	CreatedName    string     `json:"created_by_name"`
	DateCreated    *time.Time `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated"`
	DueDate        *time.Time `json:"date_updated,omitempty"`
	ID             string     `json:"id,omitempty"`
	IsComplete     bool       `json:"is_complete"`
	IsDateLess     bool       `json:"is_dateless"`
	LeadID         string     `json:"lead_id,omitempty"`
	LeadName       string     `json:"lead_name,omitempty"`
	ObjectID       string     `json:"object_id,omitempty"`
	ObjectType     string     `json:"object_type,omitempty"`
	OrganizationID string     `json:"organization_id,omitempty"`
	Text           string     `json:"text,omitempty"`
	UpdatedBy      string     `json:"updated_by,omitempty"`
	UpdatedByName  string     `json:"updated_by_name,omitempty"`
	View           string     `json:"view,omitempty"`
}

type Opportunity struct {
	Confidence     int        `json:"confidence"`
	ContactID      string     `json:"contact_id"`
	ContactName    string     `json:"contact_name"`
	CreationDate   *time.Time `json:"date_created,omitempty"`
	CreatedBy      string     `json:"created_by"`
	CreatedName    string     `json:"created_by_name"`
	DateCreated    *time.Time `json:"date_created"`
	DateLost       *time.Time `json:"date_lost"`
	DateUpdated    *time.Time `json:"date_updated"`
	DateWon        *time.Time `json:"date_won"`
	ID             string     `json:"id"`
	LeadID         string     `json:"lead_id"`
	LeadName       string     `json:"lead_name"`
	Note           string     `json:"note,omitempty"`
	OrganizationID string     `json:"organization_id,omitempty"`
	StatusID       string     `json:"status_id,omitempty"`
	StatusLabel    string     `json:"status_label,omitempty"`
	StatusType     string     `json:"status_type,omitempty"`
	UpdatedBy      string     `json:"updated_by,omitempty"`
	UpdatedByName  string     `json:"updated_by_name,omitempty"`
	UserID         string     `json:"user_id"`
	UserName       string     `json:"user_name"`
	Value          float64    `json:"value"`
	ValueCurrency  string     `json:"value_currency"`
	ValueFormatted string     `json:"value_formatted"`
	ValuePeriod    string     `json:"value_period"`
}
