package todo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTodo(t *testing.T) {
	service := NewService()
	if err := service.CreateTodo(context.Background(), Todo{
		ID:     primitive.NewObjectID(),
		TodoID: uuid.New().String(),
		Detail: "Hello World",
		Done:   false,
	}); err != nil {
		t.Error(err)
	}
}
