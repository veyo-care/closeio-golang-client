package closeio

import "time"

type Contact struct {
	CreatedBy   string     `json:"created_by,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Emails      []Email    `json:"emails,omitempty"`
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Phones      []Phone    `json:"phones,omitempty"`
	Title       string     `json:"title,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
}

type Email struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Phone struct {
	Phone          string `json:"phone,omitempty"`
	PhoneFormatted string `json:"phone_formatted,omitempty"`
	Type           string `json:"type,omitempty"`
}

type Activity struct {
	Type string `json:"_type,omitempty"`
	ID   string `json:"id,omitempty"`

	//Email

	Attachments         []Attachment `json:"attachments,omitempty"`
	Bcc                 []string     `json:"bcc,omitempty"`
	BodyHtml            string       `json:"body_html,omitempty"`
	BodyHtmlQuoted      []HtmlQuoted `json:"body_html_quoted,omitempty"`
	BodyPreview         string       `json:"body_preview,omitempty"`
	BodyText            string       `json:"body_text,omitempty"`
	BodyTextQuoted      []TextQuoted `json:"body_text_quoted,omitempty"`
	Cc                  []string     `json:"cc,omitempty"`
	ContactID           string       `json:"contact_id,omitempty"`
	DateScheduled       *time.Time   `json:"date_scheduled,omitempty"`
	DateSent            *time.Time   `json:"date_sent,omitempty"`
	EmailAccountID      string       `json:"email_account_id,omitempty"`
	Enveloppe           Enveloppe    `json:"enveloppe,omitempty"`
	InReplyToID         string       `json:"in_reply_to_id,omitempty"`
	LeadID              string       `json:"lead_id,omitempty"`
	MessageIDs          []string     `json:"message_ids,omitempty"`
	NeedSmtpCredentials bool         `json:"need_smtp_credentials,omitempty"`
	Opens               []Opens      `json:"opens,omitempty"`
	OpensSummary        string       `json:"opens_summary,omitempty"`
	References          []string     `json:"references,omitempty"`
	SendAttempts        []string     `json:"send_attempts,omitempty"`
	Sender              string       `json:"sender,omitempty"`
	Status              string       `json:"status,omitempty"`
	Subject             string       `json:"subject,omitempty"`
	TemplateID          string       `json:"template_id,omitempty"`
	TemplateName        string       `json:"template_name,omitempty"`
	ThreadID            string       `json:"thread_id,omitempty"`
	To                  []string     `json:"to,omitempty"`

	//LeadStatusChange

	NewStatusID string `json:"new_status_id,omitempty"`
	NewStatus   string `json:"new_status_label,omitempty"`
	OldStatusID string `json:"old_status_id,omitempty"`
	OldStatus   string `json:"old_status_label,omitempty"`

	//OpportunityStatusChange

	NewStatusType             string     `json:"new_status_type,omitempty"`
	OldStatusType             string     `json:"old_status_type,omitempty"`
	OpportunityConfidence     int        `json:"opportunity_confidence,omitempty"`
	OpportunityDateWon        *time.Time `json:"opportunity_date_won,omitempty"`
	OpportunityID             string     `json:"opportunity_id,omitempty"`
	OpportunityValue          float64    `json:"opportunity_value,omitempty"`
	OpportunityValueCurrency  string     `json:"opportunity_value_currency,omitempty"`
	OpportunityValueFormatted string     `json:"opportunity_value_formatted,omitempty"`
	OpportunityValuePeriod    string     `json:"opportunity_value_period,omitempty"`

	//Call

	Direction         string `json:"direction,omitempty"`
	CallDuration      int64  `json:"duration,omitempty"`
	LocalPhone        string `json:"local_phone,omitempty"`
	Note              string `json:"note,omitempty"`
	Phone             string `json:"phone,omitempty"`
	RecordingUrl      string `json:"recording_url,omitempty"`
	RemotePhone       string `json:"remote_phone,omitempty"`
	Source            string `json:"source,omitempty"`
	TransferredFrom   string `json:"transferred_from,omitempty"`
	TransferredTo     string `json:"transferred_to,omitempty"`
	VoiceMailDuration int64  `json:"voicemail_duration,omitempty"`
	VoiceMailUrl      string `json:"voicemail_url,omitempty"`

	//TaskCompleted

	TaskAssignedTo     string `json:"task_assigned_to,omitempty"`
	TaskAssignedToName string `json:"task_assigned_to_name,omitempty"`
	TaskID             string `json:"task_id,omitempty"`
	TaskText           string `json:"task_text,omitempty"`

	//Created

	ImportID string `json:"import_id,omitempty"`

	//General
	CreatedBy      string     `json:"created_by,omitempty"`
	CreatedByName  string     `json:"created_by_name,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	OrganizationID string     `json:"organization_id,omitempty"`
	UpdatedBy      string     `json:"updated_by,omitempty"`
	UpdatedByName  string     `json:"updated_by_name,omitempty"`
	UserID         string     `json:"user_id,omitempty"`
	UserName       string     `json:"user_name,omitempty"`
	Users          []string   `json:"users,omitempty"`
}

type Opens struct {
	IpAddress string     `json:"ip_address,omitempty"`
	OpenedAt  *time.Time `json:"opened_at,omitempty"`
	OpenedBy  string     `json:"opened_by,omitempty"`
	UserAgent string     `json:"user_agent,omitempty"`
}

type Enveloppe struct {
	Bcc         []string   `json:"bcc,omitempty"`
	Cc          []string   `json:"cc,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	From        []From     `json:"from,omitempty"`
	InReplyTo   string     `json:"in_reply_to,omitempty"`
	IsAutoReply bool       `json:"is_autoreply,omitempty"`
	MessageID   string     `json:"message_id,omitempty"`
	ReplyTo     []string   `json:"reply_to,omitempty"`
	Sender      []From     `json:"sender,omitempty"`
	Subject     string     `json:"subject,omitempty"`
	To          []From     `json:"to,omitempty"`
}

