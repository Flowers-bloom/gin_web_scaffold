package dao

import (
	"github.com/flowers-bloom/gin_web_scaffold/common"
	"github.com/flowers-bloom/gin_web_scaffold/common/custom"
	"github.com/flowers-bloom/gin_web_scaffold/common/db"
	"github.com/flowers-bloom/gin_web_scaffold/dto"
	"github.com/flowers-bloom/gin_web_scaffold/model"
	"time"
)

type ScoreDao struct {
}

var scoreDao ScoreDao

func (s *ScoreDao) GetScoreById(id int) model.Score {
	var Score model.Score

	result := db.GetDB().First(&Score, id)
	if result.Error != nil {
		panic(result.Error)
	}

	return Score
}

func (s *ScoreDao) GetAll() common.PageResult {
	var Scores []model.Score

	result := db.GetDB().Find(&Scores)
	if result.Error != nil {
		panic(result.Error)
	}

	rows := make([]interface{}, len(Scores))
	for i, Score := range Scores {
		rows[i] = Score
	}

	return common.PageResult{
		Total:result.RowsAffected,
		Rows:rows,
	}
}

func (s *ScoreDao) Add(Score model.Score) {
	cur := custom.DateTime(time.Now())
	Score.CreatedAt = cur
	Score.UpdatedAt = cur

	result := db.GetDB().Create(&Score)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (s *ScoreDao) DeleteById(id int) {
	db.GetDB().Delete(&model.Score{}, id)
}

func (s *ScoreDao) Update(Score model.Score) {
	db.GetDB().Save(&Score)
}

// 查询某科成绩最高的学生
func (s *ScoreDao) GetMaxScoreStudent(courseId int) dto.MaxScoreStudent {
	var maxScoreStudent dto.MaxScoreStudent

	result := db.GetDB().Raw("SELECT stu.id, stu.name, MAX(sc.value) AS score "+
		"FROM students As stu, scores As sc "+
		"WHERE sc.course_id = ? GROUP BY sc.course_id, stu.id, stu.name LIMIT 1",
		courseId).Scan(&maxScoreStudent)
	if result.Error != nil {
		panic(result.Error)
	}

	return maxScoreStudent
}