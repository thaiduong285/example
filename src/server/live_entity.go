package main

import (
	"context"
	"example/src/proto/live_entity"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LiveEntity struct {
	ID     primitive.ObjectID
	Name   string
	Status string
}

func (server *LiveServiceServer) GetLiveEntity(ctx context.Context, request *live_entity.RequestGetLiveEntity) (*live_entity.ResponseLiveEntity, error) {
	id, err := primitive.ObjectIDFromHex(request.GetId())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	fmt.Printf("id , %d", id)
	result := liveEntity.FindOne(ctx, bson.M{"_id": id})

	fmt.Printf("result, %p", result)

	data := LiveEntity{}

	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", request.GetId(), err))
	}

	dataResponse := &live_entity.ResponseLiveEntity{
		Entity: &live_entity.LiveEntity{
			Id:     id.Hex(),
			Name:   data.Name,
			Status: data.Status,
		},
	}

	return dataResponse, nil
}

func (server *LiveServiceServer) CreateLiveEntity(ctx context.Context, request *live_entity.RequestCreateLiveEntity) (*live_entity.ResponseLiveEntity, error) {
	name := request.GetName()

	data := LiveEntity{
		Name:   name,
		Status: "stop",
	}

	result, err := liveEntity.InsertOne(ctx, data)

	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	id := result.InsertedID.(primitive.ObjectID)

	dataResponse := &live_entity.ResponseLiveEntity{
		Entity: &live_entity.LiveEntity{
			Id:     id.Hex(),
			Name:   name,
			Status: "stop",
		},
	}

	return dataResponse, nil
}

func (server *LiveServiceServer) UpdateLiveEntity(ctx context.Context, request *live_entity.RequestUpdateLiveEntity) (*live_entity.ResponseLiveEntity, error) {
	Id, err := primitive.ObjectIDFromHex(request.GetId())
	Name, Status := request.GetName(), request.GetStatus()

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	filter := bson.M{"_id": Id}

	updateData := bson.M{
		"name":   Name,
		"status": Status,
	}

	result := liveEntity.FindOneAndUpdate(ctx, filter, bson.M{"$set": updateData}, options.FindOneAndUpdate().SetReturnDocument(1))

	decodeLive := LiveEntity{}

	if err := result.Decode(&decodeLive); err != nil {
		return nil, status.Errorf(404, fmt.Sprintf("Could not find entity with Object Id %s: %v", request.GetId(), err))
	}

	dataResponse := &live_entity.ResponseLiveEntity{
		Entity: &live_entity.LiveEntity{
			Id:     Id.Hex(),
			Name:   Name,
			Status: Status,
		},
	}

	return dataResponse, nil
}

func (server *LiveServiceServer) DeleteLiveEntity(ctx context.Context, request *live_entity.RequestDeleteLiveEntity) (*live_entity.ResponseLiveEntity, error) {
	return nil, nil
}
