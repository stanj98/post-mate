package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stanj98/post-mate/routes"
)

func main() {
	/*
		Plan:
		3. Saving image attachment as s3 bucket, and storing URL link into object retrieved
		4. Use mongodb instead of memory data in this file
		5. Refer to this: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

	*/

	router := gin.Default()
	router.Static("/static", "./templates")
	router.LoadHTMLGlob("templates/*.html")

	viewRoutes := router.Group("/") 
	{
		viewRoutes.GET("/view-notes", ViewNotes)
		viewRoutes.GET("/:id", ViewNote)
		viewRoutes.GET("/:id/edit", EditNote)
		viewRoutes.GET("/:id/clone", CloneNote)
		viewRoutes.GET("/", CreateNote)
	}

	apiRoutes := router.Group("/api") 
	{
		apiRoutes.GET("/notes", getNotes)
		apiRoutes.GET("/notes/:id", getNote)
		apiRoutes.POST("/notes", createNote)
		apiRoutes.PUT("/notes/:id", editNote)
		apiRoutes.DELETE("/notes/:id", deleteNote)
	}
	
	router.Run("localhost:8080")
}