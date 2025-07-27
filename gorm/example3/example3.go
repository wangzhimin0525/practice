package example3

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

// 评论状态枚举
const (
	CommentStatusNone        = "无评论"
	CommentStatusHasComments = "有评论"
)

// User 用户表
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	PostCount int    `json:"postCount" gorm:"default:0"`
	Posts     []Post `json:"posts"`
}

// Post 文章表
type Post struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title" gorm:"size:200;not null"`
	UserID        uint      `json:"userID"`
	CommentCount  int       `json:"comment_count" gorm:"default:0"`                       // 文章评论数量统计
	CommentStatus string    `json:"comment_status" gorm:"type:varchar(20);default:'无评论'"` // 评论状态
	Comments      []Comment `json:"comments"`
}

// BeforeCreate 钩子 - 在创建文章时自动更新用户的文章数量统计
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 增加用户的文章计数
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return err
	}

	// 初始化评论状态
	p.CommentStatus = CommentStatusNone
	p.CommentCount = 0

	return nil
}

// Comment 评论表
type Comment struct {
	ID uint `json:"id"`
	//UserID  uint //简化评论人id
	Content string `json:"content" gorm:"type:text"`
	PostID  uint   `json:"postID"`
}

// AfterDelete 钩子 - 在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 文章可能已不存在
	var post Post
	if err := tx.Where("id = ?", c.PostID).First(&post).Error; err != nil {
		return err
	}

	// 减少文章的评论计数
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		return err
	}

	// 检查评论数量并更新状态
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)

	if commentCount == 0 {
		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).
			Updates(map[string]interface{}{
				"comment_status": CommentStatusNone,
				"comment_count":  0,
			}).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetUserPostsWithComments 查询用户及其文章和评论
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]User, error) {
	var user []User

	// 使用Preload预加载关联数据
	err := db.Preload(clause.Associations).Preload("Posts.Comments").Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query user posts: %v", err)
	}
	return user, nil
}

// PostWithCommentCount 相当于PostVO
type PostWithCommentCount struct {
	PostID       uint
	PostTitle    string
	CommentCount int64
}

// 查询评论数量最多的文章信息
func getPostWithMostComments(db *gorm.DB) (*PostWithCommentCount, error) {
	var result PostWithCommentCount

	err := db.Debug().Table("posts").
		Select("posts.id as post_id, posts.title as post_title, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		return nil, fmt.Errorf("failed to query post with most comments: %v", err)
	}

	return &result, nil
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	//// 查询用户ID为3的所有文章及其评论
	//userID := uint(3)
	//users, err := GetUserPostsWithComments(db, userID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//jsonUsers, _ := json.MarshalIndent(users, "", "  ")
	//fmt.Println(string(jsonUsers))

	// 查询评论最多的文章
	post, err := getPostWithMostComments(db)
	if err != nil {
		log.Fatal(err)
	}
	jsonPost, _ := json.MarshalIndent(post, "", "  ")
	fmt.Println(string(jsonPost))
}
