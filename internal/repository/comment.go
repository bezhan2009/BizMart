package repository

import (
	"BizMart/internal/app/models"
	db2 "BizMart/pkg/db"
	"BizMart/pkg/errs"
	"errors"
	"gorm.io/gorm"
)

// GetProductComments - получение всех комментариев для продукта
func GetProductComments(productID uint) ([]models.Comment, map[uint][]models.Comment, error) {
	var comments []models.Comment
	db := db2.GetDBConn()

	result := db.Where("product_id = ?", productID).Find(&comments)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	// Создаем словарь для комментариев
	commentsDict := make(map[uint][]models.Comment)
	for _, comment := range comments {
		commentsDict[comment.ParentID] = append(commentsDict[comment.ParentID], comment)
	}

	// Получаем корневые комментарии
	var mainComments []models.Comment
	for _, comment := range comments {
		if comment.ParentID == 0 {
			mainComments = append(mainComments, comment)
		}
	}

	return mainComments, commentsDict, nil
}

// CreateComment - создание комментария
func CreateComment(productID uint, userID uint, req models.CommentRequest) (models.Comment, error) {
	comment := models.Comment{
		ProductID:   productID,
		UserID:      userID,
		CommentText: req.CommentText,
		ParentID:    req.ParentID,
	}
	db := db2.GetDBConn()

	// Сохраняем комментарий в БД
	if err := db.Create(&comment).Error; err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

// DeleteComment - удаление комментария с его дочерними комментариями
func DeleteComment(commentID uint, userID uint) error {
	var comment models.Comment
	db := db2.GetDBConn()

	result := db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	// Удаление всех дочерних комментариев
	err := deleteCommentChain(commentID)
	if err != nil {
		return err
	}

	// Удаляем сам комментарий
	return db.Delete(&comment).Error
}

// Удаление дочерних комментариев
func deleteCommentChain(commentID uint) error {
	var childComments []models.Comment
	db := db2.GetDBConn()
	result := db.Where("parent_id = ?", commentID).Find(&childComments)
	if result.Error != nil {
		return result.Error
	}

	for _, childComment := range childComments {
		if err := deleteCommentChain(childComment.ID); err != nil {
			return err
		}
	}

	return db.Where("id = ?", commentID).Delete(&models.Comment{}).Error
}

// BuildCommentTree - построение дерева комментариев
func BuildCommentTree(mainComments []models.Comment, commentsDict map[uint][]models.Comment) []models.CommentTree {
	var result []models.CommentTree
	for _, comment := range mainComments {
		tree := models.CommentTree{
			Comment:  comment,
			Children: buildChildrenTree(comment.ID, commentsDict),
		}
		result = append(result, tree)
	}
	return result
}

func buildChildrenTree(parentID uint, commentsDict map[uint][]models.Comment) []models.CommentTree {
	children := commentsDict[parentID]
	var childrenTree []models.CommentTree

	for _, child := range children {
		childTree := models.CommentTree{
			Comment:  child,
			Children: buildChildrenTree(child.ID, commentsDict),
		}
		childrenTree = append(childrenTree, childTree)
	}

	return childrenTree
}
