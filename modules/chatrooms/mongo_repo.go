package mongodb

import (
	"chat-hex/business"
	"chat-hex/business/chatrooms"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	collection *mongo.Collection
}

type collection struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	Code string `bson:"code"`
}

func newCollection(chatroom chatrooms.Chatroom) *collection {
	return &collection{
		chatroom.Id,
		chatroom.Name,
		chatroom.Desc,
		chatroom.Code,
	}
}

func (col *collection) ToChatroom() chatrooms.Chatroom {
	var chatroom chatrooms.Chatroom

	chatroom.Id = col.Id
	chatroom.Name = col.Name
	chatroom.Desc = col.Desc
	chatroom.Code = col.Code

	return chatroom
}

func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("chatrooms"),
	}
}

func (repo *MongoDBRepository) FindChatroomByCode(code string) (*chatrooms.Chatroom, error) {
	var col collection

	filter := bson.M{"code": code}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&col)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, business.ErrNotFound
		}

		return nil, err
	}

	chatroom := col.ToChatroom()

	return &chatroom, nil
}