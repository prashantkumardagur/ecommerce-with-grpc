package main

import (
	"context"

	pb "github.com/prashantkumardagur/ecommerce-with-grpc/proto"
	"go.mongodb.org/mongo-driver/bson"

	"golang.org/x/crypto/bcrypt"
)

//==================================================================================================

func (s *server) Ping(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "Pong"}, nil
}

//==================================================================================================

func (s *server) GetProfile(ctx context.Context, req *pb.Id) (*pb.User, error) {
	// Get user from database
	dbUser := Collection.FindOne(context.Background(), bson.M{"email": req.Id})
	if dbUser.Err() != nil {
		return nil, dbUser.Err()
	}

	// Decode user
	var user pb.User
	err := dbUser.Decode(&user)
	HandleError(err)

	return &user, nil
}

//==================================================================================================

func (s *server) Register(ctx context.Context, req *pb.User) (*pb.Response, error) {
	// Check if user already exists
	dbUser := Collection.FindOne(context.Background(), bson.M{"email": req.Email})
	if dbUser.Err() == nil {
		return &pb.Response{Success: false, Message: "User already exists"}, nil
	}

	// Hash password
	hash, hasherr := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	HandleError(hasherr)
	req.Password = string(hash)

	// Insert user into database
	_, err := Collection.InsertOne(context.Background(), req)
	HandleError(err)

	return &pb.Response{Success: true, Message: "User registered successfully"}, nil
}

//==================================================================================================

func (s *server) Login(ctx context.Context, req *pb.User) (*pb.Response, error) {
	// Get user from database
	dbUser := Collection.FindOne(context.Background(), bson.M{"email": req.Email})
	if dbUser.Err() != nil {
		return nil, dbUser.Err()
	}

	// Decode user
	var user pb.User
	err := dbUser.Decode(&user)
	HandleError(err)

	// Compare passwords
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if bcryptErr != nil {
		return &pb.Response{Success: false, Message: "Invalid credentials"}, nil
	}

	// Generate JWT token
	token, jwtErr := generateToken(user.Email)
	HandleError(jwtErr)

	return &pb.Response{Success: true, Message: token}, nil
}

//==================================================================================================

func (s *server) Verify(ctx context.Context, req *pb.Id) (*pb.Response, error) {
	// Verify JWT token
	email, err := verifyToken(req.Id)
	if err != nil {
		return &pb.Response{Success: false, Message: "Invalid token"}, nil
	}

	return &pb.Response{Success: true, Message: email}, nil
}
