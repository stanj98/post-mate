package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stanj98/post-mate/handlers"
	"github.com/stanj98/post-mate/database"
)

func main() {
	/*
		Plan:
		3. Saving image attachment as s3 bucket, and storing URL link into object retrieved
		4. Use mongodb instead of memory data in this file
		5. Refer to this: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

	*/
	
	err := database.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal("Connected to MongoDB!")

	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	router.Static("/static", "./templates")
	router.LoadHTMLGlob("templates/*.html")

	viewRoutes := router.Group("/") 
	{
		viewRoutes.GET("/view-notes", handlers.ViewNotes)
		viewRoutes.GET("/:id", handlers.ViewNote)
		viewRoutes.GET("/:id/edit", handlers.EditNote)
		viewRoutes.GET("/:id/clone", handlers.CloneNote)
		viewRoutes.GET("/", handlers.CreateNote)
	}

	apiRoutes := router.Group("/api") 
	{
		apiRoutes.GET("/notes", handlers.GetNotesAPI)
		apiRoutes.GET("/notes/:id", handlers.GetNoteAPI)
		apiRoutes.POST("/notes", handlers.CreateNoteAPI)
		apiRoutes.PUT("/notes/:id", handlers.EditNoteAPI)
		apiRoutes.DELETE("/notes/:id", handlers.DeleteNoteAPI)
	}
	
	router.Run("localhost:8080")
}