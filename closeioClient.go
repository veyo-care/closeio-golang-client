package closeio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CloseIoClient interface {
	GetLead(leadID string) (*Lead, error)
	SendActivity(activity *Activity) error
	SendLead(lead *Lead) (*Lead, error)
	SendTask(task *Task) error
	GetAllActivities() ([]Activity, error)
	GetActivities(leadId string) ([]Activity, error)
	GetAllLeads() ([]Lead, error)
	GetLeads(queryFields map[string]string) ([]Lead, error)
	DeleteLead(leadID string) error
	GetOpportunities() ([]Opportunity, error)

	CreateContact(contact *Contact) (*Contact, error)
	UpdateContact(contact *Contact) (*Contact, error)
	GetContact(contactID string) (*Contact, error)
	DeleteContact(contactID string) error
	DeleteTask(taskID string) error
	GetAllUsers() ([]User, error)
}

const limit = 100

type HttpCloseIoClient struct {
	apiKey string
}

func NewCloseIoClient(apiKey string) *HttpCloseIoClient {
	return &HttpCloseIoClient{apiKey: apiKey}
}

func (c HttpCloseIoClient) SendLead(lead *Lead) (*Lead, error) {
	content, _ := json.Marshal(lead)
	body := bytes.NewBuffer(content)

	responseBody, err := c.getResponse("POST", "lead", nil, body)

	if err != nil {
		return nil, err
	}
	var responseLead Lead
	err = json.Unmarshal(responseBody, &responseLead)
	if err != nil {
		return nil, err
	}
	return &responseLead, nil
}

func (c HttpCloseIoClient) GetLead(leadID string) (*Lead, error) {

	responseBody, err := c.getResponse("GET", fmt.Sprintf("lead/%s", leadID), nil, nil)

	if err != nil {
		return nil, err
	}

	var responseLead Lead
	err = json.Unmarshal(responseBody, &responseLead)
	if err != nil {
		return nil, err
	}
	return &responseLead, nil
}

func (c HttpCloseIoClient) DeleteLead(leadID string) error {

	_, err := c.getResponse("DELETE", fmt.Sprintf("lead/%s", leadID), nil, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c HttpCloseIoClient) GetAllLeads() ([]Lead, error) {
	return c.GetLeads(nil)
}

func (c HttpCloseIoClient) GetLeads(queryFields map[string]string) ([]Lead, error) {

	elements, err := c.getElements("lead", queryFields)

	if err != nil {
		return nil, err
	}

	leads := make([]Lead, len(elements), len(elements))

	for i, element := range elements {
		var lead Lead

		err = json.Unmarshal([]byte(element), &lead)

		if err != nil {
			return nil, err
		}

		leads[i] = lead

	}

	return leads, nil
}

func (c HttpCloseIoClient) CreateContact(contact *Contact) (*Contact, error) {
	content, _ := json.Marshal(contact)
	body := bytes.NewBuffer(content)

	responseBody, err := c.getResponse("POST", "contact", nil, body)

	if err != nil {
		return nil, err
	}
	var responseContact Contact
	err = json.Unmarshal(responseBody, &responseContact)
	if err != nil {
		return nil, err
	}
	return &responseContact, nil
}

func (c HttpCloseIoClient) UpdateContact(contact *Contact) (*Contact, error) {
	content, _ := json.Marshal(contact)
	body := bytes.NewBuffer(content)

	responseBody, err := c.getResponse("PUT", fmt.Sprintf("contact/%s", contact.ID), nil, body)

	if err != nil {
		return nil, err
	}
	var responseContact Contact
	err = json.Unmarshal(responseBody, &responseContact)
	if err != nil {
		return nil, err
	}
	return &responseContact, nil
}

func (c HttpCloseIoClient) GetContact(contactID string) (*Contact, error) {

	responseBody, err := c.getResponse("GET", fmt.Sprintf("contact/%s", contactID), nil, nil)

	if err != nil {
		return nil, err
	}

	var responseContact Contact
	err = json.Unmarshal(responseBody, &responseContact)
	if err != nil {
		return nil, err
	}
	return &responseContact, nil
}

func (c HttpCloseIoClient) DeleteContact(contactID string) error {

	_, err := c.getResponse("DELETE", fmt.Sprintf("contact/%s", contactID), nil, nil)

	if err != nil {
		return err
	}

	return nil
}

func convertQueryFields(queryFields map[string]string) string {
	if queryFields == nil {
		return ""
	}

	var query = ""
	for key, value := range queryFields {
		if query == "" {
			query = fmt.Sprintf(`%s:"%s"`, key, value)
		} else {
			query = fmt.Sprintf(`%s AND %s:"%s"`, query, key, value)
		}
	}

	return query
}

func (c HttpCloseIoClient) GetAllActivities() ([]Activity, error) {
	return c.getActivities(nil)
}

func (c HttpCloseIoClient) GetActivities(leadId string) ([]Activity, error) {

	query := map[string]string{"lead_id": leadId}
	return c.getActivities(query)
}

func (c HttpCloseIoClient) getActivities(queryFields map[string]string) ([]Activity, error) {

	elements, err := c.getElements("activity", queryFields)

	if err != nil {
		return nil, err
	}

	activities := make([]Activity, len(elements), len(elements))

	for i, element := range elements {
		var activitiy Activity

		err = json.Unmarshal([]byte(element), &activitiy)

		if err != nil {
			return nil, err
		}

		activities[i] = activitiy

	}

	return activities, nil
}

func (c HttpCloseIoClient) SendActivity(activity *Activity) error {
	var path string
	switch activity.Type {
	case "Email":
		path = "activity/email"
	case "Note":
		path = "activity/note"
	case "Call":
		path = "activity/call"
	default:
		return fmt.Errorf("Activity type %s is not supported for creation", activity.Type)
	}
	content, _ := json.Marshal(activity)
	body := bytes.NewBuffer(content)

	_, err := c.getResponse("POST", path, nil, body)

	if err != nil {
		return err
	}
	return nil
}

func (c HttpCloseIoClient) GetOpportunities() ([]Opportunity, error) {

	elements, err := c.getElements("opportunity", nil)

	if err != nil {
		return nil, err
	}

	opportunities := make([]Opportunity, len(elements), len(elements))

	for i, element := range elements {
		var opportunity Opportunity

		err = json.Unmarshal([]byte(element), &opportunity)

		if err != nil {
			return nil, err
		}

		opportunities[i] = opportunity

	}

	return opportunities, nil
}

func (c HttpCloseIoClient) getElements(route string, queryFields map[string]string) ([]json.RawMessage, error) {

	skip := 0

	blobs := []json.RawMessage{}

	finish := false
	//Stop when a get a bad request
	for !finish {

		query := map[string]string{"_skip": strconv.Itoa(skip), "_limit": strconv.Itoa(limit)}

		if queryString := convertQueryFields(queryFields); queryString != "" {
			query["query"] = queryString
		}

		body, err := c.getResponse("GET", route, query, nil)

		if err != nil {
			return nil, err
		}

		response := struct {
			Blobs   []json.RawMessage `json:"data"`
			HasMore bool              `json:"has_more"`
		}{}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("Error while deserializing json %s", err.Error())
		}

		finish = !response.HasMore
		blobs = append(blobs, response.Blobs...)
		skip = skip + limit

	}

	return blobs, nil
}

