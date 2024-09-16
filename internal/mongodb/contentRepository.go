package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"webhook/internal/domain/entities"
	"webhook/pkg/log"
)

type ContentRepository interface {
	Create(ctx context.Context, content entities.Content) error
	GetContent(ctx context.Context, id string) (*entities.Content, error)
	IsContentExist(ctx context.Context, title string) (bool, error)
}

type contentRepository struct {
	mc     *mongo.Collection
	logger *log.Logrus
}

func (c contentRepository) IsContentExist(ctx context.Context, title string) (bool, error) {
	filter := bson.M{"title": title}
	var content entities.Content
	err := c.mc.FindOne(ctx, filter).Decode(&content)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c contentRepository) Create(ctx context.Context, content entities.Content) error {
	_, err := c.mc.InsertOne(ctx, content)
	if err != nil {
		return err
	}
	return nil
}

func (c contentRepository) GetContent(ctx context.Context, id string) (*entities.Content, error) {
	var content entities.Content
	err := c.mc.FindOne(ctx, bson.M{"_id": id}).Decode(&content)
	if err != nil {
		return &content, err
	}
	return &content, nil
}

func NewContentRepository(logger *log.Logrus) ContentRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	// TODO
	opts.Auth = &options.Credential{
		Username: "root",
		Password: "example",
	}
	mc, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	collection := mc.Database("hook").Collection("content")
	return &contentRepository{mc: collection, logger: logger}
}
