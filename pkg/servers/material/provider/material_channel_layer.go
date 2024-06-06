package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
)

type MaterialChannelLayerProvider struct {
	proto.UnimplementedMaterialChannelLayerServer
}

func (u *MaterialChannelLayerProvider) Get(ctx context.Context, in *proto.GetMaterialChannelLayerRequest) (*proto.GetAllMaterialChannelLayerResponse, error) {
	resp := &proto.GetAllMaterialChannelLayerResponse{
		Code: proto.Code_Success,
	}
	f, err := logic.GetMaterialChannelLayer(in)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelLayersToPB(f)
	}
	return resp, nil
}
