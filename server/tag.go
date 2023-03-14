package server

import (
	"context"
	"encoding/json"
	"github.com/wow-unbelievable/tag-service/pkg/bapi"
	pb "github.com/wow-unbelievable/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewApi("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}
	tagList := pb.GetTagListReply{}
	if err := json.Unmarshal(body, &tagList); err != nil {
		return nil, err
	}
	return &tagList, nil
}
