package questionmanager

import (
	"context"
	"server-go/database"
	"strings"
)

func GetRandomCategories(n int) (categories []database.Category, err error) {
	err = database.DB.NewSelect().Model(&categories).Order("BY RANDOM()").Limit(n).Scan(context.Background(), &categories)
	return
}

func GetCategoryWithID(id int) (category database.Category, err error) {
	err = database.DB.NewSelect().Model(&category).Where("id = ?", id).Scan(context.Background(), &category)
	return
}

func GetRandomQuestionWithCategoryID(categoryID int) (question database.Question, err error) {
	err = database.DB.NewSelect().Model(&question).Where("category_id = ?", categoryID).OrderExpr("RANDOM()").Limit(1).Scan(context.Background(), &question)
	question.PossibleAnswers = strings.Split(strings.ToLower(question.PossibleAnswersString), ",")
	
	// remove spaces
	for i := range question.PossibleAnswers {
		question.PossibleAnswers[i] = strings.TrimSpace(question.PossibleAnswers[i])
	}
	
	return
}
