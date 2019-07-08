package models

type Article struct {
	Model

	Tag Tag
	TagId int
	Title string
	Desc string
	CreatedBy  int
	ModifiedBy int
	Content string
	State int
}

//func (article *Article) ModelFields() []interface{} {
//
//}

func GetArticles(pageNum int, pageSize int, userId interface{}) ([]Article, error) {
	var articles []Article
	// 查询多条
	err := db.Preload("Tag").Offset(pageNum).Limit(pageSize).
		Find(&articles, "created_by = ?", userId).Error
	//rows, err := db.Table("blog_article").Select("blog_article.id,blog_article.title").
	//	Joins("left join blog_tag on blog_article.tag_id=blog_tag.id").
	//	Where("blog_article.created_by = ?", 1).Rows()
	//if err == nil{
	//
	//	for rows.Next() {
	//		var article Article
	//		a := [] interface{} {&article.ID, &article.Title}
	//		_ = rows.Scan(a...)
	//		log.Println(article.ID, article.Title)
	//	}
	//}

	// 查询一条
	//db.Find(&articles, "created_by = ?", 1).Related(&article.Tag, "TagId")
	return articles, err
}

func GetArticleCount(userId interface{}) (count int, err error) {
	err = db.Model(&Article{}).Where("created_by = ?", userId).Count(&count).Error
	return
}

func CreateArticle(article *Article) (*Article, error) {
	err := db.Create(article).Error
	return article, err
}
