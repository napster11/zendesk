package zendeskService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/napster11/zendesk/zendeskModel"
	"github.com/napster11/zendesk/zendeskUtil"
)

//Test TicketList API
func TestTicketListAPI(t *testing.T) {
	client := &http.Client{}
	URL := "https://singh782.zendesk.com/api/v2/tickets.json?per_page=10"
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", zendeskUtil.BasicAuth)
	resp, err := client.Do(req)
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		t.Log("Request Failed", err)
		t.FailNow()
	}
	fmt.Println("First Test Case Passed")
}

//Test if TicketList is Empty or not
func TestTicketEmpty(t *testing.T) {
	client := &http.Client{}
	URL := "https://singh782.zendesk.com/api/v2/tickets.json?per_page=10"
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", zendeskUtil.BasicAuth)
	resp, err := client.Do(req)
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		t.Log("Request Failed", err)
		t.FailNow()
	} else {
		var data zendeskModel.APIResponse
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			t.Log("Error is parsing the response", err)
			t.FailNow()
		} else {
			if len(data.Tickets) == 0 {
				t.Log("Empty Ticket List", err)
				t.FailNow()
			}
		}
	}
	fmt.Println("Second Test Case Passed")
}
func TestAPIAuthentication(t *testing.T) {
	client := &http.Client{}
	URL := "https://singh782.zendesk.com/api/v2/tickets.json?per_page=10"
	req, _ := http.NewRequest("GET", URL, nil)

	//Uncomment to Pass this test case
	//req.Header.Set("Authorization", zendeskUtil.BasicAuth)
	resp, err := client.Do(req)
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		t.Log("Unauthorized ", err)
	} else {
		t.Log("API passed without Authentication")
		t.FailNow()
	}
	fmt.Println("Third Test Case Passed")
}
