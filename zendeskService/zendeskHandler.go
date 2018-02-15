package zendeskService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/napster11/zendesk/zendeskModel"
	"github.com/napster11/zendesk/zendeskUtil"
	"github.com/unrolled/render"
)

//getTicketList is the Handler function to generate the API Response
func getTicketList(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	pageCount := r.URL.Query().Get("per_page") //per_page defines the number of tickets in single request
	pageNumber := r.URL.Query().Get("page")    //page number defines the page sequence
	userName, authToken, _ := r.BasicAuth()

	//If page number query param is missing then default page number will be 1
	if len(pageNumber) == 0 {
		pageNumber = zendeskUtil.DefaultPage
	}
	if len(userName) == 0 || len(authToken) == 0 {
		ErrorStatus := zendeskUtil.AuthenticationError
		res := zendeskUtil.ConstructResponse(http.StatusUnauthorized, ErrorStatus, nil)
		render.JSON(w, http.StatusUnauthorized, res)
		return
	}

	//Group name of the account is same as email ID Like username = singh782 for email: singh782@umn.edu
	groupName := strings.Split(userName, "@")[0]

	//API Endpoint
	URL := "https://" + groupName + zendeskUtil.BaseURL + "?per_page=" + pageCount + "&page=" + pageNumber
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(userName, authToken)

	// Make request
	resp, error := client.Do(req)

	// If Response Code denotes Success then return the success Response otherwise returns Response with Error Message
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data zendeskModel.APIResponse
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			if len(data.Tickets) == 0 {
				res := zendeskUtil.ConstructResponse(http.StatusBadRequest, "No More Data", nil)
				render.JSON(w, http.StatusOK, res)
				return
			}
			render.JSON(w, http.StatusOK, data)
		} else {
			ErrorStatus := zendeskUtil.InvalidResponse
			res := zendeskUtil.ConstructResponse(http.StatusBadRequest, ErrorStatus, nil)
			render.JSON(w, http.StatusNotAcceptable, res)
			return
		}
	} else {
		ErrorStatus := zendeskUtil.InvalidResponse
		if error != nil {
			ErrorStatus = error.Error()
		}
		res := zendeskUtil.ConstructResponse(resp.StatusCode, ErrorStatus, nil)
		render.JSON(w, http.StatusInternalServerError, res)
		return
	}
}
