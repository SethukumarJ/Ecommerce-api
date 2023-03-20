package repository

import (
	"context"
	"fmt"

	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminDatabaseMongo struct {
	DB *mongo.Client
}

// DeleteAdmin implements interfaces.AdminRepository
func (*adminDatabaseMongo) DeleteAdmin(ctx context.Context, userId string) error {
	panic("unimplemented")
}

// CreateAdmin implements interfaces.AdminRepository
func (db *adminDatabaseMongo) CreateAdmin(ctx context.Context, user domain.Admins) (domain.AdminResponse, error) {
	// Get the "users" collection.
	collection := db.DB.Database("mongo_demo").Collection("users")

	// Insert the user document.
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domain.AdminResponse{}, fmt.Errorf("error inserting user: %v", err)
	}

	// Get the ID of the inserted document and set it on the user.
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return domain.AdminResponse{}, fmt.Errorf("error getting inserted ID: %v", err)
	}
	fmt.Println("id", id)

	return domain.AdminResponse{}, nil
}

// FindAdmin implements interfaces.AdminRepository
func (db *adminDatabaseMongo) FindAdmin(ctx context.Context, id string) (domain.AdminResponse, error) {
	collection := db.DB.Database("mongo_demo").Collection("users")
	var user domain.AdminResponse

	// string to primitive.ObjectID
	pid, _ := primitive.ObjectIDFromHex(id)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": pid}

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return domain.AdminResponse{}, fmt.Errorf("error while finding user %v", err.Error())
	}
	return user, nil
}

func NewAdminMongoRepository(DB *mongo.Client) interfaces.AdminRepository {

	return &adminDatabaseMongo{DB: DB}
}
