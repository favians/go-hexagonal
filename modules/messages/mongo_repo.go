package mongodb

import (
	"context"
	"go-hexagonal/business/messages"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *MongoDBRepository) GetMessagesByChatroom(chatroom string) ([]messages.Message, error) {
	var messages []messages.Message

	filter := bson.M{"chatroom": chatroom}

	option := options.Find()
	option.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	option.SetLimit(50)

	cursor, err := repo.collection.Find(context.Background(), filter, option)
	if err != nil {
		return messages, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var col collection

		err := cursor.Decode(&col)
		if err != nil {
			return messages, err
		}

		message := col.ToMessage()
		messages = append(messages, message)
	}

	return messages, nil
}