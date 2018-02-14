package zendeskService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/napster11/zendesk/zendeskModel"
	"github.com/napster11/zendesk/zendeskUtil"
	"github.com/unrolled/render"
)

//getQuotes is the Handler function to generate the API Response
func getTicketList(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	pageCount := r.URL.Query().Get("per_page")
	pageNumber := r.URL.Query().Get("page")
	userName, authToken, _ := r.BasicAuth()
	fmt.Println("Auth" + r.Header.Get("Authorization"))
	if len(pageNumber) == 0 {
		pageNumber = zendeskUtil.DefaultPage
	}
	if len(userName) == 0 || len(authToken) == 0 {
		ErrorStatus := zendeskUtil.AuthenticationError
		res := zendeskUtil.ConstructResponse(http.StatusUnauthorized, ErrorStatus, nil)
		render.JSON(w, http.StatusUnauthorized, res)
		return
	}
	groupName := strings.Split(userName, "@")[0]

	//Base URL of API Endpoint
	URL := "https://" + groupName + zendeskUtil.BaseURL + "?per_page=" + pageCount + "&page=" + pageNumber
	fmt.Println(URL)
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
			render.JSON(w, http.StatusOK, data)
		} else {
			res := zendeskUtil.ConstructResponse(http.StatusNotAcceptable, err.Error(), nil)
			render.JSON(w, http.StatusNotAcceptable, res)
			return
		}
	} else {
		ErrorStatus := zendeskUtil.InvalidRequest
		if error != nil {
			ErrorStatus = error.Error()
		}
		res := zendeskUtil.ConstructResponse(resp.StatusCode, ErrorStatus, nil)
		render.JSON(w, http.StatusInternalServerError, res)
		return
	}
}
