package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1/ method use map
// data := map[string]interface{}{
// 	"bookName":    "harry potter and the prisoner of azkaban",
// 	"author":      "J. K. Rowling",
// 	"publishDate": 1999,
// }

// data := gin.H{
// 	"bookName":    "harry potter and the prisoner of azkaban",
// 	"author":      "J. K. Rowling",
// 	"publishDate": 1999,
// }

// 2/ method use struct
type msg struct {
	BookName    string `json:"bookName"`
	PublishDate int    `json:"publishDate"`
	Author      string `json:"author"`
}

// ** gin ShouldBind **

type UserInfo struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required"`
}

func main() {
	r := gin.Default()

	// *** return json generated from struct ***
	r.GET("/book", func(c *gin.Context) {
		// 2/ instanciate a struct msg
		data := msg{
			BookName:    "Harry Potter and the prisoner of azkaban",
			PublishDate: 1999,
			Author:      "J.K.Rowling",
		}
		c.JSON(http.StatusOK, data)
	})

	// *** retrieve query parameters ***

	// GET request, after ? is the query string key and value
	// format: key=value, if there are several key value pairs it is concatenated with &
	// example: /web/query=peach&age=21
	r.GET("/web", func(c *gin.Context) {
		// 1/ c.Query
		// name := c.Query("query")
		// age := c.Query("age")

		// 2/ c.DefaultQuery(key, defaultValue)
		name := c.DefaultQuery("query", "somebody") // if can't retrive parameter with "query" key, if would use the default value "somebody"
		age := c.DefaultQuery("age", "18")

		// 3/ c.GetQuery(key)(string, bool) // returns false if the key is not retrieved from query parameter
		// name, ok := c.GetQuery("query")
		// if !ok {
		// 	name = "somebody"
		// }

		// if request look like this => localhost:9090/web?query=mario
		// response look like this => {"name": "mario", "age": "18"}
		// if request look like this => localhost:9090/web?query=mario&age=56
		// response look like this => {"name": "mario", "age": "56"}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// *** retrieve path (URI) parameter ***

	// the data returned by path parameter is always string type
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		// retrieve path (URI) parameter
		year := c.Param("year")   // type string
		month := c.Param("month") // type string

		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	// ** gin ShouldBind **

	// ** bind query parameter, struct shuold have form tag, for example, form:"name"
	r.POST("/user", func(c *gin.Context) {
		var u UserInfo // initialize a variable with UserInfo type
		if err := c.ShouldBind(&u); err == nil {
			fmt.Printf("user info: %#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"name": u.Name,
				"age":  u.Age,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	// ** Gin redirect **
	r.GET("/youtubeRedirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/")
	})

	// start the service
	r.Run(":9090")
}
