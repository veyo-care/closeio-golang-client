package closeio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type CloseIoClient interface {
	SendLead(lead *Lead) (*Lead, error)
	GetLead(leadID string) (*Lead, error)
	GetLeads(queryFields map[string][]string) ([]Lead, error)
	GetLeadsWithRawQuery(queryString string) ([]Lead, error)
	GetAllLeads() ([]Lead, error)
	DeleteLead(leadID string) error
	UpdateLead(lead *Lead) (*Lead, error)
	UpdateLeadStatus(leadId, statusId string) (*Lead, error)

	GetTasks(model Task) ([]Task, error)

	SendActivity(activity *Activity) error
	GetAllActivities() ([]Activity, error)
	GetActivities(leadId string) ([]Activity, error)
	UpdateActivity(activity *Activity) error

	SendOpportunity(opportunity *Opportunity) error
	GetOpportunities() ([]Opportunity, error)

	SendTask(task *Task) error
	DeleteTask(taskID string) error
	UpdateTask(task *Task) error

	CreateContact(contact *Contact) (*Contact, error)
	UpdateContact(contact *Contact) (*Contact, error)
	GetContact(contactID string) (*Contact, error)
	DeleteContact(contactID string) error

	UpdateAddress(address Address, leadID string) error

	GetAllUsers() ([]User, error)

	GetLeadStatuses() ([]Status, error)
}

const limit = 100 //maximum set by closeio

type HttpCloseIoClient struct {
	apiKey string
}

func NewCloseIoClient(apiKey string) *HttpCloseIoClient {
	return &HttpCloseIoClient{apiKey: apiKey}
}

func (c HttpCloseIoClient) SendLead(lead *Lead) (*Lead, error) {
	return c.actionOnLead("POST", "lead", lead)
}

func (c HttpCloseIoClient) GetLead(leadID string) (*Lead, error) {

	responseBody, err := c.getResponse("GET", fmt.Sprintf("lead/%s", leadID), nil, nil)

	if err != nil {
		return nil, err
	}

	lead, err := JSONToLead(responseBody)
	if err != nil {
		return nil, fmt.Errorf("error deserializing lead - %s", err.Error())
	}
	return lead, nil
}

