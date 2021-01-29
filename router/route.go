package router

import (
	"github.com/flowers-bloom/gin_web_scaffold/api"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) *gin.Engine {
	var studentApi api.StudentApi
	var courseApi api.CourseApi
	var scoreApi api.ScoreApi

	student := router.Group("/student")
	{
		student.GET(":id", studentApi.GetStudentById)
		student.GET("", studentApi.GetAll)
		student.POST("register", studentApi.Register)
		student.PUT("update", studentApi.Update)
		student.DELETE("delete/:id", studentApi.DeleteById)
	}

	course := router.Group("/course")
	{
		course.GET(":id", courseApi.GetCourseById)
		course.GET("", courseApi.GetAll)
		course.POST("add", courseApi.Add)
		course.PUT("update", courseApi.Update)
		course.DELETE("delete/:id", courseApi.DeleteById)
	}

	score := router.Group("/score")
	{
		score.GET("", scoreApi.GetScoreById)
		score.GET("all", scoreApi.GetAll)
		score.GET("max_score_student", scoreApi.MaxScoreStudent)
		score.POST("add", scoreApi.Add)
		score.PUT("update", scoreApi.Update)
		score.DELETE("delete/:id", scoreApi.DeleteById)
	}

	return router
}