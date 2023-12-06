package redis_ser

const (
	articleDiggPrefix         = "article_digg"
	articleLookPrefix         = "article_look"
	ArticleCommentCountPrefix = "article_comment_count"
	commentDiggPrefix         = "comment_igg"
)

func NewDigg() CountDB {
	return CountDB{articleDiggPrefix}
}

func NewCommentCount() CountDB {
	return CountDB{articleLookPrefix}
}

func NewArticleLook() CountDB {
	return CountDB{ArticleCommentCountPrefix}
}

func NewCommentDigg() CountDB {
	return CountDB{commentDiggPrefix}
}
