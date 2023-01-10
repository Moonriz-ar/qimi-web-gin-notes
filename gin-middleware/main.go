package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ** middleware **

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")
	if !ok {
		name = "anonymous user"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// define a middleware m1: 統計請求處理函數的耗時
func m1(c *gin.Context) {
	fmt.Println("m1 in...")

	start := time.Now()
	c.Next() // 調用後續的處理函數
	// c.Abort() // 阻止調用後續的處理函數
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "mario")
	c.Next()
	// c.Abort() // stops the functions after this middleware, remaining of this function is executed
	// return // steps out of this middleware，remaining of this function is not executed
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// connect to database
	// or other setup tasks
	return func(c *gin.Context) {
		if doCheck {
			// specific logic to determine if logged in
			// if logged in
			c.Next()
			// else (AKA not logged in)
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() // default uses Logger and Recovery middleware

	r.Use(m1, m2) // 全局注冊中間件函數m1
	/*
		below is how the log looks like:
		m1 in...
		m2 in...
		index
		m2 out...
		cost: 661.3µs
		m1 out...
	*/

	// GET(relativePath string, handlers ...HandlerFunc) IRoutes
	r.GET("/index", indexHandler)

	// for example, auth middleware can be assigned to certain route group
	dashboardGroup := r.Group("/dashboard", authMiddleware(true))
	{
		dashboardGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "dashboard"})
		})
	}

	r.Run() // :8080 default
}
