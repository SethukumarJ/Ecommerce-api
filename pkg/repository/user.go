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

type userDatabaseMongo struct {
	DB *mongo.Client
}

// DeleteUser implements interfaces.UserRepository
func (*userDatabaseMongo) DeleteUser(ctx context.Context, userid string) error {
	panic("unimplemented")
}

// FindUserByID implements interfaces.UserRepository
func (*userDatabaseMongo) FindUserByID(ctx context.Context, id string) (domain.UserResponse, error) {
	panic("unimplemented")
}

// FindUserByName implements interfaces.UserRepository
func (*userDatabaseMongo) FindUserByName(ctx context.Context, id string) (domain.UserResponse, error) {
	panic("unimplemented")
}

// StoreVerificationDetails implements interfaces.UserRepository
func (*userDatabaseMongo) StoreVerificationDetails(email string, code string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.UserRepository
func (*userDatabaseMongo) VerifyAccount(ct context.Context, email string, code string) error {
	panic("unimplemented")
}

// Delete implements interfaces.UserRepository
func (db *userDatabaseMongo) Delete(ctx context.Context, userId string) error {
	collection := db.DB.Database("mongo_demo").Collection("users")

	// string to primitive.ObjectID
	pid, _ := primitive.ObjectIDFromHex(userId)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": pid}

	// err := collection.FindOne(ctx, filter).Decode(&user)
	res, err := collection.DeleteOne(ctx, filter)
	fmt.Printf("response : %v", res)

	return err
}

// FindByID implements interfaces.UserRepository
func (db *userDatabaseMongo) FindUser(ctx context.Context, id string) (domain.UserResponse, error) {
	collection := db.DB.Database("mongo_demo").Collection("users")
	var user domain.UserResponse

	// string to primitive.ObjectID
	pid, _ := primitive.ObjectIDFromHex(id)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": pid}

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return domain.UserResponse{}, fmt.Errorf("error while finding user %v", err.Error())
	}
	return user, nil
}

// Save implements interfaces.UserRepository
func (db *userDatabaseMongo) CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error) {
	// Get the "users" collection.
	collection := db.DB.Database("mongo_demo").Collection("users")

	// Insert the user document.
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domain.UserResponse{}, fmt.Errorf("error inserting user: %v", err)
	}

	// Get the ID of the inserted document and set it on the user.
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return domain.UserResponse{}, fmt.Errorf("error getting inserted ID: %v", err)
	}
	fmt.Println("id", id)

	return domain.UserResponse{}, nil
}

// FindAll implements interfaces.UserRepository
func (db *userDatabaseMongo) AllUsers(ctx context.Context) ([]domain.UserResponse, error) {
	// Create a slice to hold the users.
	var users []domain.UserResponse

	collection := db.DB.Database("mongo_demo").Collection("users")
	// Retrieve all the documents from the "users" collection.
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding users: %v", err)
	}
	defer cursor.Close(ctx)

	// Loop through the documents and decode them into user structs.
	// Loop through the documents and decode them into user structs.
	for cursor.Next(ctx) {
		var user domain.UserResponse
		err := cursor.Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("error decoding user: %v", err)
		}

		// Check if the ID field is present in the MongoDB document.
		if user.ID != nil {
			fmt.Println(user.ID, "/////////////")
		}

		users = append(users, user)
	}

	// Check for any errors during the iteration.
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users cursor: %v", err)
	}

	return users, nil
}

func NewUserMongoRepository(DB *mongo.Client) interfaces.UserRepository {

	return &userDatabaseMongo{DB: DB}
}
