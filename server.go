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
		viewRoutes.GET("/view-notes", routes.ViewNotes)
		viewRoutes.GET("/:id", routes.ViewNote)
		viewRoutes.GET("/:id/edit", routes.EditNote)
		viewRoutes.GET("/:id/clone", routes.CloneNote)
		viewRoutes.GET("/", routes.CreateNote)
	}

	apiRoutes := router.Group("/api") 
	{
		apiRoutes.GET("/notes", routes.GetNotesAPI)
		apiRoutes.GET("/notes/:id", routes.GetNoteAPI)
		apiRoutes.POST("/notes", routes.CreateNoteAPI)
		apiRoutes.PUT("/notes/:id", routes.EditNoteAPI)
		apiRoutes.DELETE("/notes/:id", routes.DeleteNoteAPI)
	}
	
	router.Run("localhost:8080")
}