func (c HttpCloseIoClient) SendTask(task *Task) error {
	content, _ := json.Marshal(task)

	body := bytes.NewBuffer(content)

	_, err := c.getResponse("POST", "task", nil, body)
	if err != nil {
		return err
	}
	return nil
}

func (c HttpCloseIoClient) DeleteTask(taskID string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("task/%s", taskID), nil, nil)

	if err != nil {
		return err
	}
	return nil
}

func (c HttpCloseIoClient) GetLeadStatuses() ([]Status, error) {
	responseBody, err := c.getResponse("GET", "status/lead", nil, nil)
	if err != nil {
		return make([]Status, 0), err
	}
	dataStatuses := struct {
		Statuses []Status `json:"data"`
	}{}
	err = json.Unmarshal(responseBody, &dataStatuses)
	if err != nil {
		return make([]Status, 0), err
	}
	return dataStatuses.Statuses, nil

}

func (c HttpCloseIoClient) GetOpportunityStatuses() ([]Status, error) {
	responseBody, err := c.getResponse("GET", "status/opportunity", nil, nil)
	if err != nil {
		return make([]Status, 0), err
	}
	dataStatuses := struct {
		Statuses []Status `json:"data"`
	}{}
	err = json.Unmarshal(responseBody, &dataStatuses)
	if err != nil {
		return make([]Status, 0), err
	}
	return dataStatuses.Statuses, nil

}

func (c HttpCloseIoClient) SendOpportunity(opportunity *Opportunity) error {
	content, _ := json.Marshal(opportunity)

	body := bytes.NewBuffer(content)

	_, err := c.getResponse("POST", "opportunity", nil, body)
	if err != nil {
		return err
	}
	return nil
}

func (c HttpCloseIoClient) GetAllUsers() ([]User, error) {
	responseBody, err := c.getResponse("GET", "user", nil, nil)
	if err != nil {
		return make([]User, 0), err
	}
	dataUsers := struct {
		Users []User `json:"data"`
	}{}
	err = json.Unmarshal(responseBody, &dataUsers)
	if err != nil {
		return make([]User, 0), err
	}
	return dataUsers.Users, nil
}

func (c HttpCloseIoClient) getResponse(method, route string, query map[string]string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://app.close.io/api/v1/%s/", route), body)

	if err != nil {
		return nil, fmt.Errorf("Error while creating http request %s", err.Error())
	}

	if query != nil && len(query) > 0 {
		values := req.URL.Query()

		for key, value := range query {
			values.Add(key, value)
		}

		req.URL.RawQuery = values.Encode()
	}

	c.fillRequest(req)

	return getResponse(req)
}

func getResponse(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("Failed to get response - err %s, request %+v", err.Error(), request)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Could not read response body - err %s, request %+v", err.Error(), request)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Got status %d from request %+v, %s", resp.StatusCode, request, string(raw))
	}

	return raw, nil
}

func (c HttpCloseIoClient) fillRequest(request *http.Request) {
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.apiKey, "")
}
