package mongodb

import (
	"context"
	"go-hexagonal/business/messages"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	collection *mongo.Collection
}

type collection struct {
	Id        string    `bson:"_id"`
	Content   string    `bson:"content"`
	Sender    string    `bson:"sender"`
	Timestamp time.Time `bson:"timestamp"`
	ChatRoom  string    `bson:"chatroom"`
}

func newCollection(message messages.Message) *collection {
	return &collection{
		message.Id,
		message.Content,
		message.Sender,
		message.Timestamp,
		message.ChatRoom,
	}
}

func (col *collection) ToMessage() messages.Message {
	var message messages.Message

	message.Id = col.Id
	message.Content = col.Content
	message.Sender = col.Sender
	message.Timestamp = col.Timestamp
	message.ChatRoom = col.ChatRoom

	return message
}

func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("messages"),
	}
}

func (repo *MongoDBRepository) InsertMessage(message messages.Message) error {
	col := newCollection(message)

	_, err := repo.collection.InsertOne(context.Background(), col)
	if err != nil {
		return err
	}

	return nil
}