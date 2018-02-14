package zendeskService

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketListAPI(t *testing.T) {
	res := httptest.NewRecorder()
	a := assert.New(t)

	body := `{
    "username": "sdayanand@tekion.com",
    "password": "Tekion123"}`
	_, err := http.NewRequest("GET", "/tloginservice/login")

	a.NoError(err)
	a.Equal(res.Code, 400)
	a.Equal(res.Body.String(), "{\"meta\":{\"code\":400,\"msg\":\"Something Went Wrong.\"}}")
}