func (c HttpCloseIoClient) DeleteLead(leadID string) error {

	_, err := c.getResponse("DELETE", fmt.Sprintf("lead/%s", leadID), nil, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c HttpCloseIoClient) UpdateLeadStatus(leadID, statusId string) (*Lead, error) {
	message := struct {
		StatusID string `json:"status_id"`
	}{statusId}

	content, _ := json.Marshal(message)
	return c.actionOnObject("PUT", fmt.Sprintf("lead/%s", leadID), content)
}

func (c HttpCloseIoClient) UpdateLead(lead *Lead) (*Lead, error) {
	return c.actionOnLead("PUT", fmt.Sprintf("lead/%s", lead.ID), lead)
}

func (c HttpCloseIoClient) actionOnLead(method string, route string, lead *Lead) (*Lead, error) {
	if lead == nil {
		return nil, nil
	}

	content, err := LeadToJSON(*lead)

	if err != nil {
		return nil, err
	}

	return c.actionOnObject(method, route, content)
}

func (c HttpCloseIoClient) actionOnObject(method string, route string, content []byte) (*Lead, error) {
	body := bytes.NewBuffer(content)

	responseBody, err := c.getResponse(method, route, nil, body)

	if err != nil {
		return nil, err
	}

	lead, err := JSONToLead(responseBody)

	if err != nil {
		return nil, err
	}

	return lead, nil
}

func (c HttpCloseIoClient) GetAllLeads() ([]Lead, error) {
	return c.GetLeads(nil)
}

func (c HttpCloseIoClient) GetLeads(queryFields map[string][]string) ([]Lead, error) {

	query := make(map[string]string)

	if queryString := convertQueryFields(queryFields); queryString != "" {
		query["query"] = queryString
	}

	elements, err := c.getElements("lead", query)

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

func (c HttpCloseIoClient) GetTasks(task Task) ([]Task, error) {
	query := make(map[string]string)

	if task.AssignedTo != "" {
		query["assigned_to"] = task.AssignedTo
	}

	if task.IsComplete != nil {
		query["is_complete"] = strconv.FormatBool(*task.IsComplete)
	}

	if task.LeadID != "" {
		query["lead_id"] = task.LeadID
	}

	if task.View != "" {
		query["view"] = task.View
	}

	elements, err := c.getElements("task", query)

	if err != nil {
		return nil, err
	}

	tasks := make([]Task, len(elements), len(elements))

	for i, element := range elements {
		var task Task

		err = json.Unmarshal([]byte(element), &task)

		if err != nil {
			return nil, err
		}

		tasks[i] = task

	}

	return tasks, nil
}

func (c HttpCloseIoClient) GetLeadsWithRawQuery(queryString string) ([]Lead, error) {
	query := make(map[string]string)

	query["query"] = queryString

	elements, err := c.getElements("lead", query)

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

func convertQueryFields(queryFields map[string][]string) string {
	if queryFields == nil {
		return ""
	}

	andValues := make([]string, len(queryFields), len(queryFields))

	i := 0
	for key, values := range queryFields {

		orValues := make([]string, len(values), len(values))

		for k, value := range values {
			orValues[k] = fmt.Sprintf(`%s:"%s"`, key, value)
		}

		andValues[i] = fmt.Sprintf("(%s)", strings.Join(orValues, " OR "))
		i++
	}

	return strings.Join(andValues, " AND ")
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
		return nil, fmt.Errorf("Could not retrieve activities for query %+v %s", queryFields, err.Error())
	}

	activities := make([]Activity, len(elements), len(elements))

	for i, element := range elements {
		var activitiy Activity

		err = json.Unmarshal([]byte(element), &activitiy)

		if err != nil {
			return nil, fmt.Errorf("Could not deserialize %s %s", string(element), err.Error())
		}

		activities[i] = activitiy

	}

	return activities, nil
}

func (c HttpCloseIoClient) SendActivity(activity *Activity) error {
	path, err := getActivityPath(activity)
	if err != nil {
		return err
	}
	content, _ := json.Marshal(activity)
	body := bytes.NewBuffer(content)

	_, err = c.getResponse("POST", path, nil, body)

	if err != nil {
		return err
	}
	return nil
}

func (c HttpCloseIoClient) UpdateActivity(activity *Activity) error {
	path, err := getActivityPath(activity)
	if err != nil {
		return err
	}
	completePath := fmt.Sprintf("%s/%s", path, activity.ID)
	content, _ := json.Marshal(activity)
	body := bytes.NewBuffer(content)

	_, err = c.getResponse("PUT", completePath, nil, body)

	if err != nil {
		return err
	}
	return nil
}

func getActivityPath(activity *Activity) (string, error) {
	switch activity.Type {
	case "Email":
		return "activity/email", nil
	case "Note":
		return "activity/note", nil
	case "Call":
		return "activity/call", nil
	default:
		return "", fmt.Errorf("Activity type %s is not supported for creation", activity.Type)
	}
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

const n = 5

func (c HttpCloseIoClient) getElements(route string, query map[string]string) ([]json.RawMessage, error) {

	if query == nil {
		query = make(map[string]string)
	}

	results := make(chan resp)
	jobs := make(chan req)
	inWork := 0

	for i := 0; i < n; i++ {
		go c.httpWorker(jobs, results)
		sendJob(jobs, i, route, query)
		inWork++
	}

	blobs := []json.RawMessage{}

	id := n
	lastRound := false
	for result := range results {
		inWork--
		if result.err != nil {
			return nil, result.err
		}
		blobs = append(blobs, result.blobs...)
		if result.hasMore && !lastRound {
			sendJob(jobs, id, route, query)
			inWork++
			id++
		} else {
			lastRound = true
		}
		if inWork == 0 {
			close(results)
			close(jobs)
		}
	}
	return blobs, nil
}

func copyQuery(query map[string]string) map[string]string {
	copy := make(map[string]string)

	for key, value := range query {
		copy[key] = value
	}

	return copy
}

func sendJob(jobs chan req, id int, route string, query map[string]string) {
	copy := copyQuery(query)
	skip := id * limit
	copy["_limit"] = strconv.Itoa(limit)
	copy["_skip"] = strconv.Itoa(skip)
	jobs <- req{id, route, copy}
}

type req struct {
	id    int
	route string
	query map[string]string
}

type resp struct {
	id      int
	blobs   []json.RawMessage
	hasMore bool
	err     error
}

func (c HttpCloseIoClient) httpWorker(jobs chan req, results chan resp) {
	for j := range jobs {
		body, err := c.getResponse("GET", j.route, j.query, nil)
		if err != nil {
			results <- resp{err: err}
		}
		response := struct {
			Blobs   []json.RawMessage `json:"data"`
			HasMore bool              `json:"has_more"`
		}{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			results <- resp{err: fmt.Errorf("Error while deserializing json %s", err.Error())}
		}
		results <- resp{j.id, response.Blobs, response.HasMore, err}
	}
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

func (c HttpCloseIoClient) UpdateTask(task *Task) error {
	content, _ := json.Marshal(task)

	body := bytes.NewBuffer(content)

	_, err := c.getResponse("PUT", "task/"+task.ID, nil, body)
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

func (c HttpCloseIoClient) UpdateAddress(address Address, leadID string) error {
	lead := Lead{
		Addresses: []Address{address},
	}
	content, _ := json.Marshal(lead)
	body := bytes.NewBuffer(content)
	_, err := c.getResponse("PUT", fmt.Sprintf("lead/%s", leadID), nil, body)
	return err
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
