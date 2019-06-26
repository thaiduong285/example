package main

import (
	"context"
	"fmt"

	proto "github.com/user/examples_app/src/proto/rectangle"
)

func (server *RectangleServiceServer) CalcRectangleArea(ctx context.Context, request *proto.RequestRectangle) (*proto.ResponseRectangleArea, error) {
	length, height := request.GetWidth(), request.GetHeight()
	fmt.Println("width, height", length, height)
	area := length * height
	return &proto.ResponseRectangleArea{Area: area}, nil
}

func (server *RectangleServiceServer) CalcRectanglePerimeter(ctx context.Context, request *proto.RequestRectangle) (*proto.ResponseRectanglePerimeter, error) {
	length, height := request.GetWidth(), request.GetHeight()
	fmt.Println(length, height)
	perimeter := (length + height) / 2
	return &proto.ResponseRectanglePerimeter{Perimeter: perimeter}, nil
}
