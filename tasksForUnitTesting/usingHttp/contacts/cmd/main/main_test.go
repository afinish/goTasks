package main_test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/config"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/controllers"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/models"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/routes"
	"github.com/stretchr/testify/assert"
)

// var router *gin.Engine
// var e *gin.Engine

func registerNewRouter() *gin.Engine {
	r := gin.Default()
	v0 := r.Group("/")
	{
		// v0.GET("contacts", controllers.GetContacts)
		v0.GET("contact", controllers.GetContact)
		// v0.POST("contact", controllers.CreateContact)
		// v0.PUT("contact", controllers.UpdateContact)
		// v0.DELETE("contact", controllers.DeleteContact)
	}
	return r
}

func TestMain(m *testing.M) {
	var err error
	config.DB, err = sql.Open("postgres", config.DbURL())
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()
	r := routes.RegisterNewRouter()
	// e = r
	r.Run()

	code := m.Run()

	os.Exit(code)
}

func TestGetContact(t *testing.T) {
	config.DB.Exec(`
		DELETE FROM contacts;
		INSERT INTO contacts (first_name, last_name, phone, email, position) VALUES ('TFirstName', 'TLastName', 'T999119991199', 'T@mail.com', 'TPosition');
		ALTER SEQUENCE seq RESTART;
		UPDATE contacts SET id = DEFAULT;
	`)

	req, _ := http.NewRequest("GET", "/contact/1", nil)
	// response := executeRequest(req)
	res := httptest.NewRecorder()
	registerNewRouter().ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "OK response is expected")
	expected := models.Contact{
		ID: 1,
		FirstName: "TFirstName", 
		LastName: "TLastName", 
		Phone: "T999119991199", 
		Email: "T@mail.com", 
		Position: "TPosition",
	}
	actual := &models.Contact{}
	parseBody(res, actual)
	assert.Equal(t, expected, actual)
}

func parseBody(r *httptest.ResponseRecorder, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}