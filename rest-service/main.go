package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dev-test/rest-service/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/people/:id", getPersonOnID)
	e.GET("/people", getPeople)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Respond to GET /people with a 200 OK response containing all people in the system, if no optional parameters are given
func getPeople(c echo.Context) error {
	firstName := c.QueryParam("first_name")
	lastName := c.QueryParam("last_name")
	if firstName != "" && lastName != "" {
		return c.JSON(http.StatusOK, getPeopleByFirstLastName(firstName, lastName))
	}

	phoneNumber := strings.TrimSpace(c.QueryParam("phone_number"))
	if phoneNumber != "" {
		return c.JSON(http.StatusOK, getPeopleByPhoneNumber(phoneNumber))
	}

	people := models.AllPeople()

	return c.JSON(http.StatusOK, people)
}

// Respond to GET /people/:id with a 200 OK response containing the requested person
// or a 404 Not Found response if the :id doesn't exist
func getPersonOnID(c echo.Context) error {
	uid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a valid UUID: "+err.Error())
	}

	person, err := models.FindPersonByID(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

// Respond to GET /people?first_name=:first_name&last_name=:last_name with a 200 OK response
// containing the people with that first and last name or an empty array if no people were found
func getPeopleByFirstLastName(firstName, lastName string) []*models.Person {
	people := models.FindPeopleByName(firstName, lastName)

	return people
}

// Respond to GET /people?phone_number=:phone_number with a 200 OK response containing the
// people with that phone number or an empty array if no people were found
func getPeopleByPhoneNumber(phoneNumber string) []*models.Person {
	//the plus sign represent a space in a query string so when we retrieve the value, the phone number is return with an
	//empty string in front of the phone number without the plus sign, so we need to trim the string and concat the "+" back to
	//the string without a space so that it will be able to match correctly when we pass it to the FindPeopleByPhoneNumber function
	people := models.FindPeopleByPhoneNumber(fmt.Sprintf("%s%s", "+", phoneNumber))

	return people
}
