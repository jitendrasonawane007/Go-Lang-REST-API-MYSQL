package main

import (
	"fmt"
	helper "go-gin/helpers"
	"log"

	"net/http"

	"github.com/gchaincl/dotsql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*******************************************************************/
type sectorMaster struct {
	SectorId   int    `json:"sectorId"`
	SectorName string `json:"sectorName"`
}

/*******************************************************************/
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Contex) {
		var encodedToken = c.Request.Header.Get("Authorization")

		parsedToken, err := ParseToken(encodedToken)
	}
}

/*******************************************************************/

func main() {
	db := helper.DBConnection()
	dot, err := dotsql.LoadFromFile("./database/getMastersQuery.sql")
	defer db.Close()
	result, err := dot.Query(db, "getSectors", 1)
	if err != nil {
		log.Fatal(err)
	}

	var sector sectorMaster
	var sectorArr []sectorMaster
	for result.Next() {
		err := result.Scan(&sector.SectorId, &sector.SectorName)
		if err != nil {
			log.Fatal(err)
		}
		sectorArr = append(sectorArr, sector)
	}
	r := gin.New()

	// this are the middleware to intercept each request
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// group for the  different types of req associated with modules
	v1 := r.Group("/v1")
	{
		v1.GET("api/getSector", func(c *gin.Context) {
			c.JSON(200, sectorArr)
		})
	}
	// r.GET()
	// })
	v2 := r.Group("/v2")
	{
		v2.GET("api/welcome", func(c *gin.Context) {
			firstName := c.DefaultQuery("firstName", "Guest")
			lastName := c.Query("lastName")
			c.String(http.StatusOK, "Hello %s %s", firstName, lastName)

		})

		v2.GET("api/queryMap", func(c *gin.Context) {
			ids := c.QueryMap("id")
			names := c.PostFormMap("names")
			fmt.Printf("Id:%v; names: %v", ids, names)
			c.String(http.StatusOK, "Id %s", ids)

		})
	}
	// authorized group
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/api/login", loginEndPoint)
		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndPoint)

		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}

/*******************************************************************/
