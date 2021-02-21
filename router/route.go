package router

import (
	"github.com/flowers-bloom/gin_web_scaffold/api"
	"github.com/flowers-bloom/gin_web_scaffold/middleware"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由器列表
func Init(router *gin.Engine) *gin.Engine {
	var studentAPI api.StudentAPI
	var courseAPI api.CourseAPI
	var scoreAPI api.ScoreAPI

	student := router.Group("/student")
	{
		student.POST("register", studentAPI.Register)
		student.POST("login", studentAPI.Login)
		student.GET("", studentAPI.GetAll)

		authorize := student.Group("", middleware.Authorize())
		{
			authorize.GET(":id", studentAPI.GetStudentById)
			authorize.PUT("update", studentAPI.Update)
			authorize.DELETE("delete/:id", studentAPI.DeleteById)
		}
	}

	course := router.Group("/course")
	{
		course.GET(":id", courseAPI.GetCourseById)
		course.GET("", courseAPI.GetAll)
		course.POST("add", courseAPI.Add)
		course.PUT("update", courseAPI.Update)
		course.DELETE("delete/:id", courseAPI.DeleteById)
	}

	score := router.Group("/score")
	{
		score.GET("", scoreAPI.GetScoreById)
		score.GET("all", scoreAPI.GetAll)
		score.GET("max_score_student", scoreAPI.MaxScoreStudent)
		score.POST("add", scoreAPI.Add)
		score.PUT("update", scoreAPI.Update)
		score.DELETE("delete/:id", scoreAPI.DeleteById)
	}

	return router
}
