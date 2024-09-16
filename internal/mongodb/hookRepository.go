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

// HookRepository defines the interface for hook repository operations.
type HookRepository interface {
	// CreateHook creates a new hook in the database.
	CreateHook(ctx context.Context, hook entities.Hook) error
	// GetHook returns a hook by its ID.
	GetHook(ctx context.Context, id string) (entities.Hook, error)
	CheckIsSendBefore(ctx context.Context, id, to string) (bool, error)
	GetInitialHooks(ctx context.Context) ([]entities.Hook, error)
	SetHookStatus(ctx context.Context, id, status string) error
}

// hookRepository is the concrete implementation of HookRepository.
type hookRepository struct {
	// mc is the MongoDB collection used to interact with the database.
	mc *mongo.Collection
	// logger is used for logging within the repository.
	logger *log.Logrus
}

func (h hookRepository) SetHookStatus(ctx context.Context, id, status string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := h.mc.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (h hookRepository) GetInitialHooks(ctx context.Context) ([]entities.Hook, error) {
	filter := bson.M{"status": "initial"}
	cursor, err := h.mc.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hooks []entities.Hook
	err = cursor.All(ctx, &hooks)
	if err != nil {
		return nil, err
	}
	return hooks, nil
}

func (h hookRepository) CheckIsSendBefore(ctx context.Context, id, to string) (bool, error) {
	filter := bson.M{"_id": id, "to": to, "status": "sent"}
	var hook entities.Hook
	err := h.mc.FindOne(ctx, filter).Decode(&hook)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateHook inserts a new hook into the MongoDB collection.
func (h hookRepository) CreateHook(ctx context.Context, hook entities.Hook) error {
	_, err := h.mc.InsertOne(ctx, hook)
	if err != nil {
		return err
	}
	return nil
}

// GetHook retrieves a hook from the MongoDB collection by its ID.
func (h hookRepository) GetHook(ctx context.Context, id string) (entities.Hook, error) {
	var hook entities.Hook
	err := h.mc.FindOne(ctx, id).Decode(&hook)
	if err != nil {
		return hook, err
	}
	return hook, nil
}

// NewHookRepository creates a new instance of hookRepository with a MongoDB collection and logger.
func NewHookRepository(logger *log.Logrus) HookRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	opts.Auth = &options.Credential{
		Username: "root",
		Password: "example",
	}
	mc, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	collection := mc.Database("hook").Collection("hooks")
	return &hookRepository{mc: collection, logger: logger}
}
