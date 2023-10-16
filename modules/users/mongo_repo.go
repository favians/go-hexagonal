package mongodb

import (
	"chat-hex/business/users"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	collection *mongo.Collection
}

type collection struct {
	Id	string	`bson:"id"`
	Email	string	`bson:"email"`
	Password	string	`bson:"password"`
	Name	string	`bson:"name"`
	CurrentRoom	string	`bson:"currentRoom"`
}

func newCollection(user users.User) *collection {
	return &collection{
		user.Id,
		user.Email,
		user.Password,
		user.Name,
		user.CurrentRoom,
	}
}

func (col *collection) ToUser() users.User {
	var user users.User

	user.Id = col.Id
	user.Email = col.Email
	user.Password = col.Password
	user.Name = col.Name
	user.CurrentRoom = col.CurrentRoom

	return user
}

func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("users"),
	}
}

func (repo *MongoDBRepository) EnterChatroom(email string, chatroom string) error {
	filter := bson.M{"email": email}

	update := bson.M{"$set": bson.M{"currentRoom": chatroom}}

    _, err := repo.collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
			return err
    }

		return nil
}

func (repo *MongoDBRepository) LeaveChatroom(email string) error {
	filter := bson.M{"email": email}

	update := bson.M{"$set": bson.M{"currentRoom": nil}}

    _, err := repo.collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
			return err
    }

		return nil
}