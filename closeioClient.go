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
	SendLead(lead *Lead) error
	GetActivities(leadId string) ([]Activity, error)
	GetAllLeads() ([]Lead, error)
	GetLeads(channel, leadType string) ([]Lead, error)
	DeleteLead(leadID string) error
	GetOpportunities() ([]Opportunity, error)
}

const limit = 100

type HttpCloseIoClient struct {
	apiKey string
}

func NewCloseIoClient(apiKey string) *HttpCloseIoClient {
	return &HttpCloseIoClient{apiKey: apiKey}
}

func (c HttpCloseIoClient) SendLead(lead *Lead) error {
	content, _ := json.Marshal(lead)
	body := bytes.NewBuffer(content)

	_, err := c.getResponse("POST", "lead", nil, body)

	if err != nil {
		return err
	}

	return nil
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

	skip := 0

	leads := []Lead{}
	finish := false
	//Stop when a get a bad request
	for !finish {

		query := map[string]string{"_skip": strconv.Itoa(skip), "_limit": strconv.Itoa(limit)}
		if queryString := convertQueryFields(queryFields); queryString != "" {
			query["query"] = queryString
		}

		body, err := c.getResponse("GET", "lead", query, nil)

		fmt.Println(string(body))
		fetched, err := ParseLeads(body)
		if err != nil {
			return nil, err
		}

		if len(fetched) < limit {
			finish = true
		}

		leads = append(leads, fetched...)
		skip = skip + limit

	}

	return leads, nil
}

func convertQueryFields(queryFields map[string]string) string {
	if queryFields == nil {
		return nil
	}

	var query = ""
	for key, value := range queryFields {
		if query == "" {
			query = fmt.Sprintf(`%s:"%s`, key, value)
		} else {
			query = fmt.Sprintf(`%s AND %s:"%s`, query, key, value)
		}
	}

	return query
}

func (c HttpCloseIoClient) GetActivities(leadId string) ([]Activity, error) {

	query := map[string]string{"lead_id": leadId}
	body, err := c.getResponse("GET", "activity", query, nil)

	if err != nil {
		return nil, err
	}

	return ParseActivities(body)
}

func (c HttpCloseIoClient) SendMailActivity(activity *Activity) error {
	if activity.Type != "Email" {
		return fmt.Errorf("Activity type %s is not supported for creation", activity.Type)
	}

	content, _ := json.Marshal(activity)
	body := bytes.NewBuffer(content)

	_, err := c.getResponse("POST", "activity/email", nil, body)

	if err != nil {
		return err
	}

	return nil
}

func ParseActivities(content []byte) ([]Activity, error) {

	dataActivities := struct {
		Activities []Activity `json:"data"`
	}{}

	err := json.Unmarshal(content, &dataActivities)
	if err != nil {
		return nil, fmt.Errorf("Error while deserializing json %s", err.Error())
	}

	return dataActivities.Activities, nil
}

func ParseLeads(content []byte) ([]Lead, error) {
	dataLeads := struct {
		Leads []Lead `json:"data"`
	}{}

	err := json.Unmarshal(content, &dataLeads)
	if err != nil {
		return nil, fmt.Errorf("Error while deserializing json %s \n %s \n", err.Error(), string(content))
	}

	return dataLeads.Leads, nil
}

func (c HttpCloseIoClient) GetOpportunities() ([]Opportunity, error) {

	skip := 0

	opportunities := []Opportunity{}
	finish := false
	//Stop when a get a bad request
	for !finish {

		query := map[string]string{"_skip": strconv.Itoa(skip), "_limit": strconv.Itoa(limit)}
		body, err := c.getResponse("GET", "opportunity", query, nil)

		if err != nil {
			return nil, err
		}

		opportunityResponse := struct {
			Opportunities []Opportunity `json:"data"`
		}{}

		err = json.Unmarshal(body, &opportunityResponse)
		if err != nil {
			return nil, fmt.Errorf("Error while deserializing json %s", err.Error())
		}

		if len(opportunityResponse.Opportunities) < limit {
			finish = true
		}
		opportunities = append(opportunities, opportunityResponse.Opportunities...)
		skip = skip + limit

	}

	return opportunities, nil

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