type From struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type HtmlQuoted struct {
	Expand bool   `json:"expand,omitempty"`
	Html   string `json:"html,omitempty"`
}

type TextQuoted struct {
	Expand bool   `json:"expand,omitempty"`
	Text   string `json:"text,omitempty"`
}

type Attachment struct {
	ContentID   string `json:"content_id,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	FileName    string `json:"filename,omitempty"`
	Size        int    `json:"size,omitempty"`
	URL         string `json:"url,omitempty"`
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
	Tasks          []Task            `json:"tasks,omitempty"`
	UpdatedBy      string            `json:"updated_by,omitempty"`
	UpdatedByName  string            `json:"updated_by_name,omitempty"`
	URL            string            `json:"url,omitempty"`
	CreatedBy      string            `json:"created_by,omitempty"`
	CreatedName    string            `json:"created_by_name,omitempty"`

	Activities []Activity `json:"activities,omitempty`
}

type Task struct {
	Type           string     `json:"_type"`
	AssignedTo     string     `json:"assigned_to,omitempty"`
	AssignedToName string     `json:"assigned_to_name,omitempty"`
	ContactID      string     `json:"contact_id,omitempty"`
	ContactName    string     `json:"contact_name,omitempty"`
	CreatedBy      string     `json:"created_by,omitempty"`
	CreatedName    string     `json:"created_by_name,omitempty"`
	Date           string     `json:"date,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	ID             string     `json:"id,omitempty"`
	IsComplete     bool       `json:"is_complete,omitempty"`
	IsDateLess     bool       `json:"is_dateless,omitempty"`
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
	Confidence     int        `json:"confidence,omitempty"`
	ContactID      string     `json:"contact_id,omitempty"`
	ContactName    string     `json:"contact_name,omitempty"`
	CreationDate   *time.Time `json:"date_created,omitempty"`
	CreatedBy      string     `json:"created_by,omitempty"`
	CreatedName    string     `json:"created_by_name,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DateLost       *time.Time `json:"date_lost,omitempty"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	DateWon        string     `json:"date_won,omitempty"` //not classic format
	ID             string     `json:"id,omitempty"`
	LeadID         string     `json:"lead_id,omitempty"`
	LeadName       string     `json:"lead_name,omitempty"`
	Note           string     `json:"note,omitempty"`
	OrganizationID string     `json:"organization_id,omitempty"`
	StatusID       string     `json:"status_id,omitempty"`
	StatusLabel    string     `json:"status_label,omitempty"`
	StatusType     string     `json:"status_type,omitempty"`
	UpdatedBy      string     `json:"updated_by,omitempty"`
	UpdatedByName  string     `json:"updated_by_name,omitempty"`
	UserID         string     `json:"user_id,omitempty"`
	UserName       string     `json:"user_name,omitempty"`
	Value          float64    `json:"value,omitempty"`
	ValueCurrency  string     `json:"value_currency,omitempty"`
	ValueFormatted string     `json:"value_formatted,omitempty"`
	ValuePeriod    string     `json:"value_period,omitempty"`
}

type Status struct {
	ID    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
}
