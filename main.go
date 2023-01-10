package main

import (
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

	// start the service
	r.Run(":9090")
}
