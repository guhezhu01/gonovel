package model

import "gonovel/utils/errMsg"

type Article struct {
	Title     string `gorm:"type:varchar(255);not null" json:"title"`
	Id        string `gorm:"type:varchar(20);not null" json:"id"`
	Type      string `gorm:"type:varchar(20)" json:"type"`
	Img       string `gorm:"type:varchar(255)" json:"img"`
	Introduce string `gorm:"type:varchar(255)" json:"introduce"`
	Author    string `gorm:"type:varchar(20);not null" json:"author"`
	Clicks    string `gorm:"type:varchar(255)" json:"clicks"`
	Chapter   string `gorm:"type:varchar(20)" json:"chapter"`
	Words     string `gorm:"type:varchar(255)" json:"words"`
	State     string `gorm:"type:varchar(255)" json:"state"`
}

func GetTypeArticles(type_ string) ([]Article, int, int) {
	var articles []Article
	var total int

	err = db.Model(&articles).Where("type LIKE ?", type_+"%").Find(&articles).Error
	total = len(articles)

	if err != nil {
		return articles, 0, errMsg.CateNotExist
	}
	return articles, total, errMsg.SUCCESS
}

func GetArticle(title string) (Article, int) {
	var article Article
	err := db.Where("title = ?", title).First(&article).Error
	if err != nil {
		return article, errMsg.ArtNotExist
	}
	return article, errMsg.SUCCESS
}

func GetRandArticle() ([]Article, int, int) {
	var data []Article
	var total int
	sql := "SELECT * FROM article WHERE 1  ORDER BY rand() limit 12;"
	err := db.Raw(sql).Scan(&data).Error
	total = len(data)
	if err != nil {
		return data, 0, errMsg.ERROR
	}
	return data, total, errMsg.SUCCESS
}

func DeleteArticle(id string, type_ string) int {
	var article Article
	err := db.Unscoped().Where("id =? and type=?", id, type_).Delete(&article).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}
