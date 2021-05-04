package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

//const url = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/coWin/getSlot/:age/:districtID/:date", getSlot)
	e.Logger.Fatal(e.Start(":1323"))

}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	// id := c.Param("id")
	// return c.String(http.StatusOK, id)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	//log.Printf(sb)
	return c.String(http.StatusOK, sb)
}

func getSlot(c echo.Context) error {

	a := c.Param("age")
	dst := c.Param("districtID")
	dt := c.Param("date")
	URL := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=" + dst + "&date=" + dt
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(a)
	return c.String(http.StatusOK, sb)
}
