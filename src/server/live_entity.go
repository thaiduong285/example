package main

import (
	"context"
	"fmt"

	"github.com/user/examples_app/src/proto/live_entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Printf("id", id)
	result := liveEntity.FindOne(ctx, bson.M{"_id": id})

	fmt.Printf("result", result)

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

	result, err := liveEntity.InsertOne(mongoCtx, data)

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
	// Id, Name, Status := request.GetId(), request.GetName(), request.GetStatus()

	// data := LiveEntity{
	// 	Name, Status,
	// }

	// filter := bson.D{{"id", "name", "status"}}

	// result, err := liveEntity.UpdateOne(ctx.TODO(), filter, data)

	// fmt.Println("result update", result)

	return nil, nil
}

func (server *LiveServiceServer) DeleteLiveEntity(ctx context.Context, request *live_entity.RequestDeleteLiveEntity) (*live_entity.ResponseLiveEntity, error) {
	return nil, nil
}
