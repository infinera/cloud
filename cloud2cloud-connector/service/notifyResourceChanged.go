package service

import (
	"context"

	raCqrs "github.com/plgd-dev/cloud/resource-aggregate/cqrs"
	pbCQRS "github.com/plgd-dev/cloud/resource-aggregate/pb"
	pbRA "github.com/plgd-dev/cloud/resource-aggregate/pb"
	"github.com/plgd-dev/go-coap/v2/message"
	kitNetGrpc "github.com/plgd-dev/kit/net/grpc"
	kitHttp "github.com/plgd-dev/kit/net/http"
)

func notifyResourceChanged(ctx context.Context, raClient pbRA.ResourceAggregateClient, deviceID, href, userID string, contentType string, body []byte, cmdMeta pbCQRS.CommandMetadata) error {
	coapContentFormat := int32(-1)
	switch contentType {
	case message.AppCBOR.String():
		coapContentFormat = int32(message.AppCBOR)
	case message.AppOcfCbor.String():
		coapContentFormat = int32(message.AppOcfCbor)
	case message.AppJSON.String():
		coapContentFormat = int32(message.AppJSON)
	}

	_, err := raClient.NotifyResourceChanged(kitNetGrpc.CtxWithUserID(ctx, userID), &pbRA.NotifyResourceChangedRequest{
		AuthorizationContext: &pbCQRS.AuthorizationContext{
			DeviceId: deviceID,
		},
		ResourceId:      raCqrs.MakeResourceId(deviceID, kitHttp.CanonicalHref(href)),
		CommandMetadata: &cmdMeta,
		Content: &pbRA.Content{
			Data:              body,
			ContentType:       contentType,
			CoapContentFormat: coapContentFormat,
		},
	})
	return err
}
