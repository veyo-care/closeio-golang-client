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
	Type string `json:"_type"`
	ID   string `json:"id"`

	//Email

	Attachments         []Attachment `json:"attachments"`
	Bcc                 []string     `json:"bcc"`
	BodyHtml            string       `json:"body_html"`
	BodyHtmlQuoted      HtmlQuoted   `json:"body_html_quoted"`
	BodyPreview         string       `json:"body_preview"`
	BodyText            string       `json:"body_text"`
	BodyTextQuoted      TextQuoted   `json:"body_text_quoted"`
	Cc                  []string     `json:"cc"`
	ContactID           string       `json:"contact_id"`
	DateScheduled       *time.Time   `json:"date_scheduled"`
	DateSent            *time.Time   `json:"date_sent"`
	EmailAccountID      string       `json:"email_account_id"`
	Enveloppe           Enveloppe    `json:"enveloppe"`
	InReplyToID         string       `json:"in_reply_to_id"`
	LeadID              string       `json:"lead_id"`
	MessageIDs          []string     `json:"message_ids"`
	NeedSmtpCredentials bool         `json:"need_smtp_credentials"`
	Opens               Opens        `json:"opens"`
	OpensSummary        string       `json:"opens_summary"`
	References          []string     `json:"references"`
	SendAttempts        []string     `json:"send_attempts"`
	Sender              string       `json:"sender"`
	Status              string       `json:"status"`
	Subject             string       `json:"subject"`
	TemplateID          string       `json:"template_id"`
	TemplateName        string       `json:"template_name"`
	ThreadID            string       `json:"thread_id"`
	To                  []string     `json:"to"`

	//LeadStatusChange

	NewStatusID string `json:"new_status_id"`
	NewStatus   string `json:"new_status_label"`
	OldStatusID string `json:"old_status_id"`
	OldStatus   string `json:"old_status_label"`

	//OpportunityStatusChange

	NewStatusType             string     `json:"new_status_type"`
	OldStatusType             string     `json:"old_status_type"`
	OpportunityConfidence     int        `json:"opportunity_confidence"`
	OpportunityDateWon        *time.Time `json:"opportunity_date_won"`
	OpportunityID             string     `json:"opportunity_id"`
	OpportunityValue          int        `json:"opportunity_value"`
	OpportunityValueCurrency  string     `json:"opportunity_value_currency"`
	OpportunityValueFormatted string     `json:"opportunity_value_formatted"`
	OpportunityValuePeriod    string     `json:"opportunity_value_period"`

	//Call
	Direction         string `json:"direction"`
	CallDuration      int64  `json:"duration"`
	LocalPhone        string `json:"local_phone"`
	Note              string `json:"note"`
	Phone             string `json:"phone"`
	RecordingUrl      string `json:"recording_url"`
	RemotePhone       string `json:"remote_phone"`
	Source            string `json:"source"`
	TransferredFrom   string `json:"transferred_from"`
	TransferredTo     string `json:"transferred_to"`
	VoiceMailDuration int64  `json:"voicemail_duration"`
	VoiceMailUrl      string `json:"voicemail_url"`

	//TaskCompleted

	TaskAssignedTo     string `json:"task_assigned_to"`
	TaskAssignedToName string `json:"task_assigned_to_name"`
	TaskID             string `json:"task_id"`
	TaskText           string `json:"task_text"`

	//Created

	ImportID string `json:"import_id"`

	//General
	CreatedBy      string     `json:"created_by"`
	CreatedByName  string     `json:"created_by_name"`
	DateCreated    *time.Time `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated"`
	OrganizationID string     `json:"organization_id"`
	UpdatedBy      string     `json:"updated_by"`
	UpdatedByName  string     `json:"updated_by_name"`
	UserID         string     `json:"user_id"`
	UserName       string     `json:"user_name"`
	Users          []string   `json:"users"`
}

type Opens struct {
	IpAddress string     `json:"ip_address"`
	OpenedAt  *time.Time `json:"opened_at"`
	OpenedBy  string     `json:"opened_by"`
	UserAgent string     `json:"user_agent"`
}

type Enveloppe struct {
	Bcc         []string   `json:"bcc"`
	Cc          []string   `json:"cc"`
	Date        *time.Time `json:"date"`
	From        []From     `json:"from"`
	InReplyTo   string     `json:"in_reply_to"`
	IsAutoReply bool       `json:"is_autoreply"`
	MessageID   string     `json:"message_id"`
	ReplyTo     []string   `json:"reply_to"`
	Sender      []From     `json:"sender"`
	Subject     string     `json:"subject"`
	To          []From     `json:"to"`
}

type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type HtmlQuoted struct {
	Expand bool   `json:"expand"`
	Html   string `json:"html"`
}

type TextQuoted struct {
	Expand bool   `json:"expand"`
	Text   string `json:"text"`
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
