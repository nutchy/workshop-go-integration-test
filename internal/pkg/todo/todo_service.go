package todo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type repo interface {
	CreateTodo(context.Context, Todo) error
}

type repoImpl struct {
	client *mongo.Client
}

func (r repoImpl) CreateTodo(ctx context.Context, todo Todo) error {
	_, err := r.client.Database("WorkshopApp").Collection("todos").InsertOne(ctx, todo)
	return err
}

type Todo struct {
	ID     primitive.ObjectID `bson:"_id"`
	TodoID string             `bson:"todo_id"`
	Detail string             `bson:"detail"`
	Done   bool               `bson:"done"`
}

type serviceImpl struct {
	repo repo
}

type service interface {
	CreateTodo(context.Context, Todo) error
}

func NewService() service {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://admin:1234abc@localhost:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.PrimaryPreferred()); err != nil {
		panic(err)
	}

	return &serviceImpl{
		repo: repoImpl{client: client},
	}
}

func (s serviceImpl) CreateTodo(ctx context.Context, todo Todo) error {
	return s.repo.CreateTodo(ctx, todo)
}