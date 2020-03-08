package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ArticleDao dao
type ArticleDao struct {
}

// NewArticleDao new dao
func NewArticleDao() *ArticleDao {
	return &ArticleDao{}
}

// getCol get collection
func (a *ArticleDao) getCol(client *mongo.Client) *mongo.Collection {
	return getCollect(client, "article")
}

// Insert insert function
func (a *ArticleDao) Insert(article string, chapter *Article) (err error) {
	client, context := GetSession()
	coll := a.getCol(client)

	// insert article
	if article == "" {
		_, err = coll.InsertOne(context, chapter)
		return
	}

	// insert chapter
	coll.UpdateOne(context, bson.M{"name": article}, bson.M{
		"$push": bson.M{"chapters": chapter},
	})

	return
}

// UpdateTypeName update Type
func (a *ArticleDao) UpdateTypeName(oldName, newName string) (err error) {
	client, context := GetSession()
	coll := a.getCol(client)
	_, err = coll.UpdateMany(context,
		bson.M{"type": oldName},
		bson.M{
			"$set": bson.M{
				"type": newName,
			},
		},
	)
	return
}

// UpdateArticle update article
func (a *ArticleDao) UpdateArticle(oldType, oldArticle string, article *Article) error {
	client, context := GetSession()
	coll := a.getCol(client)
	_, err := coll.UpdateOne(context,
		bson.M{"type": oldType, "name": oldArticle},
		bson.M{
			"$set": bson.M{
				"name":        article.Name,
				"title":       article.Title,
				"type":        article.Type,
				"description": article.Description,
				"content":     article.Content,
				"sort":        article.Sort,
				"prev":        article.Prev,
				"next":        article.Next,
				"good":        article.Good,
				"top":         article.Top,
				"tags":        article.Tags,
				"hits":        article.Hits,
				"author":      article.Author,
				"createdat":   article.CreatedAt,
			},
		},
	)
	return err
}

// Get one
func (a *ArticleDao) Get(ty, articleName, chapter string) (article *Article, err error) {
	client, context := GetSession()
	coll := a.getCol(client)
	if chapter == "" {
		err = coll.FindOne(context, bson.M{"type": ty, "name": articleName}).Decode(&article)
	} else {
		filter := bson.D{
			{Key: "type", Value: ty},
			{Key: "name", Value: articleName},
			{Key: "chapters.name", Value: chapter},
		}
		projection := bson.D{
			{Key: "name", Value: 1},
			{Key: "type", Value: 1},
			{Key: "chapters.$", Value: 1},
		}
		var cursor *mongo.Cursor
		cursor, err = coll.Find(context, filter, options.Find().SetProjection(projection))
		err = cursor.Decode(&a)
	}
	if err != nil {
		return
	}
	return
}
