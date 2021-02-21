package api

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/dao"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"github.com/flowers-bloom/gin_web_scaffold/util/convert"
	"github.com/gin-gonic/gin"
)

type ScoreAPI struct {
}

var scoreDao dao.ScoreDao

func (s *ScoreAPI) GetScoreById(c *gin.Context) {
	id := convert.StringToInt(c.Query("id"))
	c.JSON(200, common.Success(0, "查询成功", scoreDao.GetScoreById(id)))
}

func (s *ScoreAPI) GetAll(c *gin.Context) {
	c.JSON(200, common.Success(0, "查询成功", scoreDao.GetAll()))
}

func (s *ScoreAPI) Add(c *gin.Context) {
	var Score model.Score
	if err := c.ShouldBind(&Score); err != nil {
		panic(err)
	}

	scoreDao.Add(Score)
	c.JSON(200, common.Success(0, "添加成功", nil))
}

func (s *ScoreAPI) DeleteById(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))
	scoreDao.DeleteById(id)

	c.JSON(200, common.Success(0, "删除成功", nil))
}

func (s *ScoreAPI) Update(c *gin.Context) {
	var Score model.Score
	if err := c.ShouldBind(&Score); err != nil {
		panic(err)
	}
	scoreDao.Update(Score)

	c.JSON(200, common.Success(0, "更新成功", Score))
}

func (s *ScoreAPI) MaxScoreStudent(c *gin.Context) {
	courseId := convert.StringToInt(c.Query("course_id"))

	c.JSON(200, common.Success(0, "查询成功", scoreDao.GetMaxScoreStudent(courseId)))
}